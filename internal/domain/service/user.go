package service

import (
	"context"
	"log"
	"time"

	"goex1/internal/data/cache"
	"goex1/internal/domain/do"
	"goex1/internal/domain/enum"
	"goex1/internal/domain/repo"
	"goex1/pkg/code"
	"goex1/pkg/util"
)

type UserService struct {
	*Service
	userRepo repo.UserRepo
}

func NewUserService(service *Service, userRepo repo.UserRepo) *UserService {
	return &UserService{Service: service, userRepo: userRepo}
}

func (u *UserService) RegisterUser(ctx context.Context, user *do.User) error {
	_, exist, err := u.userRepo.GetUserByUserName(ctx, user.UserName)
	if err != nil {
		return err
	}

	if exist {
		return code.ErrUserExist.WithArgs(user.UserName)
	}

	user.Password, err = util.EncryptPass(user.Password)
	if err != nil {
		return code.Wrap("密码加密失败", err)
	}

	return u.userRepo.Create(ctx, user)
}

func (u *UserService) LoginUser(ctx context.Context, userName, password, platform string) (*do.TokenInfo, error) {
	user, exist, err := u.userRepo.GetUserByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, code.ErrUserNotExist.WithArgs(userName)
	}

	if err = util.ComparePass(user.Password, password); err != nil {
		return nil, code.ErrUserPassInvalid.WithArgs(userName)
	}

	return u.GenToken(ctx, user.ID, platform, "")
}

func (u *UserService) RefreshToken(ctx context.Context, refreshToken string) (*do.TokenInfo, error) {
	ok, err := cache.LockRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, code.ErrUserTokenRefreshFaild.WithCause(err).WithArgs(refreshToken)
	}
	defer cache.UnLockRefreshToken(ctx, refreshToken)

	if !ok {
		return nil, code.ErrTooManyRequest.WithArgs("刷新token")
	}

	session, err := cache.GetSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, code.ErrUserTokenInvalid.WithArgs(refreshToken)
	}

	platformSession, err := cache.GetUserSession(ctx, session.UserID, session.Platform)
	if err != nil {
		return nil, err
	}

	if refreshToken != platformSession.RefreshToken {
		log.Println("refresh token已经失效", refreshToken, platformSession.RefreshToken)
		return nil, code.ErrUserTokenInvalid
	}

	return u.GenToken(ctx, session.UserID, session.Platform, session.SessionID)
}

func (u *UserService) GenToken(ctx context.Context, uid uint, platform, sessionID string) (*do.TokenInfo, error) {
	user, exist, err := u.userRepo.GetUserByUserID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, code.ErrUserNotExist.WithArgs(uid)
	}

	if sessionID == "" {
		sessionID = util.GenSessionId(int64(uid))
	}

	accessToken, refreshToken, err := util.GenUserAuthToken(int64(uid))
	if err != nil {
		return nil, code.ErrUserTokenGenFaild.WithCause(err).WithArgs(uid)
	}

	sessionInfo := &do.SessionInfo{
		UserID:       uid,
		SessionID:    sessionID,
		Platform:     platform,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// 删除旧token
	if err = cache.DelOldToken(ctx, sessionInfo); err != nil {
		return nil, code.ErrUserTokenDelFaild.WithCause(err).WithArgs(uid, sessionInfo.Platform)
	}

	// 缓存新token和session
	if err = cache.SetUserSession(ctx, sessionInfo); err != nil {
		return nil, code.ErrUserTokenSetFaild.WithCause(err).WithArgs(uid, sessionInfo.Platform)
	}

	return &do.TokenInfo{
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
		Duration:      int64(enum.AccessTokenDuration.Seconds()),
		SrvCreateTime: time.Now(),
	}, nil
}

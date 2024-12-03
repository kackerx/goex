package appservice

import (
	"context"
	"time"

	"goex1/api/v1/reply"
	"goex1/api/v1/request"
	"goex1/internal/appservice/assembler"
	"goex1/internal/domain/service"
)

type UserAppService struct {
	*AppService
	userService *service.UserService
}

func NewUserAppService(appService *AppService, userService *service.UserService) *UserAppService {
	return &UserAppService{AppService: appService, userService: userService}
}

func (u *UserAppService) Register(ctx context.Context, req *request.RegisterReq) error {
	return u.userService.RegisterUser(ctx, assembler.RegisterDtoToUserDo(req))
}

func (u *UserAppService) Login(ctx context.Context, req *request.LoginReq) (*reply.TokenResp, error) {
	token, err := u.userService.LoginUser(ctx, req.Body.UserName, req.Body.Password, req.Header.Platform)
	if err != nil {
		return nil, err
	}

	return &reply.TokenResp{
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		Duration:      token.Duration,
		SrvCreateTime: token.SrvCreateTime.Format(time.DateTime),
	}, nil
}

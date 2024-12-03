package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"goex1/internal/domain/do"
	"goex1/internal/domain/enum"
)

func DelOldToken(ctx context.Context, sessionInfo *do.SessionInfo) error {
	old, err := GetUserSession(ctx, sessionInfo.UserID, sessionInfo.Platform)
	if err != nil {
		return err
	}

	if old == nil {
		return nil
	}

	if err = Redis().Del(ctx, fmt.Sprintf(enum.RedisKeyAccessToken, old.AccessToken)).Err(); err != nil {
		return err
	}

	return Redis().Expire(ctx, fmt.Sprintf(enum.RedisKeyRefreshToken, old.RefreshToken), enum.OldRefreshTokenHoldingDuration).Err()
}

func GetUserSession(ctx context.Context, uid uint, platform string) (*do.SessionInfo, error) {
	res, err := Redis().HGet(ctx, fmt.Sprintf(enum.RedisKeySession, uid), platform).Bytes()
	if err != nil {
		return nil, err
	}

	var session *do.SessionInfo
	if err = json.Unmarshal(res, session); err != nil {
		return nil, err
	}

	return session, nil
}

func GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*do.SessionInfo, error) {
	result, err := Redis().Get(ctx, fmt.Sprintf(enum.RedisKeyRefreshToken, refreshToken)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var session *do.SessionInfo
	return session, json.Unmarshal([]byte(result), session)
}

func SetUserSession(ctx context.Context, session *do.SessionInfo) error {
	sessionByte, _ := json.Marshal(session)
	if err := Redis().HSet(
		ctx,
		fmt.Sprintf(enum.RedisKeySession, session.SessionID),
		session.Platform,
		sessionByte,
	).Err(); err != nil {
		return err
	}

	return SetUserToken(ctx, session.AccessToken, session.RefreshToken, sessionByte)
}

func SetUserToken(ctx context.Context, accessToken, refreshToken string, sessionBytes []byte) error {
	if err := Redis().Set(
		ctx,
		fmt.Sprintf(enum.RedisKeyAccessToken, accessToken),
		sessionBytes,
		enum.AccessTokenDuration,
	).Err(); err != nil {
		return err
	}

	return Redis().Set(
		ctx,
		fmt.Sprintf(enum.RedisKeyRefreshToken, refreshToken),
		sessionBytes,
		enum.RefreshTokenDuration,
	).Err()
}

func LockRefreshToken(ctx context.Context, refreshToken string) (bool, error) {
	return Redis().SetNX(
		ctx,
		fmt.Sprintf(enum.RedisKeyLockRefreshToken, refreshToken),
		"locked",
		time.Second*10,
	).Result()
}

func UnLockRefreshToken(ctx context.Context, refreshToken string) error {
	return Redis().Del(ctx, fmt.Sprintf(enum.RedisKeyLockRefreshToken, refreshToken)).Err()
}

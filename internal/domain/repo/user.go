package repo

import (
	"context"

	"goex1/internal/domain/do"
)

type UserRepo interface {
	Create(ctx context.Context, user *do.User) error
	GetUserByUserName(ctx context.Context, userName string) (*do.User, bool, error)
	GetUserByUserID(ctx context.Context, userID uint) (*do.User, bool, error)
}

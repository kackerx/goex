package data

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"goex1/internal/data/convertor"
	"goex1/internal/data/model"
	"goex1/internal/domain/do"
	"goex1/internal/domain/repo"
	"goex1/pkg/code"
)

type UserRepo struct {
	*Data
}

func NewUserRepo(data *Data) repo.UserRepo {
	return &UserRepo{Data: data}
}

func (u *UserRepo) Create(ctx context.Context, user *do.User) error {
	userModel := convertor.UserDo2Model(user)
	if err := u.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return code.ErrDBUnknow.WithCause(err)
	}

	return nil
}

func (u *UserRepo) GetUserByUserName(ctx context.Context, userName string) (*do.User, bool, error) {
	return u.getUserByCond(ctx, map[string]any{"user_name": userName})
}

func (u *UserRepo) GetUserByUserID(ctx context.Context, userID uint) (*do.User, bool, error) {
	return u.getUserByCond(ctx, map[string]any{"user_id": userID})
}

func (u *UserRepo) getUserByCond(ctx context.Context, conds map[string]any) (*do.User, bool, error) {
	user := new(model.User)
	if err := u.db.WithContext(ctx).Where(conds).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, code.ErrDBUnknow.WithCause(err)
	}

	return convertor.UserModel2Do(user), true, nil
}

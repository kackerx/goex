package convertor

import (
	"goex1/internal/data/model"
	"goex1/internal/domain/do"
	"goex1/internal/domain/enum"
)

func UserModel2Do(po *model.User) *do.User {
	return &do.User{
		ID:        po.ID,
		UserName:  po.UserName,
		Password:  po.Password,
		NickName:  po.NickName,
		Email:     po.Email,
		Slogan:    po.Slogan,
		Status:    enum.GetUserStatus(po.Status),
		Gender:    enum.GetUserGender(po.Gender),
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}
}

func UserDo2Model(do *do.User) *model.User {
	return &model.User{
		UserName: do.UserName,
		Password: do.Password,
		NickName: do.NickName,
		Email:    do.Email,
		Slogan:   do.Slogan,
	}
}

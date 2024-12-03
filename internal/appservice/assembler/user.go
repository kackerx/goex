package assembler

import (
	"goex1/api/v1/request"
	"goex1/internal/domain/do"
	"goex1/internal/domain/enum"
)

func RegisterDtoToUserDo(req *request.RegisterReq) *do.User {
	return &do.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Slogan:   "",
		NickName: "",
		Gender:   enum.UserGender(req.Gender),
	}
}

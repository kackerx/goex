package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"

	"goex1/api/v1/request"
	"goex1/internal/appservice"
	"goex1/pkg/code"
	"goex1/pkg/resp"
	"goex1/pkg/validator"
)

type UserHandler struct {
	*Handler
	app *appservice.UserAppService
}

func NewUserHandler(handler *Handler, app *appservice.UserAppService) *UserHandler {
	return &UserHandler{Handler: handler, app: app}
}

func (u *UserHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func (u *UserHandler) GetUserInfo(c *gin.Context) {
	uid := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	req := new(request.RegisterReq)
	trans, _ := validator.GetLocalTrans("")
	if err := c.ShouldBind(req); err != nil {
		var errs validator2.ValidationErrors
		if errors.As(err, &errs) {
			resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(validator.RemoveTopStruct(errs.Translate(trans))))
			return
		}
		resp.HandleErr(c, code.ErrParams.WithCause(err))
		return
	}

	if err := u.app.Register(c, req); err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c)
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	req := new(request.LoginReq)
	if err := c.ShouldBind(req.Body); err != nil {
		resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(err))
		return
	}

	if err := c.ShouldBindHeader(req.Header); err != nil {
		resp.HandleErr(c, code.ErrParams.WithCause(err).WithArgs(err))
		return
	}

	token, err := u.app.Login(c, req)
	if err != nil {
		resp.HandleErr(c, err)
		return
	}

	resp.HandleSuccess(c, token)
}

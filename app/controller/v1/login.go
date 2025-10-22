package v1

import (
	"gin/app/middleware"
	"gin/app/model"
	"gin/app/request"
	"gin/app/service"
	"gin/common/base"
	"gin/common/errcode"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	base.BaseController
}

// Token token信息
type Token struct {
	AccessToken        string `json:"accessToken"`
	RefreshToken       string `json:"refreshToken"`
	TokenExpire        int64  `json:"tokenExpire" example:"7200"`
	RefreshTokenExpire int64  `json:"refreshTokenExpire" example:"172800"`
}

type LoginResponse struct {
	Token Token `json:"token"`
	User  model.User
}

// Login 登录
// @Tags 登录相关
// @Summary 登录
// @Description 用户登录
// @Accept json
// @Produce json
// @Param data body request.UserLogin true "登录参数"
// @Success 200 {object} errcode.SuccessResponse{data=LoginResponse} "登录成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/login [post]
func (s *LoginController) Login(c *gin.Context) {
	var (
		srv service.LoginService
		req request.Login
		jwt middleware.Jwt
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.Login{}.GetValidate(req, "login")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	userModel, err := srv.Login(req.Username, req.Password)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	accessToken, refreshToken, tokenExpire, refreshTokenExpire, err := jwt.WithRefresh(userModel.ID, 2*60*60, 2*24*60*60)
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	res := map[string]interface{}{
		"token": map[string]interface{}{
			"accessToken":        accessToken,
			"refreshToken":       refreshToken,
			"tokenExpire":        tokenExpire,
			"refreshTokenExpire": refreshTokenExpire,
		},
		"user": userModel,
	}

	s.Success(c, errcode.Success().WithData(res))
}

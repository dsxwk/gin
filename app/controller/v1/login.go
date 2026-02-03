package v1

import (
	"gin/app/event"
	"gin/app/middleware"
	"gin/app/model"
	"gin/app/request"
	"gin/app/service"
	"gin/common/base"
	"gin/common/errcode"
	"gin/common/global"
	"gin/pkg/eventbus"
	"gin/pkg/lang"
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
// @Success 200 {object} errcode.SuccessResponse{data=LoginResponse} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/login [post]
func (s *LoginController) Login(c *gin.Context) {
	var (
		svc service.LoginService
		req request.Login
		jwt middleware.Jwt
		ctx = c.Request.Context()
	)

	svc.Context.Set(ctx)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.Login{}.GetValidate(req, "Login")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	userModel, err := svc.Login(req.Username, req.Password)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(lang.T(ctx, err.Error(), nil)))
		return
	}

	accessToken, refreshToken, tokenExpire, refreshTokenExpire, err := jwt.WithRefresh(userModel.ID, global.Config.Jwt.Exp, global.Config.Jwt.RefreshExp)
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	// 发布事件
	eventbus.Publish(ctx, event.UserLoginEvent{
		UserId:   userModel.ID,
		Username: userModel.Username,
	})

	s.Success(
		c, errcode.Success().WithMsg(
			lang.T(ctx, "login.success", map[string]interface{}{
				"name": userModel.Username,
			}),
		).WithData(LoginResponse{
			Token{
				AccessToken:        accessToken,
				RefreshToken:       refreshToken,
				TokenExpire:        tokenExpire,
				RefreshTokenExpire: refreshTokenExpire,
			},
			userModel,
		}),
	)
}

// RefreshToken 刷新token
// @Tags 登录相关
// @Summary 刷新token
// @Description 刷新token
// @Accept json
// @Produce json
// @Param token header string true "刷新Token"
// @Success 200 {object} errcode.SuccessResponse{data=Token} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/refresh-token [post]
func (s *LoginController) RefreshToken(c *gin.Context) {
	var (
		svc service.LoginService
		req request.Login
		ctx = c.Request.Context()
	)

	svc.Context.Set(ctx)

	token := c.Request.Header.Get("token")
	req.RefreshToken.Token = token
	// 验证
	err := request.Login{}.GetValidate(req, "RefreshToken")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	accessToken, refreshToken, tokenExpire, refreshTokenExpire, err := svc.RefreshToken(token)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(lang.T(ctx, err.Error(), nil)))
		return
	}

	s.Success(c, errcode.Success().WithData(Token{
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		TokenExpire:        tokenExpire,
		RefreshTokenExpire: refreshTokenExpire,
	}))
}

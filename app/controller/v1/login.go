package v1

import (
	"gin/common/base"
	"gin/common/errcode"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	base.BaseController
}

// Login 登录
// @Router /login [post]
func (s *LoginController) Login(c *gin.Context) {
	s.Success(c, nil, errcode.Success())
}

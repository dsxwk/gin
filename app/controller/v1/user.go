package v1

import (
	"gin/common/base"
	"gin/common/errcode"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

// List 列表
// @Router /user [get]
func (s *UserController) List(c *gin.Context) {
	s.Success(c, nil, errcode.Success())
}

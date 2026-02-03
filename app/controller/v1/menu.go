package v1

import (
	"gin/app/request"
	"gin/app/service"
	"gin/common/base"
	"gin/common/errcode"
	"gin/pkg/lang"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	base.BaseController
}

// List 列表
// @Tags 菜单管理
// @Summary 列表
// @Description 菜单列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Success 200 {object} errcode.SuccessResponse{data=request.PageData{list=[]model.Menu}} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/menu [get]
func (s *MenuController) List(c *gin.Context) {
	var (
		svc service.MenuService
		req request.Menu
		_s  request.Search
		ctx = c.Request.Context()
	)

	svc.Context.Set(ctx)

	err := c.ShouldBind(&_s)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.Menu{}.GetValidate(req, "List")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	res, err := svc.List(req, _s.Search)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(lang.T(ctx, err.Error(), nil)))
		return
	}

	s.Success(c, errcode.Success().WithData(res))
}

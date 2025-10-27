package v1

import (
	"gin/app/model"
	"gin/app/request"
	"gin/app/service"
	"gin/common/base"
	"gin/common/errcode"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
)

type UserController struct {
	base.BaseController
}

// List 列表
// @Tags 用户管理
// @Summary 列表
// @Description 用户列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Success 200 {object} errcode.SuccessResponse{data=request.PageData{list=[]model.User}} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/user [get]
func (s *UserController) List(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.User{}.GetValidate(req, "List")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	res, err := srv.List(req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	s.Success(c, errcode.Success().WithData(res))
}

// Create 创建
// @Tags 用户管理
// @Summary 创建
// @Description 用户创建
// @Param token header string true "认证Token"
// @Param data body request.UserCreate true "创建参数"
// @Success 200 {object} errcode.SuccessResponse{data=model.User} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/user [post]
func (s *UserController) Create(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
		m   model.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.User{}.GetValidate(req, "Create")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	err = copier.Copy(&m, &req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	m, err = srv.Create(m)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	s.Success(c, errcode.Success().WithData(m))
}

// Update 更新
// @Tags 用户管理
// @Summary 更新
// @Description 用户更新
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Param data body request.UserUpdate true "更新参数"
// @Success 200 {object} errcode.SuccessResponse{data=model.User} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/user/{id} [put]
func (s *UserController) Update(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
		m   model.User
		id  int64
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}
	req.ID = id

	// 验证
	err = request.User{}.GetValidate(req, "Update")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	err = copier.Copy(&m, &req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	m, err = srv.Update(m)

	s.Success(c, errcode.Success().WithData(m))
}

// Detail 详情
// @Tags 用户管理
// @Summary 详情
// @Description 用户详情
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Success 200 {object} errcode.SuccessResponse{data=model.User} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/user/{id} [get]
func (s *UserController) Detail(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
		id  int64
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}
	req.ID = id

	// 验证
	err = request.User{}.GetValidate(req, "Detail")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	m, err := srv.Detail(req.ID)

	s.Success(c, errcode.Success().WithData(m))
}

// Delete 删除
// @Tags 用户管理
// @Summary 删除
// @Description 用户删除
// @Param token header string true "认证Token"
// @Param id path int true "用户ID"
// @Success 200 {object} errcode.SuccessResponse{data=model.User} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/user/{id} [delete]
func (s *UserController) Delete(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
		id  int64
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}
	req.ID = id

	// 验证
	err = request.User{}.GetValidate(req, "Delete")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	m, err := srv.Delete(req.ID)

	s.Success(c, errcode.Success().WithData(m))
}

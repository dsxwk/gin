package service

import (
	"gin/app/model"
	"gin/app/request"
	"gin/common/base"
	"gin/common/global"
)

type UserService struct {
	base.BaseService
}

// List 列表
func (s *UserService) List(req request.User) (pageData request.PageData, err error) {
	var (
		m []model.User
	)

	offset, limit := request.Pagination(req.Page, req.PageSize)

	db := global.DB.Model(&m)
	err = db.Count(&pageData.Total).Error
	if err != nil {
		return pageData, err
	}

	err = db.Offset(offset).Limit(limit).Order("id DESC").Find(&m).Error
	if err != nil {
		return pageData, err
	}
	pageData.List = m

	return pageData, nil
}

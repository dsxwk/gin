package service

import (
	"errors"
	"gin/app/model"
	"gin/app/request"
	"gin/common/base"
	"gin/common/global"
	"gin/utils"
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

// Create 创建
func (s *UserService) Create(m model.User) (model.User, error) {
	var (
		count int64
	)

	// 校验用户名是否重复
	err := global.DB.Model(&model.User{}).Where("username = ?", m.Username).Count(&count).Error
	if err != nil {
		return m, err
	}
	if count > 0 {
		return m, errors.New("用户名已存在")
	}

	// 处理密码
	m.Password = utils.BcryptHash(m.Password)
	err = global.DB.Create(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Update 更新
func (s *UserService) Update(m model.User) (model.User, error) {
	var (
		count int64
	)

	// 校验用户名是否重复
	err := global.DB.Model(&model.User{}).Where("username = ? AND id <> ?", m.Username, m.ID).Count(&count).Error
	if err != nil {
		return m, err
	}
	if count > 0 {
		return m, errors.New("用户名已存在")
	}

	err = global.DB.Updates(&m).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Detail 详情
func (s *UserService) Detail(id int64) (m model.User, err error) {
	err = global.DB.First(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

// Delete 删除
func (s *UserService) Delete(id int64) (m model.User, err error) {
	err = global.DB.Delete(&m, id).Error
	if err != nil {
		return m, err
	}

	return m, nil
}

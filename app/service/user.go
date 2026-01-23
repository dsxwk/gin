package service

import (
	"errors"
	"gin/app/model"
	"gin/app/request"
	"gin/common/base"
	"gin/common/global"
	"gin/utils"
	"gin/utils/gorm/search"
	"time"
)

type UserService struct {
	base.BaseService
}

// List 列表
func (s *UserService) List(req request.User, _search map[string]interface{}) (pageData request.PageData, err error) {
	var (
		m []model.User
	)

	offset, limit := request.Pagination(req.Page, req.PageSize)

	db := global.DB.Model(&model.User{}).
		Preload("UserRoles")

	if _search != nil {
		whereSql, args, _err := search.BuildCondition(_search, db, model.User{})
		if _err != nil {
			return pageData, err
		}

		if whereSql != "" {
			db = db.Where(whereSql, args...)
		}
	}

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
func (s *UserService) Update(id int64, m model.User) error {
	var (
		count int64
	)

	// 校验用户名是否重复
	err := global.DB.Model(&model.User{}).Where("username = ? AND id <> ?", m.Username, m.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	data := request.FilterMapByKeys(m, request.UserUpdateKeys)
	data["updated_at"] = time.DateTime

	err = global.DB.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
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

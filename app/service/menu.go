package service

import (
	"gin/app/model"
	"gin/app/request"
	"gin/common/base"
	"gin/common/global"
	"gin/utils/gorm/search"
)

type MenuService struct {
	base.BaseService
}

// List 列表
func (s *MenuService) List(req request.Menu, _search map[string]interface{}) (pageData request.PageData, err error) {
	var (
		m []model.Menu
	)

	offset, limit := request.Pagination(req.Page, req.PageSize)

	db := global.DB.Model(&model.Menu{})

	if _search != nil {
		whereSql, args := search.BuildCondition(_search, db, model.Menu{})

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

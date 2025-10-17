package request

type PageListValidate struct {
	Page     int   `form:"page" validate:"required|int|gt:0" label:"页码"`
	PageSize int   `form:"pageSize" validate:"required|int|gt:0" label:"每页数量"`
	IsPage   *bool `form:"isPage" validate:"required|bool" label:"是否分页"`
}

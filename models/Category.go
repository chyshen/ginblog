// @Author scy
// @Time 2024/7/23 14:32
// @File Category.go

package models

// Category 文章分类表
type Category struct {
	CategoryId uint   `json:"category_id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"type:VARCHAR(20);comment:'分类名称'" binding:"required" label:"分类名称" example:"go"`
	UserId     uint   `json:"user_id" gorm:"comment:'用户ID'"`
}

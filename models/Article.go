// @Author scy
// @Time 2024/7/23 14:06
// @File Article.go

package models

import (
	"gorm.io/gorm"
	"time"
)

// Article 文章
type Article struct {
	ArticleId    uint           `json:"article_id" gorm:"primaryKey;autoIncrement"`
	Title        string         `json:"title" gorm:"type:VARCHAR(200);comment:'标题'" binding:"required" label:"标题" example:"go语言"`
	Content      string         `json:"content" gorm:"type:LONGTEXT;comment:'内容'" binding:"required" label:"内容" example:"Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。"`
	Status       int            `json:"status" gorm:"type:TINYINT UNSIGNED;default:1;comment:'状态'" binding:"required" label:"状态" example:"1"`
	PraiseCount  int            `json:"praise_count" gorm:"default:0;comment:'点赞数'" binding:"required" label:"点赞数" example:"10"`
	CommentCount int            `json:"comment_count" gorm:"default:0;comment:'评论数'" binding:"required" label:"评论数" example:"10"`
	CreatedAt    time.Time      `json:"created_at" gorm:"comment:'创建时间'"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"comment:'更新时间'"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:'删除时间'"`
	CategoryId   uint           `json:"category_id" gorm:"comment:'文章分类ID'"`
	UserId       uint           `json:"user_id" gorm:"comment:'用户ID'"`
}

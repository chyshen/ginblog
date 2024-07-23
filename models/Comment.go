// @Author scy
// @Time 2024/7/23 14:37
// @File Comment.go

package models

import "time"

// Comment 评论
type Comment struct {
	CommentId uint      `json:"comment_id" gorm:"primaryKey;autoIncrement"`
	Content   string    `json:"content" gorm:"type:TEXT;comment:'评论内容'" binding:"required" label:"评论内容" example:"观点阐述到位"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'创建时间'"`
	ArticleId uint      `json:"article_id" gorm:"comment:'文章ID'"`
	UserId    uint      `json:"user_id" gorm:"comment:'用户ID'"`
}

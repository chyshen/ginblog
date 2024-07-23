// @Author scy
// @Time 2024/7/23 14:39
// @File Praise.go

package models

import "time"

// Praise 点赞
type Praise struct {
	ArticleId uint      `json:"article_id" gorm:"primaryKey;comment:'文章ID'"`
	UserId    uint      `json:"user_id" gorm:"primaryKey;comment:'用户ID'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'创建时间'"`
}

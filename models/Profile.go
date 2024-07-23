// @Author scy
// @Time 2024/7/23 15:01
// @File Profile.go

package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name      string `json:"name" gorm:"varchar(20);comment:'昵称'"`
	Desc      string `json:"desc" gorm:"varchar(200);comment:'简介'"`
	Qqchat    string `json:"qq_chat" gorm:"varchar(100);comment:'QQ号'"`
	Wechat    string `json:"wechat" gorm:"varchar(100);comment:'微信号'"`
	Email     string `json:"email" gorm:"varchar(200);comment:'邮箱'"`
	Img       string `json:"img" gorm:"varchar(200);comment:'图片地址'"`
	Avatar    string `json:"avatar" gorm:"varchar(200);comment:'头像'"`
	IcpRecord string `json:"icp_record" gorm:"varchar(200)"`
}

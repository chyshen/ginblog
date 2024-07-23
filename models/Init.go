// @Author scy
// @Time 2024/7/23 15:21
// @File Init.go

package models

import (
	"fmt"
	"github.com/chyshen/ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB
var err error

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.Vcf.GetString("mysql.user"),
		utils.Vcf.GetString("mysql.password"),
		utils.Vcf.GetString("mysql.host"),
		utils.Vcf.GetString("mysql.port"),
		utils.Vcf.GetString("mysql.dbname"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志模式
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用gorm默认事务
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 启动用单数表面，默认`User`会为`Users`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&User{}, &Article{}, &Category{}, &Praise{}, &Comment{}, &Profile{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := DB.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

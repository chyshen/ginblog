// @Author scy
// @Time 2024/7/22 22:52
// @File router.go

package routers

import (
	"github.com/chyshen/ginblog/api/v1"
	"github.com/chyshen/ginblog/middlewares"
	"github.com/chyshen/ginblog/tests/api"
	"github.com/gin-gonic/gin"

	_ "github.com/chyshen/ginblog/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	test := r.Group("/tests/api")
	test.GET("/test", api.Test)

	g1 := r.Group("/api/v1")
	{
		user := g1.Group("/user")
		user.Use(middlewares.JwtToken())
		{
			// 新增用户
			user.POST("/add", v1.UserAdd)
			// 查询单个用户
			user.GET("/query/:id", v1.UserQuery)
			// 查询用户列表
			user.GET("/list", v1.UserList)
			// 更新用户
			user.PUT("/update", v1.UserUpdate)
			// 修改密码
			user.PUT("/password/:id", v1.UserChangePassword)
			// 删除用户
			user.DELETE("/del/:id", v1.UserDel)

		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

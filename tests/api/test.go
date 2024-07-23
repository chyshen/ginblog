//	@Author	scy
//	@Time	2024/7/22 23:16
//	@File	test.go

package api

import "github.com/gin-gonic/gin"

// Test godoc
//	@Summary		测试
//	@Description	测试接口
//	@Tags			测试
//	@Router			/test [get]
func Test(c *gin.Context) {
	c.String(200, "Hello World")
}

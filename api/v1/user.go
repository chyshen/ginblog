//	@Author	scy
//	@Time	2024/7/23 15:41
//	@File	user.go

package v1

import (
	"errors"
	"fmt"
	"github.com/chyshen/ginblog/models"
	"github.com/chyshen/ginblog/types"
	"github.com/chyshen/ginblog/utils/gcode"
	"github.com/chyshen/ginblog/utils/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// UserAdd 添加用户
//
//	@Security		ApiKeyAuth
//	@Summary		新增用户
//	@Description	接口鉴权，需在接口中配置 @Security ApiKeyAuth ,同时主函数中需先启用鉴权，可参考main.go中鉴权注释
//	@Accept			json
//	@Produce		json
//	@Tags			User
//	@Param			User		body		models.UserAddModel	true	"新增用户"
//	@Success		200			{object}	models.UserAddModel
//	@Failure		400			{object}	models.HTTPError
//	@Failure		404			{object}	models.HTTPError
//	@Failure		500			{object}	models.HTTPError
//	@Router			/user/add	[post]
func UserAdd(ctx *gin.Context) {
	var user models.UserAddModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		var validErr validator.ValidationErrors
		if errors.As(err, &validErr) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    gcode.Error,
				"message": validErr.Translate(translator.Trans),
				"data":    user,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    gcode.Error,
			"message": err,
			"data":    user,
		})
		return
	}
	// 验证用户是否存在
	code := user.CheckUser(user.Username)
	if code == gcode.Success {
		models.UserAdd(&user)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data":    user,
	})
}

// UserQuery 查询单个用户
//
//	@Security	ApiKeyAuth
//	@Summary	查询单个用户
//	@Accept		json
//	@Produce	json
//	@Tags		User
//	@Param		id					path	int	true	"用户id"	default(1)
//	@Router		/user/query/{id}	[get]
func UserQuery(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	data, code := models.UserQuery(id)
	fmt.Println(data)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data": types.M{
			"user":  data,
			"total": 1,
		},
	})
}

// UserList 用户列表
//
//	@Security	ApiKeyAuth
//	@Summary	用户列表
//	@Accept		json
//	@Produce	json
//	@Tags		User
//	@Param		username	query	string	false	"用户名"
//	@Param		pagesize	query	int		true	"每页显示的条目"	default(10)
//	@Param		pagenum		query	int		true	"页码"		default(1)
//	@Router		/user/list	[get]
func UserList(ctx *gin.Context) {
	username := ctx.Query("username")
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := models.UserList(username, pageSize, pageNum)
	code := gcode.Success
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data": types.M{
			"user":  data,
			"total": total,
		},
	})
}

// UserUpdate 更新用户
//
//	@Security	ApiKeyAuth
//	@Summary	更新用户
//	@Accept		json
//	@Produce	json
//	@Tags		User
//	@Param		id				query		int						true	"用户id"
//	@Param		User			body		models.UserUpdateModel	true	"更新用户"
//	@Success	200				{object}	models.UserUpdateModel
//	@Router		/user/update	[put]
func UserUpdate(ctx *gin.Context) {
	var user models.UserUpdateModel
	// Query查询参数，如：/api/v1/user/edit?id=6
	id, _ := strconv.Atoi(ctx.Query("id"))
	_ = ctx.ShouldBindJSON(&user)
	var code int
	// 判断用户名是否存在，但是在此如果用户名不变，只修改其他项时也会提示用户名已存在，修改不成功，{"username":"test", "role": 100}
	code = user.CheckUpUser(id, user.Username)
	if code == gcode.Success {
		code = models.UserUpdate(id, &user)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data":    user,
	})
}

// UserChangePassword 修改用户密码
//
//	@Security	ApiKeyAuth
//	@Summary	修改用户密码
//	@Tags		User
//	@Param		id					path	int							true	"用户id"
//	@Param		Password			body	models.UserPasswordModel	true	"修改密码"
//	@Router		/user/password/{id}	[put]
func UserChangePassword(ctx *gin.Context) {
	var user models.UserPasswordModel
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	if err = ctx.ShouldBindJSON(&user); err != nil {
		var validErr validator.ValidationErrors
		if errors.As(err, &validErr) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    gcode.Error,
				"message": validErr.Translate(translator.Trans),
				"data":    user,
			})
			return
		}
	}
	code, data := models.UserChangePassword(id, &user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data": types.M{
			"id":       id,
			"password": data,
		},
	})
}

// UserDel 删除用户
//
//	@Security	ApiKeyAuth
//	@Summary	删除用户
//	@Tags		User
//	@Param		id				path	int	false	"用户id"
//	@Router		/user/del/{id}	[delete]
func UserDel(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	code := models.UserDel(id)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": gcode.Message(code),
		"data":    types.M{"id": id},
	})
}

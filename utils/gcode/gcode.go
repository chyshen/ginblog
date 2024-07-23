// @Author scy
// @Time 2024/7/23 16:07
// @File gcode.go

package gcode

const (
	Success = 200
	Error   = 500

	// 用户模块错误码
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorUserNoRight    = 1004
	ErrorTokenExist     = 1005
	ErrorTokenRuntime   = 1006
	ErrorTokenWrong     = 1007
	ErrorTokenTypeWrong = 1008

	// 分类模块错误码
	ErrorCategoryUsed     = 2001
	ErrorCategoryNotExist = 2002

	// 文章模块错误码
	ErrorArticleNotExist = 3001
)

var msg = map[int]string{
	Success: "请求成功",
	Error:   "请求失败",
	// 用户模块
	ErrorUsernameUsed:  "用户名已存在",
	ErrorPasswordWrong: "密码错误",
	ErrorUserNotExist:  "用户不存在",
	ErrorUserNoRight:   "该用户无权限",
	// Token模块
	ErrorTokenExist:     "Token不存在，请重新登录",
	ErrorTokenRuntime:   "Token已过期，请重新登录",
	ErrorTokenWrong:     "Token不正确，请重新登录",
	ErrorTokenTypeWrong: "Token格式错误，请重新登录",
	// 分类模块
	ErrorCategoryUsed:     "分类已存在",
	ErrorCategoryNotExist: "分类不存在",
	// 文章模块
	ErrorArticleNotExist: "文章不存在",
}

// Message 返回给定code的消息
func Message(code int) string {
	return msg[code]
}

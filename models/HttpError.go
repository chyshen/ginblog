// @Author scy
// @Time 2024/7/23 15:50
// @File HttpError.go

package models

type HTTPError struct {
	Code    int    `json:"code" example:"300"`
	Message string `json:"message" example:"请求失败"`
	Data    any    `json:"data"`
}

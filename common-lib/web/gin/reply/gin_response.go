package reply

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"net/http"
)

// 成功响应 部分 --------
// 提示： 不提供自定义成功响应状态码,方便接收端处理

// Success 默认提示信息为 msg = "操作成功"
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, r.Success())
}

// SuccessMsg 自定义 提示消息
func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, r.SuccessMsg(msg))
}

// SuccessData 使用默认提示消息，并携带数据
func SuccessData[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, r.SuccessData(data))
}

// 失败响应 部分 --------

// FailMsg 自定义错误提示信息，默认 code = 5000
func FailMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, r.FailMsg(msg))
}

// FailCode 从 statusCode 定义错误提示信息
func FailCode(c *gin.Context, code r.StatusCode) {
	c.JSON(http.StatusOK, r.FailCode(code))
}

// FailMsgDetails 自定义错误提示信息和错误细节，默认 code = 5000
func FailMsgDetails[T any](c *gin.Context, msg string, data T) {
	c.JSON(http.StatusOK, r.FailMsgDetails(msg, data))
}

// Fail 自定义 code 和错误提示信息
func Fail(c *gin.Context, code r.StatusCode, msg string) {
	c.JSON(http.StatusOK, r.Fail(code, msg))
}

// FailDetails 自定义 code 和错误提示信息，错误细节
func FailDetails[T any](c *gin.Context, code r.StatusCode, msg string, data T) {
	c.JSON(http.StatusOK, r.FailDetails(code, msg, data))
}

// 通用构造 部分 ----
// 用于不确定是 成功响应还是错误响应的场景

// Code 从 StatusCode 定义响应提示
func Code(c *gin.Context, code r.StatusCode) {
	c.JSON(http.StatusOK, r.NewCode(code))
}

// Of 自定义 code 和 msg
func Of(c *gin.Context, code r.StatusCode, msg string) {
	c.JSON(http.StatusOK, r.New(code, msg))
}

// WithData 自定义 code 和提示信息，并携带数据
func WithData[T any](c *gin.Context, code r.StatusCode, msg string, data T) {
	c.JSON(http.StatusOK, r.NewWithData(code, msg, data))
}

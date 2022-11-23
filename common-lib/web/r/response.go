package r

type Resp struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
}

// RespData 响应结构体
// T 返回数据的数据类型
type RespData[T any] struct {
	Resp
	Data T `json:"data"`
}

// 成功响应 部分 --------
// 提示： 不提供自定义成功响应状态码,方便接收端处理

// Success 默认提示信息为 msg = "操作成功"
func Success() *Resp {
	return New(Ok, Status[Ok])
}

// SuccessMsg 自定义 提示消息
func SuccessMsg(msg string) *Resp {
	return New(Ok, msg)
}

// SuccessData 使用默认提示消息，并携带数据
func SuccessData[T any](data T) *RespData[T] {
	return NewWithData(Ok, Status[Ok], data)
}

// 失败响应 部分 --------

// FailMsg 自定义错误提示信息，默认 code = 5000
func FailMsg(msg string) *Resp {
	return New(Internal, msg)
}

// FailCode 从 statusCode 定义错误提示信息
func FailCode(code StatusCode) *Resp {
	return New(code, Status[code])
}

// FailMsgDetails 自定义错误提示信息和错误细节，默认 code = 5000
func FailMsgDetails[T any](msg string, data T) *RespData[T] {
	return NewWithData(Internal, msg, data)
}

// Fail 自定义 code 和错误提示信息
func Fail(code StatusCode, msg string) *Resp {
	return New(code, msg)
}

// FailDetails 自定义 code 和错误提示信息，错误细节
func FailDetails[T any](code StatusCode, msg string, data T) *RespData[T] {
	return NewWithData(code, msg, data)
}

// 通用构造 部分 ----
// 用于不确定是 成功响应还是错误响应的场景

// NewCode 从 StatusCode 定义响应提示
func NewCode(code StatusCode) *Resp {
	return New(code, Status[code])
}

// New 自定义 code 和 msg
func New(code StatusCode, msg string) *Resp {
	return &Resp{code, msg}
}

// NewWithData 自定义 code 和提示信息，并携带数据
func NewWithData[T any](code StatusCode, msg string, data T) *RespData[T] {
	return &RespData[T]{
		Resp{code, msg},
		data,
	}
}

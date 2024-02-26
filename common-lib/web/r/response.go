package r

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
	"net/http"
	"strconv"
)

type Response[T any] struct {
	Code codes.Code `json:"code"`
	Data T          `json:"data"`
	Msg  string     `json:"message"`
}

func Reply(c *gin.Context, data any) {
	httpCode := http.StatusOK
	resp := Response[any]{Code: codes.OK}
	if data == nil {
		resp.Data = struct{}{}
	} else if s, ok := data.(*status.Status); ok {
		//httpCode = codesToHttpCode(s.Code)
		resp.Code = s.GetCode()
		resp.Msg = s.GetMessage()
		resp.Data = struct{}{}
	} else if _, ok := data.(error); ok {
		//httpCode = http.StatusInternalServerError
		resp.Code = 5e5
		resp.Msg = "内部错误"
		resp.Data = struct{}{}
	} else {
		resp.Data = data
	}
	c.JSON(httpCode, resp)
}

func codesToHttpCode(code codes.Code) int {
	switch strconv.Itoa(int(code))[:1] {
	case "2":
		return http.StatusOK
	case "4":
		return http.StatusBadRequest
	case "5":
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

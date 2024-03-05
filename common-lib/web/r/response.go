package r

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gogoclouds/gogo-services/common-lib/app/check"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
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
	switch v := data.(type) {
	case nil:
		resp.Data = struct{}{}
	case *status.Status:
		//httpCode = codesToHttpCode(s.Code)
		resp.Code = v.GetCode()
		resp.Msg = v.GetMessage()
		if v.Details != nil {
			resp.Data = detailErrorType(c, v.Details)
		} else {
			resp.Data = struct{}{}
		}
	case error:
		//httpCode = http.StatusInternalServerError
		resp.Code = 5e5
		resp.Msg = "内部错误"
		resp.Data = struct{}{}
		logger.Error(v.Error())
	default:
		resp.Data = data
	}
	c.JSON(httpCode, resp)
}

// detailErrorType 处理 validator 的错误进行翻译
func detailErrorType(ctx *gin.Context, ds []any) []any {
	res := make([]any, 0, len(ds))
	for _, d := range ds {
		if err, ok := d.(validator.ValidationErrors); ok {
			e := check.TransErrCtx(ctx, err)
			res = append(res, e.Error())
		}
	}
	return res
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

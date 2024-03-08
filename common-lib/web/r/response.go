package r

import (
	"github.com/gin-gonic/gin"

	cvalidator "github.com/gogoclouds/gogo-services/common-lib/app/validator"

	"github.com/go-playground/validator/v10"
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
	Err  any        `json:"omitempty,err"`
}

func Reply(c *gin.Context, data any) {
	httpCode := http.StatusOK
	resp := Response[any]{Code: codes.OK, Data: struct{}{}}
	switch v := data.(type) {
	case nil:
	case *status.Status:
		//httpCode = codesToHttpCode(s.Code)
		resp.Code = v.GetCode()
		resp.Msg = v.GetMessage()
		if v.Details != nil {
			resp.Err = detailErrorType(c, v.Details)
		}
	case error:
		//httpCode = http.StatusInternalServerError
		resp.Code = 5e5
		resp.Msg = "内部错误" // TODO
		logger.Error(v.Error())
	default:
		resp.Data = data
	}
	c.JSON(httpCode, resp)
}

// detailErrorType 处理 validator 的错误进行翻译
func detailErrorType(ctx *gin.Context, ds []any) []any { // TODO key-value
	res := make([]any, 0, len(ds))
	for _, d := range ds {
		if err, ok := d.(validator.ValidationErrors); ok {
			e := cvalidator.TransErrCtx(ctx, err)
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

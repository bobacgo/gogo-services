package r

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/errs"

	cvalidator "github.com/gogoclouds/gogo-services/common-lib/app/validator"

	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gogoclouds/gogo-services/common-lib/app/logger"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
)

type Response[T any] struct {
	Code codes.Code `json:"code"`
	Data T          `json:"data"`
	Msg  string     `json:"message"`
	Err  any        `json:"err,omitempty"`
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
		resp.Code = errs.InternalError.Code
		resp.Msg = errs.InternalError.Message
		logger.Error(v.Error())
	default:
		resp.Data = data
	}
	c.JSON(httpCode, resp)
}

// detailErrorType 处理 validator 的错误进行翻译
func detailErrorType(ctx *gin.Context, ds []any) []any { // TODO key-value
	for i := 0; i < len(ds); i++ {
		switch v := ds[i].(type) {
		case validator.ValidationErrors:
			e := cvalidator.TransErrCtx(ctx, v)
			ds[i] = e.Error()
		case error:
			ds[i] = v.Error()
		}
	}
	return ds
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

package v1_test

import (
	"fmt"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uhttp"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"testing"
)

var AdminEndpoint = "http://localhost:8080/admin"

func TestLogin(t *testing.T) {
	//w := httptest.NewRecorder()
	request := v1.AdminLoginRequest{
		Username: "admin",
		Password: "admin123",
	}
	resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/login", request)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%v", resp)
	//assert.Equal(t, codes.OK, resp.Code)
	//assert.Equal(t, "pong", w.Body.String())
}

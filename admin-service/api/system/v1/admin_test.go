package v1_test

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/admin-service/api/system/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uhttp"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"strings"
	"testing"
)

var AdminEndpoint = "http://localhost:8080/admin"

var DefaultLoginRequest = &v1.AdminLoginRequest{
	UsernamePasswd: v1.UsernamePasswd{
		Username: "admin",
		Password: "admin123",
	},
}

func GetToken(request *v1.AdminLoginRequest) (string, error) {
	if request == nil {
		request = DefaultLoginRequest
	}

	resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/login", request)
	if err != nil {
		return "", err
	}
	if resp.Code != codes.OK {
		return "", fmt.Errorf("msg: %s, err: %v", resp.Msg, resp.Err)
	}
	return resp.Data.Token, nil
}

func TestRegister(t *testing.T) {
	request := v1.AdminRegisterRequest{
		UsernamePasswd: v1.UsernamePasswd{Username: "test3", Password: "test1"},
		Icon:           "",
		Email:          "",
		Note:           "",
	}
	resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/register", request)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("index:%d, codes: %d msg: %s, err: %v", 0, resp.Code, resp.Msg, resp.Err)
	}
	loginReq := v1.AdminLoginRequest{UsernamePasswd: request.UsernamePasswd}
	_, err = GetToken(&loginReq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	// 测试用例
	// 1.登录成功
	// 1.1 用户名为空,密码为空
	// 2.用户名长的过长
	// 3.用户名不存在
	// 4.单一登录
	// 5.用户登录失败次数
	// 6.用户被禁用
	// 7.用户已注销
	var tests = []struct {
		v1.AdminLoginRequest
		want codes.Code
	}{
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: "admin", Password: "admin123"}}, want: codes.OK},
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: "admin", Password: "admin1231"}}, want: errs.AdminLoginFail.Code},
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: "", Password: ""}}, want: codes.BadRequest},
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: strings.Repeat("a", 21), Password: "admin123"}}, want: codes.BadRequest},
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: "用户名不存在", Password: "admin123"}}, want: errs.AdminLoginFail.Code},
		{AdminLoginRequest: v1.AdminLoginRequest{UsernamePasswd: v1.UsernamePasswd{Username: "admin1", Password: "admin123"}}, want: errs.AdminLoginForbidden.Code},
	}

	for i, test := range tests {
		resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/login", test.AdminLoginRequest)

		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

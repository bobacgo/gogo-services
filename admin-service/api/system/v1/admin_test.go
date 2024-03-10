package v1_test

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uhttp"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"strconv"
	"strings"
	"testing"
	"time"
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
	now := strconv.FormatInt(time.Now().Unix(), 10)
	var tests = []struct {
		v1.AdminRegisterRequest
		want codes.Code
	}{
		// 测试用例
		// 1.注册新账户并登录成功
		// 2.用户名不能重复
		// 3.已删除用户名不能重复
		// 4.邮箱号不能重复
		// 5.邮箱格式

		{v1.AdminRegisterRequest{UsernamePasswd: v1.UsernamePasswd{Username: "new" + now, Password: "admin123"}, Email: now + "@qq.com"}, codes.OK},
		{v1.AdminRegisterRequest{UsernamePasswd: v1.UsernamePasswd{Username: "admin", Password: "admin123"}}, errs.AdminUsernameDuplicated.Code},
		{v1.AdminRegisterRequest{UsernamePasswd: v1.UsernamePasswd{Username: "unusername", Password: "admin123"}}, errs.AdminUnUsernameDuplicated.Code},
		{v1.AdminRegisterRequest{UsernamePasswd: v1.UsernamePasswd{Username: "new1" + now, Password: "admin123"}, Email: now + "@qq.com"}, errs.AdminEmailDuplicated.Code},
		{v1.AdminRegisterRequest{UsernamePasswd: v1.UsernamePasswd{Username: "admin", Password: "admin123"}, Email: "qq.com"}, codes.BadRequest},
	}
	for i, test := range tests {
		resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/register", test.AdminRegisterRequest)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
		if test.want == codes.OK {
			loginReq := v1.AdminLoginRequest{UsernamePasswd: test.UsernamePasswd}
			_, err = GetToken(&loginReq)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestLogin(t *testing.T) {
	var tests = []struct {
		v1.AdminLoginRequest
		want codes.Code
	}{
		// 测试用例
		// 1.登录成功
		// 1.1 用户名为空,密码为空
		// 2.用户名长的过长
		// 3.用户名不存在
		// 4.单一登录
		// 5.用户登录失败次数
		// 6.用户被禁用
		// 7.用户已注销
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

func TestLogout(t *testing.T) {

}

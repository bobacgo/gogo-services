package v1_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/common-lib/pkg/uhttp"
	"github.com/gogoclouds/gogo-services/common-lib/web/r"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/status"
	"github.com/samber/lo"
)

var AdminEndpoint = "http://localhost:8080/admin"

var DefaultLoginRequest = &v1.AdminLoginRequest{
	UsernamePasswd: v1.UsernamePasswd{
		Username: "admin",
		Password: "admin123",
	},
}

func GetToken(request *v1.AdminLoginRequest) (*v1.AdminLoginResponse, error) {
	if request == nil {
		request = DefaultLoginRequest
	}

	resp, err := uhttp.Post[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/login", request)
	if err != nil {
		return nil, err
	}
	if resp.Code != codes.OK {
		switch resp.Code {
		case errs.AdminLoginFail.Code:
			return nil, errs.AdminLoginFail
		case errs.AdminLoginForbidden.Code:
			return nil, errs.AdminLoginForbidden
		default:
			return nil, fmt.Errorf("msg: %s, err: %v", resp.Msg, resp.Err)
		}
	}
	return &resp.Data, nil
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
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := uhttp.NewHttpClient[r.Response[any]](AdminEndpoint+"/logout", http.MethodGet)
	client.Header.Add(middleware.AuthHeader, "Bearer "+token.Token)
	resp, err := client.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("codes: %d msg: %s, err: %v", resp.Code, resp.Msg, resp.Err)
	}
}

func TestRefreshToken(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		v1.AdminRefreshTokenRequest
		want codes.Code
	}{
		// 测试用例
		// 1. 有 aToken 和 rToken
		// 2. 没有 rToken
		// 3. 没有 aToken
		{AdminRefreshTokenRequest: v1.AdminRefreshTokenRequest{AToken: token.Token, RToken: token.RToken}, want: codes.OK},
		{AdminRefreshTokenRequest: v1.AdminRefreshTokenRequest{AToken: token.Token, RToken: ""}, want: codes.BadRequest},
		{AdminRefreshTokenRequest: v1.AdminRefreshTokenRequest{AToken: "", RToken: token.RToken}, want: codes.BadRequest},
	}

	for i, test := range tests {
		client := uhttp.NewHttpClient[r.Response[v1.AdminLoginResponse]](AdminEndpoint+"/refreshToken", http.MethodGet)
		client.Header.Add(middleware.AuthHeader, "Bearer "+test.RToken)
		client.Query.Set("aToken", test.AToken)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestGetSelfInfo(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := uhttp.NewHttpClient[r.Response[v1.UserInfo]](AdminEndpoint+"/info", http.MethodGet)
	client.Header.Add(middleware.AuthHeader, "Bearer "+token.Token)
	resp, err := client.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("codes: %d msg: %s, err: %v", resp.Code, resp.Msg, resp.Err)
	}
}

func TestAdminList(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := uhttp.NewHttpClient[r.Response[v1.AdminListResponse]](AdminEndpoint+"/list", http.MethodGet)
	client.Header.Add(middleware.AuthHeader, "Bearer "+token.Token)
	client.Query.Set("keyword", "admin")
	resp, err := client.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("codes: %d msg: %s, err: %v", resp.Code, resp.Msg, resp.Err)
	}
}

func TestAdminGetItem(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		v1.AdminRequest
		want codes.Code
	}{
		// 测试用例
		// 1.获取成功
		// 2.获取失败(ID不存在)
		{AdminRequest: v1.AdminRequest{ID: 1}, want: codes.OK},
		{AdminRequest: v1.AdminRequest{ID: 0}, want: errs.AdminNotFound.Code},
	}
	for i, test := range tests {
		client := uhttp.NewHttpClient[r.Response[v1.AdminResponse]](AdminEndpoint+"/"+strconv.FormatInt(test.ID, 10), http.MethodGet)
		client.Header.Add(middleware.AuthHeader, "Bearer "+token.Token)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestAdminDelete(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		v1.AdminRequest
		want codes.Code
	}{
		// 测试用例
		// 1.删除成功
		// 2.删除失败(ID不存在)
		{AdminRequest: v1.AdminRequest{ID: 4}, want: codes.OK},
		{AdminRequest: v1.AdminRequest{ID: 0}, want: errs.AdminNotFound.Code},
	}
	for i, test := range tests {
		client := uhttp.NewHttpClient[r.Response[v1.AdminResponse]](AdminEndpoint+"/delete/"+strconv.FormatInt(test.ID, 10), http.MethodPost)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestAdminUpdate(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		v1.AdminUpdateRequest
		want codes.Code
	}{
		// 测试用例
		// 1.更新成功
		// 2.邮箱号重复
		// 3.邮箱号格式错误
		{AdminUpdateRequest: v1.AdminUpdateRequest{ID: 5, Icon: lo.ToPtr("https://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/icon/github_icon_01.png"), Email: lo.ToPtr("123@qq.com"), Nickname: lo.ToPtr("更新测试"), Note: lo.ToPtr("更新测试")}, want: codes.OK},
		{AdminUpdateRequest: v1.AdminUpdateRequest{ID: 5, Email: lo.ToPtr("admin1@163.com")}, want: errs.AdminEmailDuplicated.Code},
		{AdminUpdateRequest: v1.AdminUpdateRequest{ID: 5, Email: lo.ToPtr("123")}, want: codes.BadRequest},
	}
	for i, test := range tests {
		client := uhttp.NewHttpClient[r.Response[v1.AdminResponse]](AdminEndpoint+"/update/"+strconv.FormatInt(test.ID, 10), http.MethodPost)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
		client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
		client.Body, _ = json.Marshal(test.AdminUpdateRequest)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestUpdatePasswd(t *testing.T) {
	token, err := GetToken(&v1.AdminLoginRequest{
		v1.UsernamePasswd{
			Username: "uppasswd", Password: "admin123",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		v1.UpdatePasswordRequest
		Relogin bool
		want    codes.Code
	}{
		// 测试用例
		// 1.更新成功
		// 2.令牌失效
		// 3.修改前密码登录
		// 4.修改后密码登录并获取新的令牌
		// 5.旧密码错误
		// 6.旧密码为空
		// 7.新密码为空
		// 8.账户为空
		// 9.修改为原来的密码(方便下次调试)
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin123", NewPassword: "admin1234"}, want: codes.OK},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin123", NewPassword: "admin1234"}, want: errs.TokenOut.Code},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin123"}, Relogin: true, want: errs.AdminLoginFail.Code},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin1234"}, Relogin: true, want: codes.OK},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin123", NewPassword: "admin1234"}, want: errs.AdminOldPwdErr.Code},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "", NewPassword: "admin123434"}, want: codes.BadRequest},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin123", NewPassword: ""}, want: codes.BadRequest},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "", Password: "admin123", NewPassword: "admin123434"}, want: codes.BadRequest},
		{UpdatePasswordRequest: v1.UpdatePasswordRequest{Username: "uppasswd", Password: "admin1234", NewPassword: "admin123"}, want: codes.OK},
	}

	for i, test := range tests {
		// 更新密码后登录校验
		if test.Relogin {
			token, err = GetToken(&v1.AdminLoginRequest{
				v1.UsernamePasswd{
					Username: test.Username, Password: test.Password,
				},
			})

			if err != nil {
				var serr *status.Status
				if !errors.As(err, &serr) || serr.Code != test.want {
					t.Errorf("index:%d, err: %v", i, err)
					return
				}
			}
			return
		}
		// 更新密码相关校验
		client := uhttp.NewHttpClient[r.Response[v1.AdminPwdErr]](AdminEndpoint+"/updatePassword", http.MethodPost)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
		client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
		client.Body, _ = json.Marshal(test.UpdatePasswordRequest)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestAdminStatus(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		v1.AdminUpdateStatusRequest
		relogin bool
		want    codes.Code
	}{
		// 测试用例
		// 1.修改状态成功
		// 2.登录禁止
		// 3.状态改回来
		// 4.状态不传
		{AdminUpdateStatusRequest: v1.AdminUpdateStatusRequest{ID: 7, Status: lo.ToPtr(false)}, want: codes.OK},
		{relogin: true, want: errs.AdminLoginForbidden.Code},
		{AdminUpdateStatusRequest: v1.AdminUpdateStatusRequest{ID: 7, Status: lo.ToPtr(true)}, want: codes.OK},
		{AdminUpdateStatusRequest: v1.AdminUpdateStatusRequest{ID: 7}, want: codes.BadRequest},
	}
	for i, test := range tests {
		if test.relogin {
			token, err = GetToken(&v1.AdminLoginRequest{
				v1.UsernamePasswd{
					Username: "status0", Password: "admin123",
				},
			})

			if err != nil {
				var serr *status.Status
				if !errors.As(err, &serr) || serr.Code != test.want {
					t.Errorf("index:%d, err: %v", i, err)
					return
				}
			}
			return
		}
		client := uhttp.NewHttpClient[r.Response[any]](AdminEndpoint+"/updateStatus/"+strconv.FormatInt(test.ID, 10), http.MethodPost)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
		client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
		client.Body, _ = json.Marshal(test.AdminUpdateStatusRequest)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index:%d, codes: %d msg: %s, err: %v", i, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestAdminUpdateRole(t *testing.T) {
	// TODO
}

func TestAdminRoleList(t *testing.T) {
	// TODO
}

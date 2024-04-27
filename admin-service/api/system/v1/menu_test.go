package v1_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gogoclouds/gogo-services/admin-service/api/errs"
	v1 "github.com/gogoclouds/gogo-services/admin-service/api/system/v1"
	"github.com/gogoclouds/gogo-services/admin-service/internal/model"
	"github.com/gogoclouds/gogo-services/admin-service/internal/router/middleware"
	"github.com/gogoclouds/gogo-services/framework/pkg/uhttp"
	"github.com/gogoclouds/gogo-services/framework/web/r"
	"github.com/gogoclouds/gogo-services/framework/web/r/codes"
	"github.com/gogoclouds/gogo-services/framework/web/r/page"
	"github.com/samber/lo"
)

var MenuEndpoint = "http://localhost:8080/menu"

func TestMenuList(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := uhttp.NewHttpClient[r.Response[page.Data[model.Menu]]](MenuEndpoint+"/list/"+strconv.FormatInt(12, 10), http.MethodGet)
	client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
	client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
	client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
	client.Query.Set("pageNum", "1")
	client.Query.Set("pageSize", "10")
	resp, err := client.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("codes: %d msg: %s, err: %v", resp.Code, resp.Msg, resp.Err)
	}
}

func TestMenuTreeList(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := uhttp.NewHttpClient[r.Response[any]](MenuEndpoint+"/treeList/", http.MethodGet)
	client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
	client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
	client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
	resp, err := client.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != codes.OK {
		t.Errorf("codes: %d msg: %s, err: %v", resp.Code, resp.Msg, resp.Err)
	}
}

func TestMenuDetails(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name string
		v1.MenuRequest
		want codes.Code
	}{
		{"1.获取存在的数据", v1.MenuRequest{ID: 1}, codes.OK},
		{"2.获取不存在的数据", v1.MenuRequest{ID: 10000}, errs.MenuNotFound.Code},
	}
	for _, test := range tests {
		client := uhttp.NewHttpClient[r.Response[any]](MenuEndpoint+"/"+strconv.FormatInt(test.ID, 10), http.MethodGet)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
		client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("name: %s codes: %d msg: %s, err: %v", test.name, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestMenuAdd(t *testing.T) {
	token, err := GetToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		name string
		v1.MenuCreateRequest
		want codes.Code
	}{
		{"1.创建成功", v1.MenuCreateRequest{ParentID: 0, Title: lo.ToPtr("test"), Name: "test", Icon: lo.ToPtr("test"), Level: lo.ToPtr[int32](0), Sort: lo.ToPtr[int32](0), Hidden: false}, codes.OK},
		{"2.参数校验 Title、Name、Icon 不能为空", v1.MenuCreateRequest{ParentID: 0, Level: lo.ToPtr[int32](0), Sort: lo.ToPtr[int32](0), Hidden: false}, codes.BadRequest},
	}
	for _, test := range tests {
		client := uhttp.NewHttpClient[r.Response[any]](MenuEndpoint+"/create", http.MethodPost)
		client.Header.Set(middleware.AuthHeader, "Bearer "+token.Token)
		client.Header.Add(uhttp.HeaderContentType, uhttp.MIMEJSON)
		client.Header.Add(uhttp.HeaderContentType, uhttp.ContentEncoder)
		client.Body, _ = json.Marshal(test.MenuCreateRequest)
		resp, err := client.Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != test.want {
			t.Errorf("index: %d codes: %d msg: %s, err: %v", test.want, resp.Code, resp.Msg, resp.Err)
		}
	}
}

func TestMenuUpdate(t *testing.T) {

}

func TestMenuUpdateHidden(t *testing.T) {

}

func TestMenuDelete(t *testing.T) {

}

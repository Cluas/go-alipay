package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMiniService_QueryMiniBaseInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
								"alipay_open_mini_baseinfo_query_response": {
									"code": "10000",
									"msg": "Success",
									"app_name": "小程序示例",
									"app_english_name": "demoexample",
									"app_slogan": "提供小程序示例功能",
									"app_logo": "https://appstoreisvpic.alipayobjects.com/prod/04843e84-f1fd-4717-a230-1c72de99aa5d.png",
									"category_names": "航空票务_航空公司;生活服务_室内清洁服务;",
									"app_desc": "小程序官方示例Demo，展示已支持的接口能力及组件。",
									"service_phone": "13110101010",
									"service_email": "example@mail.com",
									"safe_domains": [
										"example.com"
									],
									"package_names": [
										"小程序基础功能"
									]
								}
							}`)
	})

	got, err := client.Mini.QueryMiniBaseInfo(context.Background())
	if err != nil {
		t.Errorf("Mini.QueryMiniBaseInfo returned unexcepted error: %v", err)
	}
	want := &BaseInfo{
		AppName:        "小程序示例",
		AppEnglishName: "demoexample",
		AppSlogan:      "提供小程序示例功能",
		AppLogo:        "https://appstoreisvpic.alipayobjects.com/prod/04843e84-f1fd-4717-a230-1c72de99aa5d.png",
		CategoryNames:  "航空票务_航空公司;生活服务_室内清洁服务;",
		AppDesc:        "小程序官方示例Demo，展示已支持的接口能力及组件。",
		ServicePhone:   "13110101010",
		ServiceEmail:   "example@mail.com",
		SafeDomains:    []string{"example.com"},
		PackageNames:   []string{"小程序基础功能"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryMiniBaseInfo got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryMiniBaseInfo_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
								"alipay_open_mini_baseinfo_query_response": {
									"code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
								}
							}`)
	})
	_, err := client.Mini.QueryMiniBaseInfo(context.Background())
	if err == nil {
		t.Errorf("Mini.QueryMiniBaseInfo excepted error")
	}
}

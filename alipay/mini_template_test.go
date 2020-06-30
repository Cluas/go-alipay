package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMiniService_QueryTemplateUsage(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_template_usage_query_response": {
								"code": "10000",
								"msg": "Success",
								"template_usage_info_list": [
									{
										"mini_app_id": "2018011111111111",
										"app_version": "0.0.1"
									}
								]
							}
						}`)
	})

	got, err := client.Mini.QueryTemplateUsage(context.Background(), &QueryTemplateUsageBiz{
		TemplateID:      "1",
		PageNum:         1,
		PageSize:        10,
		TemplateVersion: "0.0.1",
		BundleID:        "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.BindQrCode returned unexcepted error: %v", err)
	}
	want := &QueryTemplateUsageResp{
		TemplateUsageInfoList: []*TemplateUsageInfo{
			{MiniAppID: "2018011111111111", AppVersion: "0.0.1"},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryTemplateUsage got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryTemplateUsage_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_template_usage_query_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryTemplateUsage(context.Background(), &QueryTemplateUsageBiz{
		TemplateID:      "1",
		PageNum:         1,
		PageSize:        10,
		TemplateVersion: "0.0.1",
		BundleID:        "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.QueryTemplateUsage excepted error")
	}
}

package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMiniService_BindQrCode(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_qrcode_bind_response": {
								"code": "10000",
								"msg": "Success",
								"route_group": "78b59c5b6b2946448bc77e17e544b813"
							}
						}`)
	})

	got, err := client.Mini.BindQrCode(context.Background(), &BindQrCodeBiz{
		RouteURL:        "https://www.yoursite.com/",
		Mode:            "FUZZY",
		PageRedirection: "pages/index/index",
	})
	if err != nil {
		t.Errorf("Mini.BindQrCode returned unexcepted error: %v", err)
	}
	want := &BindQrCodeResp{
		RouteGroup: "78b59c5b6b2946448bc77e17e544b813",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.BindQrCode got %+v, want %+v", got, want)
	}
}

func TestMiniService_BindQrCode_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_qrcode_bind_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.BindQrCode(context.Background(), &BindQrCodeBiz{
		RouteURL:        "https://www.yoursite.com/",
		Mode:            "FUZZY",
		PageRedirection: "pages/index/index",
	})
	if err == nil {
		t.Errorf("Mini.BindQrCode excepted error")
	}
}

func TestMiniService_UnbindQrCode(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
						"alipay_open_mini_qrcode_unbind_response": {
							"code": "10000",
							"msg": "Success"
						}
					}`)
	})

	err := client.Mini.UnbindQrCode(context.Background(), &UnbindQrCodeBiz{RouteGroup: "78b59c5b6b2946448bc77e17e544b813"})
	if err != nil {
		t.Errorf("Mini.UnbindQrCode returned unexcepted error: %v", err)
	}
}

func TestMiniService_UnbindQrCode_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
								"alipay_open_mini_qrcode_unbind_response": {
									"code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
								}
							}`)
	})
	err := client.Mini.UnbindQrCode(context.Background(), &UnbindQrCodeBiz{RouteGroup: "78b59c5b6b2946448bc77e17e544b813"})
	if err == nil {
		t.Errorf("Mini.UnbindQrCode excepted error")
	}
}

package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMiniService_CancelExperience(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_cancel_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.CancelExperience(context.Background(), &CancelExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})

	if err != nil {
		t.Errorf("Mini.CancelExperience returned unexcepted error: %v", err)
	}
}

func TestMiniService_CancelExperience_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_cancel_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	err := client.Mini.CancelExperience(context.Background(), &CancelExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})

	if err == nil {
		t.Errorf("Mini.CancelExperience excepted error")
	}
}

func TestMiniService_CreateExperience(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_create_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})
	err := client.Mini.CreateExperience(context.Background(), &CreateExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})

	if err != nil {
		t.Errorf("Mini.CreateExperience returned unexcepted error: %v", err)
	}
}

func TestMiniService_CreateExperience_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_create_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	err := client.Mini.CreateExperience(context.Background(), &CreateExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})

	if err == nil {
		t.Errorf("Mini.CreateExperience excepted error")
	}
}

func TestMiniService_QueryExperience(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_query_response": {
								"code": "10000",
								"msg": "Success",
								"status": "notExpVersion",
								"exp_qr_code_url": "https://mobilecodec.alipay.com/show.htm?code=s4x06980mfxeaok1f3zvq8d"
							}
						}`)
	})
	want := &ExperienceStatus{
		Status:       "notExpVersion",
		ExpQrCodeURL: "https://mobilecodec.alipay.com/show.htm?code=s4x06980mfxeaok1f3zvq8d",
	}
	got, err := client.Mini.QueryExperience(context.Background(), &QueryExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.QueryExperience returned unexcepted error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryExperience got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryExperience_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_experience_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	_, err := client.Mini.QueryExperience(context.Background(), &QueryExperienceBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})

	if err == nil {
		t.Errorf("Mini.QueryExperience excepted error")
	}
}

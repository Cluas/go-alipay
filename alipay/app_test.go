package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAppService_CreateMember(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_members_create_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.App.CreateMember(context.Background(), &CreateAppMemberBiz{
		LogonID: "test_id",
		Role:    "DEVELOPER",
	})

	if err != nil {
		t.Errorf("App.CreateMember returned unexcepted error: %v", err)
	}
}

func TestAppService_CreateMember_error(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_members_create_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	err := client.App.CreateMember(context.Background(), &CreateAppMemberBiz{
		LogonID: "test_id",
		Role:    "DEVELOPER",
	})

	if err == nil {
		t.Errorf("App.CreateMember excepted error")
	}
}

func TestAppService_DeleteMember(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_members_delete_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.App.DeleteMember(context.Background(), &DeleteMemberBiz{
		UserID: "test_id",
		Role:   "DEVELOPER",
	})
	if err != nil {
		t.Errorf("App.CreateMember returned unexcepted error: %v", err)
	}
}

func TestAppService_DeleteMember_error(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_members_delete_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	err := client.App.DeleteMember(context.Background(), &DeleteMemberBiz{
		UserID: "test_id",
		Role:   "DEVELOPER",
	})

	if err == nil {
		t.Errorf("App.DeleteMember excepted error")
	}
}

func TestAppService_QueryAppMembers(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_app_members_query_response": {
								"code": "10000",
								"msg": "Success",
								"app_member_info_list": [
									{
										"user_id": "20881234123412341234",
										"nick_name": "张三",
										"portrait": "http://imageg.alipay.com/1",
										"status": "VALID",
										"gmt_join": "2017-08-12",
										"logon_id": "test@e*****e.com",
										"gmt_invite": "2017-09-08 12:00:00",
										"role": "DEVELOPER"
									}
								]
							}
						}`)
	})
	want := &QueryAppMembersResp{AppMemberInfoList: []*AppMemberInfo{
		{
			UserID:    "20881234123412341234",
			NickName:  "张三",
			Portrait:  "http://imageg.alipay.com/1",
			Status:    "VALID",
			GmtJoin:   "2017-08-12",
			LogonID:   "test@e*****e.com",
			GmtInvite: "2017-09-08 12:00:00",
			Role:      "DEVELOPER",
		},
	}}
	got, err := client.App.QueryAppMembers(context.Background(), &QueryAppMembersBiz{Role: "DEVELOPER"})
	if err != nil {
		t.Errorf("App.QueryAppMembers returned unexcepted error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("App.QueryAppMembers got %+v, want %+v", got, want)
	}

}

func TestAppService_QueryAppMembers_error(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_app_members_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	_, err := client.App.QueryAppMembers(context.Background(), &QueryAppMembersBiz{Role: "DEVELOPER"})
	if err == nil {
		t.Errorf("App.QueryAppMembers excepted error")
	}
}

func TestAppService_CreateAppQRCode(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_qrcode_create_response": {
								"code": "10000",
								"msg": "Success",
								"qr_code_url": "http://mmtcdp.stable.alipay.net/wsdk/img?fileid=A*lSbPT5i9C1wAAAAAAAAAAABjAQAAAA&t=9005d7f574f30246b89c20c17302115f&bz=mmtcafts&"
							}
							}`)
	})
	want := &CreateAppQRCodeResp{
		QRCodeURL: "http://mmtcdp.stable.alipay.net/wsdk/img?fileid=A*lSbPT5i9C1wAAAAAAAAAAABjAQAAAA&t=9005d7f574f30246b89c20c17302115f&bz=mmtcafts&",
	}
	got, err := client.App.CreateAppQRCode(context.Background(), &CreateAppQRCodeBiz{
		URLParam:   "pages/index/index",
		QueryParam: "x=1",
		Describe:   "test",
	})
	if err != nil {
		t.Errorf("App.CreateAppQRCode returned unexcepted error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("App.CreateAppQRCode got %+v, want %+v", got, want)
	}

}

func TestAppService_CreateAppQRCode_error(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_qrcode_create_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})

	_, err := client.App.CreateAppQRCode(context.Background(), &CreateAppQRCodeBiz{
		URLParam:   "pages/index/index",
		QueryParam: "x=1",
		Describe:   "test",
	})
	if err == nil {
		t.Errorf("App.CreateAppQRCode excepted error")
	}
}

package alipay

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestMiniService_QueryMiniBaseInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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

	got, err := client.Mini.QueryBaseInfo(context.Background())
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

func TestMiniService_ModifyBaseInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_safedomain_create_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.ModifyBaseInfo(context.Background(), &ModifyBaseInfoBiz{
		AppName:         "小程序demo",
		AppEnglishName:  "demoexample",
		AppSlogan:       "这是一个小程序示例",
		AppLogo:         os.NewFile(0, "123"),
		AppCategoryIDs:  "11_12;12_13",
		AppDesc:         "这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述",
		ServicePhone:    "13110101010",
		ServiceEmail:    "example@mail.com",
		MiniCategoryIDs: "XS1001_XS2001_XS3002;XS1011_XS2089;XS1002_XS2008_XS3024",
	})
	if err != nil {
		t.Errorf("Mini.ModifyBaseInfo returned unexcepted error: %v", err)
	}
}

func TestMiniService_ModifyBaseInfo_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_baseinfo_modify_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.ModifyBaseInfo(context.Background(), &ModifyBaseInfoBiz{
		AppName:         "小程序demo",
		AppEnglishName:  "demoexample",
		AppSlogan:       "这是一个小程序示例",
		AppLogo:         os.NewFile(0, "123"),
		AppCategoryIDs:  "11_12;12_13",
		AppDesc:         "这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述",
		ServicePhone:    "13110101010",
		ServiceEmail:    "example@mail.com",
		MiniCategoryIDs: "XS1001_XS2001_XS3002;XS1011_XS2089;XS1002_XS2008_XS3024",
	})
	if err == nil {
		t.Errorf("Mini.ModifyBaseInfo excepted error")
	}
}

func TestMiniService_CreateSafeDomain(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
						"alipay_open_mini_safedomain_create_response": {
							"code": "10000",
							"msg": "Success"
						}
					}`)
	})

	err := client.Mini.CreateSafeDomain(context.Background(), &CreateSafeDomainBiz{SafeDomain: "example.com"})
	if err != nil {
		t.Errorf("Mini.CreateSafeDomain returned unexcepted error: %v", err)
	}
}

func TestMiniService_CreateSafeDomain_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
								"alipay_open_mini_baseinfo_query_response": {
									"code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
								}
							}`)
	})
	err := client.Mini.CreateSafeDomain(context.Background(), nil)
	if err == nil {
		t.Errorf("Mini.CreateSafeDomain excepted error")
	}
}

func TestMiniService_DetectRiskContent(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_security_risk_content_detect_response": {
								"code": "10000",
								"msg": "Success",
								"action": "REJECTED",
								"keywords": [ "张三", "李四" ],
								"unique_id": "0ba600421493362500440513027526"
							}
						}`)
	})

	got, err := client.Mini.DetectRiskContent(context.Background(), &DetectRiskContentBiz{
		Content: "张三，李四",
	})
	if err != nil {
		t.Errorf("Mini.DetectRiskContent returned unexcepted error: %v", err)
	}
	want := &DetectRiskContentResp{
		Action:   "REJECTED",
		Keywords: []string{"张三", "李四"},
		UniqueID: "0ba600421493362500440513027526",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.DetectRiskContent got %+v, want %+v", got, want)
	}
}

func TestMiniService_DetectRiskContent_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_security_risk_content_detect_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.DetectRiskContent(context.Background(), &DetectRiskContentBiz{
		Content: "张三，李四",
	})
	if err == nil {
		t.Errorf("Mini.DetectRiskContent excepted error")
	}
}

func TestMiniService_QueryTinyAppExist(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_tinyapp_exist_query_response": {
								"code": "10000",
								"msg": "Success",
								"exist_mini": "true"
							}
						}`)
	})

	got, err := client.Mini.QueryTinyAppExist(context.Background(), &QueryTinyAppExistBiz{
		PID: "2088301371981234",
	})
	if err != nil {
		t.Errorf("Mini.QueryTinyAppExist returned unexcepted error: %v", err)
	}
	want := &QueryTinyAppExistResp{ExistMini: "true"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryTinyAppExist got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryTinyAppExist_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_tinyapp_exist_query_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryTinyAppExist(context.Background(), &QueryTinyAppExistBiz{
		PID: "2088301371981234",
	})
	if err == nil {
		t.Errorf("Mini.QueryTinyAppExist excepted error")
	}
}

func TestMiniService_QueryCategory(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_category_query_response": {
								"code": "10000",
								"msg": "Success",
								"mini_category_list": [
									{
										"category_id": "107396",
										"category_name": "公共交通",
										"parent_category_id": "0",
										"has_child": true,
										"need_license": true,
										"need_out_door_pic": true,
										"need_special_license": true
									}
								],
								"category_list": [
									{
										"category_id": "107396",
										"category_name": "公共交通",
										"parent_category_id": "0",
										"has_child": true,
										"need_license": true,
										"need_out_door_pic": true,
										"need_special_license": true
									}
								]
							}
						}`)
	})

	got, err := client.Mini.QueryCategory(context.Background(), &QueryCategoryBiz{
		IsFilter: true,
	})
	if err != nil {
		t.Errorf("Mini.QueryTinyAppExist returned unexcepted error: %v", err)
	}
	want := &QueryCategoryResp{
		MiniCategoryList: []*MiniAppCategory{
			{CategoryID: "107396", CategoryName: "公共交通", ParentCategoryID: "0", HasChild: true, NeedLicense: true, NeedOutDoorPic: true, NeedSpecialLicense: true},
		},
		CategoryList: []*MiniAppCategory{
			{CategoryID: "107396", CategoryName: "公共交通", ParentCategoryID: "0", HasChild: true, NeedLicense: true, NeedOutDoorPic: true, NeedSpecialLicense: true},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryCategory got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryCategory_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_category_query_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryCategory(context.Background(), &QueryCategoryBiz{
		IsFilter: true,
	})
	if err == nil {
		t.Errorf("Mini.QueryCategory excepted error")
	}
}

func TestMiniService_CertifyIndividualBusiness(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_individual_business_certify_response": {
								"code": "10000",
								"msg": "Success",
								"certify_result": true
							}
						}`)
	})

	got, err := client.Mini.CertifyIndividualBusiness(context.Background(), &CertifyIndividualBusinessBiz{
		LiceseNo:  "1235234234123124234234",
		LicesePic: "/9j/Qnl0ZUFycmF5T3V0cHV0U3RyZWFtIG91dHB1dCA9IG5ldyBCeXRlQ中间缩略Skge30=",
	})
	if err != nil {
		t.Errorf("Mini.CertifyIndividualBusiness returned unexcepted error: %v", err)
	}
	want := &CertifyIndividualBusinessResp{
		CertifyResult: true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.CertifyIndividualBusiness got %+v, want %+v", got, want)
	}
}

func TestMiniService_CertifyIndividualBusiness_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_individual_business_certify_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.CertifyIndividualBusiness(context.Background(), &CertifyIndividualBusinessBiz{
		LiceseNo:  "1235234234123124234234",
		LicesePic: "/9j/Qnl0ZUFycmF5T3V0cHV0U3RyZWFtIG91dHB1dCA9IG5ldyBCeXRlQ中间缩略Skge30=",
	})
	if err == nil {
		t.Errorf("Mini.CertifyIndividualBusiness excepted error")
	}
}

func TestMiniService_SyncContent(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_content_sync_response": {
								"code": "10000",
								"msg": "Success",
								"result_data": "true"
							}
						}`)
	})

	got, err := client.Mini.SyncContent(context.Background(), &SyncContentBiz{
		ContentType: "SHOP",
		ContentData: `{"shopIds": ["2020041300077000000024065718"]}`,
		Operation:   "batchBind",
		ExtendInfo:  `{"key": "val"}`,
	})
	if err != nil {
		t.Errorf("Mini.SyncContent returned unexcepted error: %v", err)
	}
	want := &SyncContentResp{
		ResultData: "true",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.SyncContent got %+v, want %+v", got, want)
	}
}

func TestMiniService_SyncContent_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, `{
							"alipay_open_mini_content_sync_response": {
								"code": "20000",
								"msg": "Service Currently Unavailable",
								"sub_code": "isp.unknow-error",
								"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.SyncContent(context.Background(), &SyncContentBiz{
		ContentType: "SHOP",
		ContentData: `{"shopIds": ["2020041300077000000024065718"]}`,
		Operation:   "batchBind",
		ExtendInfo:  `{"key": "val"}`,
	})
	if err == nil {
		t.Errorf("Mini.CertifyIndividualBusiness excepted error")
	}
}

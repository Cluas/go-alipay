package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMiniService_ApplyVersionAudit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_audit_apply_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.ApplyVersionAudit(context.Background(), &ApplyVersionAuditBiz{
		LicenseName:      "营业执照名称",
		FirstLicensePic:  []byte("图片1"),
		SecondLicensePic: []byte("图片2"),
		ThirdLicensePic:  []byte("图片3"),
		FourthLicensePic: []byte("图片4"),
		FifthLicensePic:  []byte("图片5"),
		LicenseValidDate: "9999-12-31",
		OutDoorPic:       []byte("OutDoorPic"),
		AppVersion:       "0.0.1",
		AppName:          "小程序示例",
		AppEnglishName:   "demo example",
		AppSlogan:        "这是一个支付示例",
		AppLogo:          []byte("AppLogo"),
		AppCategoryIDs:   "11_12;12_13",
		AppDesc:          "这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述这是一个小程序的描述",
		ServicePhone:     "13110101010",
		ServiceEmail:     "example@mail.com",
		VersionDesc:      "小程序版本描述小程序版本描述小程序版本描述小程序版本描述小程序版本描述小程序版本描述",
		Memo:             "小程序示例",
		RegionType:       "LOCATION",
		ServiceRegionInfo: []*RegionInfo{
			{
				ProvinceCode: "310000",
				ProvinceName: "浙江省",
				CityCode:     "310000",
				CityName:     "杭州市",
				AreaCode:     "311100",
				AreaName:     "余杭区",
			},
		},
		FirstScreenShot:         []byte("FirstScreenShot"),
		SecondScreenShot:        []byte("SecondScreenShot"),
		ThirdScreenShot:         []byte("ThirdScreenShot"),
		FourthScreenShot:        []byte("FourthScreenShot"),
		FifthScreenShot:         []byte("FifthScreenShot"),
		LicenseNo:               "LicenseNo",
		FirstSpecialLicensePic:  []byte("FirstSpecialLicensePic"),
		SecondSpecialLicensePic: []byte("SecondSpecialLicensePic"),
		ThirdSpecialLicensePic:  []byte("ThirdSpecialLicensePic"),
		TestAccount:             "TestAccount",
		TestPassword:            "TestPassword",
		TestFileName:            []byte("TestFileName"),
		BundleID:                "com.alipay.alipaywallet",
	})

	if err != nil {
		t.Errorf("Mini.CancelExperience returned unexcepted error: %v", err)
	}
}

func TestMiniService_ApplyVersionAudit_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_app_members_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.ApplyVersionAudit(context.Background(), &ApplyVersionAuditBiz{})
	if err == nil {
		t.Errorf("Mini.ApplyVersionAudit excepted error")
	}
}

func TestMiniService_CancelGrayVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_gray_cancel_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.CancelGrayVersion(context.Background(), &CancelGrayVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.CancelGrayVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_CancelGrayVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_gray_version_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.CancelGrayVersion(context.Background(), &CancelGrayVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.CancelGrayVersion excepted error")
	}
}

func TestMiniService_CancelVersionAudit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_audit_cancel_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.CancelVersionAudit(context.Background(), &CancelVersionAuditBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.CancelVersionAudit returned unexcepted error: %v", err)
	}
}

func TestMiniService_CancelVersionAudit_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_audit_cancel_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.CancelVersionAudit(context.Background(), &CancelVersionAuditBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.CancelVersionAudit excepted error")
	}
}

func TestMiniService_CancelVersionAudited(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_audited_cancel_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.CancelVersionAudited(context.Background(), &CancelVersionAuditedBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.CancelVersionAudited returned unexcepted error: %v", err)
	}
}

func TestMiniService_CancelVersionAudited_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_audited_cancel_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.CancelVersionAudited(context.Background(), &CancelVersionAuditedBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.CancelVersionAudited excepted error")
	}
}

func TestMiniService_DeleteVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_delete_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.DeleteVersion(context.Background(), &DeleteVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.DeleteVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_DeleteVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_delete_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.DeleteVersion(context.Background(), &DeleteVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.DeleteVersion excepted error")
	}
}

func TestMiniService_OfflineVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_offline_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.OfflineVersion(context.Background(), &OfflineVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.OfflineVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_OfflineVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_offline_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.OfflineVersion(context.Background(), &OfflineVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.OfflineVersion excepted error")
	}
}

func TestMiniService_OnlineGrayVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_gray_offline_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.OnlineGrayVersion(context.Background(), &OnlineGrayVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.OfflineVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_OnlineGrayVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_gray_offline_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.OnlineGrayVersion(context.Background(), &OnlineGrayVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.OnlineGrayVersion excepted error")
	}
}

func TestMiniService_OnlineVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_online_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.OnlineVersion(context.Background(), &OnlineVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.OnlineVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_OnlineVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_online_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.OnlineVersion(context.Background(), &OnlineVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.OnlineVersion excepted error")
	}
}

func TestMiniService_QueryVersionDetail(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_detail_query_response": {
								"code": "10000",
								"msg": "Success",
								"app_version": "0.0.1",
								"app_name": "小程序demo",
								"app_english_name": "demo example",
								"app_logo": "http://image.aaaa.alipay.com",
								"version_desc": "这是一个简单的版本描述",
								"gray_strategy": "p10",
								"status": "INIT",
								"reject_reason": "名称太宽泛",
								"scan_result": "True",
								"gmt_create": "2017-12-12 12:00:00",
								"gmt_apply_audit": "2017-12-12 12:00:00",
								"gmt_online": "2017-12-12 12:00:00",
								"gmt_offline": "2017-12-12 12:00:00",
								"app_desc": "小程序demo的相关示例",
								"gmt_audit_end": "2017-12-12 12:00:00",
								"service_region_type": "LOCATION",
								"service_region_info": [
									{
										"province_code": "310000",
										"province_name": "浙江省",
										"city_code": "310000",
										"city_name": "杭州市",
										"area_code": "311100",
										"area_name": "余杭区"
									}
								],
								"screen_shot_list": [
									"http://image.aaa.alipay.com"
								],
								"app_slogan": "小程序demo简介",
								"memo": "这是一个demo示例",
								"service_phone": "13110101010",
								"service_email": "example@mail.com",
								"mini_app_category_info_list": [
									{
										"first_category_id": "1234",
										"first_category_name": "生活服务",
										"second_category_id": "12344",
										"second_category_name": "汽车服务"
									}
								],
								"package_info_list": [
									{
										"package_name": "基础能力",
										"package_desc": "这是通用能力",
										"doc_url": "http://doc.aaa.alipay.com",
										"status": "valid",
										"package_open_type": "APPLY"
									}
								]
							}
						}`)
	})

	got, err := client.Mini.QueryVersionDetail(context.Background(), &QueryVersionDetailBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.QueryVersionDetail returned unexcepted error: %v", err)
	}
	want := &VersionDetail{
		AppVersion:        "0.0.1",
		AppName:           "小程序demo",
		AppEnglishName:    "demo example",
		AppLogo:           "http://image.aaaa.alipay.com",
		VersionDesc:       "这是一个简单的版本描述",
		GrayStrategy:      "p10",
		Status:            "INIT",
		RejectReason:      "名称太宽泛",
		ScanResult:        "True",
		GmtCreate:         "2017-12-12 12:00:00",
		GmtApplyAudit:     "2017-12-12 12:00:00",
		GmtOnline:         "2017-12-12 12:00:00",
		GmtOffline:        "2017-12-12 12:00:00",
		GmtAuditEnd:       "2017-12-12 12:00:00",
		AppDesc:           "小程序demo的相关示例",
		ServiceRegionType: "LOCATION",
		ServiceRegionInfo: []*RegionInfo{
			{
				ProvinceCode: "310000",
				ProvinceName: "浙江省",
				CityCode:     "310000",
				CityName:     "杭州市",
				AreaCode:     "311100",
				AreaName:     "余杭区",
			},
		},
		ScreenShotList: []string{"http://image.aaa.alipay.com"},
		AppSlogan:      "小程序demo简介",
		Memo:           "这是一个demo示例",
		ServicePhone:   "13110101010",
		ServiceEmail:   "example@mail.com",
		MiniAppCategoryInfoList: []*MiniAppCategoryInfo{
			{
				FirstCategoryID:    "1234",
				FirstCategoryName:  "生活服务",
				SecondCategoryID:   "12344",
				SecondCategoryName: "汽车服务",
			},
		},
		PackageInfoList: []*MiniPackageInfo{
			{
				PackageName:     "基础能力",
				PackageDesc:     "这是通用能力",
				DocURL:          "http://doc.aaa.alipay.com",
				Status:          "valid",
				PackageOpenType: "APPLY",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryVersionDetail got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryVersionDetail_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_detail_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryVersionDetail(context.Background(), &QueryVersionDetailBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.OnlineVersion excepted error")
	}
}

func TestMiniService_RollbackVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_rollback_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.RollbackVersion(context.Background(), &RollbackVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.RollbackVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_RollbackVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_rollback_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.RollbackVersion(context.Background(), &RollbackVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.RollbackVersion excepted error")
	}
}

func TestMiniService_UploadVersion(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_upload_response": {
								"code": "10000",
								"msg": "Success"
							}
						}`)
	})

	err := client.Mini.UploadVersion(context.Background(), &UploadVersionBiz{
		AppVersion:      "0.0.1",
		BundleID:        "com.alipay.alipaywallet",
		TemplateID:      "1",
		TemplateVersion: "0.0.1",
		Ext:             "{\"extEnable\": true, \"extPages\": {\"pages/face/index\": {\"defaultTitle\": \"哈哈哈哈\"}},\"window\": {\"defaultTitle\": \"AI2\"}}",
	})
	if err != nil {
		t.Errorf("Mini.UploadVersion returned unexcepted error: %v", err)
	}
}

func TestMiniService_UploadVersion_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_upload_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	err := client.Mini.UploadVersion(context.Background(), &UploadVersionBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.UploadVersion excepted error")
	}
}

func TestMiniService_QueryVersionList(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_list_query_response": {
								"code": "10000",
								"msg": "Success",
								"app_versions": [
									"0.0.1"
								]
							}
						}`)
	})

	got, err := client.Mini.QueryVersionList(context.Background())
	if err != nil {
		t.Errorf("Mini.QueryVersionList returned unexcepted error: %v", err)
	}
	want := &QueryVersionListResp{
		AppVersions: []string{"0.0.1"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryVersionList got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryVersionList_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_list_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryVersionList(context.Background())
	if err == nil {
		t.Errorf("Mini.QueryVersionList excepted error")
	}
}

func TestMiniService_QueryVersionBuild(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_build_query_response": {
								"code": "10000",
								"msg": "Success",
								"need_rotation": "true",
								"create_status": "6"
							}
						}`)
	})

	got, err := client.Mini.QueryVersionBuild(context.Background(), &QueryVersionBuildBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err != nil {
		t.Errorf("Mini.QueryVersionBuild returned unexcepted error: %v", err)
	}
	want := &QueryVersionBuildResp{
		NeedRotation: "true",
		CreateStatus: "6",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mini.QueryVersionBuild got %+v, want %+v", got, want)
	}
}

func TestMiniService_QueryVersionBuild_failed(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"alipay_open_mini_version_build_query_response": {
							        "code": "20000",
									"msg": "Service Currently Unavailable",
									"sub_code": "isp.unknow-error",
									"sub_msg": "系统繁忙"
							}
						}`)
	})
	_, err := client.Mini.QueryVersionBuild(context.Background(), &QueryVersionBuildBiz{
		AppVersion: "0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	})
	if err == nil {
		t.Errorf("Mini.QueryVersionBuild excepted error")
	}
}

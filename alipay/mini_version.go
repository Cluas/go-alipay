package alipay

import (
	"context"
)

// DeleteVersionBiz 小程序删除版本
type DeleteVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	BundleID   string `json:"bundle_id"`   //小程序投放的端参数，例如投放到支付宝钱包是支付宝端。该参数可选，默认支付宝端，目前仅支持支付宝端，枚举列举：com.alipay.alipaywallet:支付宝端
}

// DeleteVersion 小程序删除版本
func (s *MiniService) DeleteVersion(ctx context.Context, biz *DeleteVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.delete"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// ApplyVersionAuditBiz 小程序提交审核
type ApplyVersionAuditBiz struct {
	LicenseName             string        `json:"license_name"`
	FirstLicensePic         []byte        `json:"first_license_pic"`
	SecondLicensePic        []byte        `json:"second_license_pic"`
	ThirdLicensePic         []byte        `json:"third_license_pic"`
	FourthLicensePic        []byte        `json:"fourth_license_pic"`
	FifthLicensePic         []byte        `json:"fifth_license_pic"`
	LicenseValidDate        string        `json:"license_valid_date"`
	OutDoorPic              []byte        `json:"out_door_pic"`
	AppVersion              string        `json:"app_version"`
	AppName                 string        `json:"app_name"`
	AppEnglishName          string        `json:"app_english_name"`
	AppSlogan               string        `json:"app_slogan"`
	AppLogo                 []byte        `json:"app_logo"`
	AppCategoryIDs          string        `json:"app_category_ids"`
	AppDesc                 string        `json:"app_desc"`
	ServicePhone            string        `json:"service_phone"`
	ServiceEmail            string        `json:"service_email"`
	VersionDesc             string        `json:"version_desc"`
	Memo                    string        `json:"memo"`
	RegionType              string        `json:"region_type"`
	ServiceRegionInfo       []*RegionInfo `json:"service_region_info"`
	FirstScreenShot         []byte        `json:"first_screen_shot"`
	SecondScreenShot        []byte        `json:"second_screen_shot"`
	ThirdScreenShot         []byte        `json:"third_screen_shot"`
	FourthScreenShot        []byte        `json:"fourth_screen_shot"`
	FifthScreenShot         []byte        `json:"fifth_screen_shot"`
	LicenseNo               string        `json:"license_no"`
	FirstSpecialLicensePic  []byte        `json:"first_special_license_pic"`
	SecondSpecialLicensePic []byte        `json:"second_special_license_pic"`
	ThirdSpecialLicensePic  []byte        `json:"third_special_license_pic"`
	TestAccount             string        `json:"test_accout"` // 官方拼写错误
	TestPassword            string        `json:"test_password"`
	TestFileName            []byte        `json:"test_file_name"`
	BundleID                string        `json:"bundle_id"`
}

// RegionInfo 省市区信息，当区域类型为LOCATION时，不能为空
type RegionInfo struct {
	ProvinceCode string `json:"province_code"`
	ProvinceName string `json:"province_name"`
	CityCode     string `json:"city_code"`
	CityName     string `json:"city_name"`
	AreaCode     string `json:"area_code"`
	AreaName     string `json:"area_name"`
}

// ApplyVersionAudit 小程序提交审核
func (s *MiniService) ApplyVersionAudit(ctx context.Context, biz *ApplyVersionAuditBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audit.apply"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CancelVersionAuditBiz 小程序撤销审核
type CancelVersionAuditBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 可不选, 默认撤消正在审核中的版本
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端(com.alipay.alipaywallet:支付宝端)
}

// CancelVersionAudit 小程序撤销审核
func (s *MiniService) CancelVersionAudit(ctx context.Context, biz *CancelVersionAuditBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audit.cancel"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CancelVersionAuditedBiz 小程序退回开发
type CancelVersionAuditedBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号
	BundleID   string `json:"bundle_id"`   //小程序投放的端参数，例如投放到支付宝钱包是支付宝端。该参数可选，默认支付宝端，目前仅支持支付宝端，枚举列举：com.alipay.alipaywallet:支付宝端
}

// CancelVersionAudited 小程序退回开发
func (s *MiniService) CancelVersionAudited(ctx context.Context, biz *CancelVersionAuditedBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.audited.cancel"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OnlineVersionBiz 小程序上架
type OnlineVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// OnlineVersion 小程序上架
func (s *MiniService) OnlineVersion(ctx context.Context, biz *OnlineVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.online"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OfflineVersionBiz 小程序下架
type OfflineVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// OfflineVersion 小程序下架
func (s *MiniService) OfflineVersion(ctx context.Context, biz *OfflineVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.offline"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// RollbackVersionBiz 小程序回滚
type RollbackVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// RollbackVersion 小程序回滚
func (s *MiniService) RollbackVersion(ctx context.Context, biz *RollbackVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.roolback"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// OnlineGrayVersionBiz 小程序灰度上架
type OnlineGrayVersionBiz struct {
	AppVersion   string `json:"app_version"`   //小程序版本号, 必选
	GrayStrategy string `json:"gray_strategy"` //小程序灰度策略值，支持p10，p30，p50, 代表百分之多少的用户
	BundleID     string `json:"bundle_id"`     //端参数，可不选，默认支付宝端
}

// OnlineGrayVersion 小程序灰度上架
func (s *MiniService) OnlineGrayVersion(ctx context.Context, biz *OnlineGrayVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.gray.online"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CancelGrayVersionBiz 小程序结束灰度
type CancelGrayVersionBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// CancelGrayVersion 小程序灰度上架
func (s *MiniService) CancelGrayVersion(ctx context.Context, biz *CancelGrayVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.gray.cancel"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// UploadVersionBiz 小程序基于模板上传版本
type UploadVersionBiz struct {
	AppVersion      string `json:"app_version"`      //小程序版本号, 必选
	BundleID        string `json:"bundle_id"`        //端参数，可不选，默认支付宝端
	TemplateID      string `json:"template_id"`      //模板id
	Ext             string `json:"ext"`              //模板的配置参数
	TemplateVersion string `json:"template_version"` //模板版本号，版本号必须满足 x.y.z, 且均为数字。不传默认使用最新在架模板版本。
}

// UploadVersion 小程序基于模板上传版本
func (s *MiniService) UploadVersion(ctx context.Context, biz *UploadVersionBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.version.upload"
	req, err := s.client.NewRequest("POST", apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// QueryVersionDetailBiz 小程序版本详情查询
type QueryVersionDetailBiz struct {
	AppVersion string `json:"app_version"` //小程序版本号, 必选
	BundleID   string `json:"bundle_id"`   //端参数，可不选，默认支付宝端
}

// MiniAppCategoryInfo 小程序类目
type MiniAppCategoryInfo struct {
	FirstCategoryID    string `json:"first_category_id"`
	FirstCategoryName  string `json:"first_category_name"`
	SecondCategoryID   string `json:"second_category_id"`
	SecondCategoryName string `json:"second_category_name"`
}

// MiniPackageInfo 小程序功能包
type MiniPackageInfo struct {
	PackageName     string `json:"package_name"`
	PackageDesc     string `json:"package_desc"`
	DocUrl          string `json:"doc_url"`
	Status          string `json:"status"`
	PackageOpenType string `json:"package_open_type"`
}

// VersionDetail 小程序版本详情
type VersionDetail struct {
	AppVersion              string                 `json:"app_version"`
	AppName                 string                 `json:"app_name"`
	AppEnglishName          string                 `json:"app_english_name"`
	AppLogo                 string                 `json:"app_logo"`
	VersionDesc             string                 `json:"version_desc"`
	GrayStrategy            string                 `json:"gray_strategy"`
	Status                  string                 `json:"status"`
	RejectReason            string                 `json:"reject_reason"`
	ScanResult              string                 `json:"scan_result"`
	GmtCreate               string                 `json:"gmt_create"`
	GmtApplyAudit           string                 `json:"gmt_apply_audit"`
	GmtOnline               string                 `json:"gmt_online"`
	GmtOffline              string                 `json:"gmt_offline"`
	GmtAuditEnd             string                 `json:"gmt_audit_end"`
	AppDesc                 string                 `json:"app_desc"`
	ServiceRegionType       string                 `json:"service_region_type"`
	ServiceRegionInfo       []*RegionInfo          `json:"service_region_info"`
	ScreenShotList          []string               `json:"screen_shot_list"`
	AppSlogan               string                 `json:"app_slogan"`
	Memo                    string                 `json:"memo"`
	ServicePhone            string                 `json:"service_phone"`
	ServiceEmail            string                 `json:"service_email"`
	MiniAppCategoryInfoList []*MiniAppCategoryInfo `json:"mini_app_category_info_list"`
	PackageInfoList         []*MiniPackageInfo     `json:"package_info_list"`
}

// QueryVersionDetail 小程序版本详情查询
func (s *MiniService) QueryVersionDetail(ctx context.Context, biz *QueryVersionDetailBiz, opts ...ValueOptions) (*VersionDetail, error) {
	apiMethod := "alipay.open.mini.version.detail.query"
	req, err := s.client.NewRequest("GET", apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	versionDetail := new(VersionDetail)
	_, err = s.client.Do(ctx, req, versionDetail)
	if err != nil {
		return nil, err
	}
	return versionDetail, nil
}

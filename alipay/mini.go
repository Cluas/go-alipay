package alipay

import "context"

// MiniService 小程序服务
//
// Docs: https://opendocs.alipay.com/apis/api_49
type MiniService service

// BaseInfo 小程序基础信息
type BaseInfo struct {
	AppName        string   `json:"app_name"`
	AppEnglishName string   `json:"app_english_name"`
	AppSlogan      string   `json:"app_slogan"`
	AppLogo        string   `json:"app_logo"`
	CategoryNames  string   `json:"category_names"`
	AppDesc        string   `json:"app_desc"`
	ServicePhone   string   `json:"service_phone"`
	ServiceEmail   string   `json:"service_email"`
	SafeDomains    []string `json:"safe_domains,omitempty"`
	PackageNames   []string `json:"package_names,omitempty"`
}

// QueryMiniBaseInfo 查询小程序基础信息
func (s *MiniService) QueryMiniBaseInfo(ctx context.Context, opts ...ValueOptions) (*BaseInfo, error) {
	apiMethod := "alipay.open.mini.baseinfo.query"
	req, err := s.client.NewRequest("GET", apiMethod, nil, opts...)
	if err != nil {
		return nil, err
	}
	baseInfo := new(BaseInfo)
	_, err = s.client.Do(ctx, req, baseInfo)
	if err != nil {
		return nil, err
	}
	return baseInfo, nil
}

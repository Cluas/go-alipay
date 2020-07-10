package alipay

import "context"

// QueryTemplateUsageBiz 查询使用模板的小程序列表
type QueryTemplateUsageBiz struct {
	TemplateID      string `json:"template_id"`                // 模板id
	PageNum         int    `json:"page_num,omitempty"`         // 查询的页数，默认第一页
	PageSize        int    `json:"page_size,omitempty"`        // 每页的数量，最多查询50个，默认查询10个
	TemplateVersion string `json:"template_version,omitempty"` // 模板小程序的版本号
	// 小程序客户端类型，默认为支付宝端。
	// 支付宝端：com.alipay.alipaywallet,
	// DINGDING端：com.alibaba.android.rimet,
	// 高德端:com.amap.app,
	// 天猫精灵端:com.alibaba.ailabs.genie.webapps,
	// 支付宝IOT:com.alipay.iot.xpaas
	BundleID string `json:"bundle_id,omitempty"`
}

// QueryTemplateUsageResp 查询使用模板的小程序列表resp
type QueryTemplateUsageResp struct {
	TemplateUsageInfoList []*TemplateUsageInfo `json:"template_usage_info_list"` // 模板使用信息
}

// TemplateUsageInfo 小程序信息
type TemplateUsageInfo struct {
	MiniAppID  string `json:"mini_app_id"` // 商家小程序appId
	AppVersion string `json:"app_version"` // 商家小程序版本号
}

// QueryTemplateUsage 查询使用模板的小程序列表
func (s *MiniService) QueryTemplateUsage(ctx context.Context, biz *QueryTemplateUsageBiz, opts ...ValueOptions) (*QueryTemplateUsageResp, error) {
	apiMethod := "alipay.open.mini.template.usage.query"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryTemplateUsageResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

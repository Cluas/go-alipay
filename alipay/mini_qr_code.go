package alipay

import "context"

// BindQrCodeBiz 关联普通二维码biz
type BindQrCodeBiz struct {
	RouteURL string `json:"route_url"` // 二维码域名，须通过ICP备案验证，支持http、https、ftp开头的链接
	//匹配规则，仅支持EXACT（精确匹配）、FUZZY（模糊匹配）两个值。
	//精确匹配：根据填写的二维码地址精确匹配，地址完全一致时才能唤起小程序（如：配置二维码地址为https://www.alipay.com/my?id=123，当用户扫这个地址的二维码可唤起小程序）。
	//模糊匹配：根据填写的二维码地址模糊匹配，只要地址前缀匹配即可唤起小程序（如：配置二维码地址为https://www.alipay.com/my/，当用户扫的二维码地址为https://www.alipay.com/my/id=123,可唤起小程序）。
	Mode            string `json:"mode"`
	PageRedirection string `json:"page_redirection"` // 小程序功能页，配置扫描二维码后打开的小程序功能页面路径
}

// BindQrCodeResp 关联普通二维码resp
type BindQrCodeResp struct {
	RouteGroup string `json:"route_group"` // 路由规则组，用于唯一标记一条路由规则
}

// BindQrCode 关联普通二维码
func (s *MiniService) BindQrCode(ctx context.Context, biz *BindQrCodeBiz, opts ...ValueOptions) (*BindQrCodeResp, error) {
	apiMethod := "alipay.open.mini.qrcode.bind"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(BindQrCodeResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UnbindQrCodeBiz 删除已关联普通二维码biz
type UnbindQrCodeBiz struct {
	RouteGroup string `json:"route_group"` // 路由规则组，用于唯一标记一条路由规则
}

// UnbindQrCode 删除已关联普通二维码
func (s *MiniService) UnbindQrCode(ctx context.Context, biz *UnbindQrCodeBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.qrcode.unbind"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	resp := new(BindQrCodeResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return err
	}
	return nil
}

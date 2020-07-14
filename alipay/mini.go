package alipay

import (
	"context"
	"io"
)

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

// QueryBaseInfo 查询小程序基础信息
func (s *MiniService) QueryBaseInfo(ctx context.Context, opts ...ValueOptions) (*BaseInfo, error) {
	apiMethod := "alipay.open.mini.baseinfo.query"
	req, err := s.client.NewRequest(apiMethod, nil, opts...)
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

// ModifyBaseInfoBiz 小程序修改基础信息
type ModifyBaseInfoBiz struct {
	AppName         string `json:"app_name,omitempty"`          // 小程序应用名称
	AppEnglishName  string `json:"app_english_name,omitempty"`  // 小程序应用英文名称
	AppSlogan       string `json:"app_slogan,omitempty"`        // 小程序应用简介，一句话描述小程序功能
	AppLogo         *File  `json:"app_logo,omitempty"`          // 小程序应用logo图标，图片格式必须为：png、jpeg、jpg，建议上传像素为180*180
	AppCategoryIDs  string `json:"app_category_ids,omitempty"`  // 11_12;12_13。小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考https://docs.alipay.com/isv/10325
	AppDesc         string `json:"app_desc,omitempty"`          // 小程序应用描述，20-200个字
	ServicePhone    string `json:"service_phone,omitempty"`     // 小程序客服电话
	ServiceEmail    string `json:"service_email,omitempty"`     // 小程序客服邮箱
	MiniCategoryIDs string `json:"mini_category_ids,omitempty"` // 新小程序前台类目，一级与二级、三级用下划线隔开，最多可以选四个类目，类目之间;隔开。使用后不再读取app_category_ids值，老前台类目将废弃
}

type MultiRender interface {
	Params() map[string]string
	MultipartParams() map[string]io.Reader
}

func (b *ModifyBaseInfoBiz) Params() map[string]string {
	params := make(map[string]string)
	if b.AppName != "" {
		params["app_name"] = b.AppName
	}
	if b.AppEnglishName != "" {
		params["app_english_name"] = b.AppEnglishName
	}
	if b.AppSlogan != "" {
		params["app_slogan"] = b.AppSlogan
	}
	if b.AppCategoryIDs != "" {
		params["app_category_ids"] = b.AppCategoryIDs
	}
	if b.AppDesc != "" {
		params["app_desc"] = b.AppDesc
	}
	if b.ServicePhone != "" {
		params["service_phone"] = b.ServicePhone
	}
	if b.ServiceEmail != "" {
		params["service_email"] = b.ServiceEmail
	}
	if b.MiniCategoryIDs != "" {
		params["mini_category_ids"] = b.MiniCategoryIDs
	}
	return params
}
func (b *ModifyBaseInfoBiz) MultipartParams() map[string]io.Reader {
	params := make(map[string]io.Reader)
	if b.AppLogo != nil {
		params["app_logo"] = b.AppLogo
	}
	return params
}

// ModifyBaseInfo 小程序修改基础信息
func (s *MiniService) ModifyBaseInfo(ctx context.Context, biz *ModifyBaseInfoBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.baseinfo.modify"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreateSafeDomainBiz 小程序添加域白名单
type CreateSafeDomainBiz struct {
	SafeDomain string `json:"safe_domain"` // httpRequest域白名单 示例值：example.com 一次只支持设置一个域名
}

// CreateSafeDomain 小程序添加域白名单
func (s *MiniService) CreateSafeDomain(ctx context.Context, biz *CreateSafeDomainBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.mini.safedomain.create"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}

// DetectRiskContentBiz 小程序风险内容检测服务biz
type DetectRiskContentBiz struct {
	Content string `json:"content"` // 最大长度2000 需要识别的文本，不要包含特殊字符以及双引号等可能引起json格式化错误问题的字符.

}

// DetectRiskContentResp 小程序风险内容检测服务resp
type DetectRiskContentResp struct {
	Action   string   `json:"action"`    //表示处理结果，REJECTED表示拦截，PASSED表示放过。
	Keywords []string `json:"keywords"`  // 命中的关键字
	UniqueID string   `json:"unique_id"` // 业务唯一识别码，可用来对应异步识别结果
}

// DetectRiskContent 小程序风险内容检测服务
func (s *MiniService) DetectRiskContent(ctx context.Context, biz *DetectRiskContentBiz, opts ...ValueOptions) (*DetectRiskContentResp, error) {
	apiMethod := "alipay.security.risk.content.detect"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(DetectRiskContentResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// QueryTinyAppExistBiz 查询是否创建过小程序biz
type QueryTinyAppExistBiz struct {
	PID string `json:"pid"` //支付宝账号ID
}

// QueryTinyAppExistResp 查询是否创建过小程序resp
type QueryTinyAppExistResp struct {
	ExistMini string `json:"exist_mini"` // 是否是小程序开发者
}

// QueryTinyAppExist 查询是否创建过小程序
func (s *MiniService) QueryTinyAppExist(ctx context.Context, biz *QueryTinyAppExistBiz, opts ...ValueOptions) (*QueryTinyAppExistResp, error) {
	apiMethod := "alipay.open.mini.tinyapp.exist.query"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryTinyAppExistResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// QueryCategoryBiz 小程序类目树查询
type QueryCategoryBiz struct {
	IsFilter bool `json:"is_filter,omitempty"` // 是否需要过滤不可用类目
}

// MiniAppCategory 小程序类别
type MiniAppCategory struct {
	CategoryID         string `json:"category_id"`
	CategoryName       string `json:"category_name"`
	ParentCategoryID   string `json:"parent_category_id"`
	HasChild           bool   `json:"has_child"`
	NeedLicense        bool   `json:"need_license"`
	NeedOutDoorPic     bool   `json:"need_out_door_pic"`
	NeedSpecialLicense bool   `json:"need_special_license"`
}

// QueryCategoryResp 小程序类目树查询resp
type QueryCategoryResp struct {
	MiniCategoryList []*MiniAppCategory `json:"mini_category_list"`
	CategoryList     []*MiniAppCategory `json:"category_list"`
}

// QueryCategory 小程序类目树查询
func (s *MiniService) QueryCategory(ctx context.Context, biz *QueryCategoryBiz, opts ...ValueOptions) (*QueryCategoryResp, error) {
	apiMethod := "alipay.open.mini.category.query"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryCategoryResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// CertifyIndividualBusinessBiz 个人账户升级为个体工商户
type CertifyIndividualBusinessBiz struct {
	LiceseNo  string `json:"license_no"`  // 营业执照
	LicesePic string `json:"license_pic"` //	营业执照图片的Base64编码字符串，图片大小不能超过2M
}

// CertifyIndividualBusinessResp 个人账户升级为个体工商户resp
type CertifyIndividualBusinessResp struct {
	CertifyResult bool `json:"certify_result"` // 个体工商户认证结果，true代表认证成功，false代表认证失败
}

// CertifyIndividualBusiness 个人账户升级为个体工商户
func (s *MiniService) CertifyIndividualBusiness(ctx context.Context, biz *CertifyIndividualBusinessBiz, opts ...ValueOptions) (*CertifyIndividualBusinessResp, error) {
	apiMethod := "alipay.open.mini.individual.business.certify"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(CertifyIndividualBusinessResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// SyncContentBiz 小程序内容接入biz
type SyncContentBiz struct {
	ContentType string `json:"content_type"` // 内容类型，例如门店、商品等
	Operation   string `json:"operation"`    // 内容类型下的具体操作，比如门店类型下，小程序批量绑定门店。可参考具体内容接入文档中的详细说明。
	ContentData string `json:"content_data"` // 具体的内容数据，采用json格式，不同类型不同操作数据不同。可参考具体内容接入文档中的详细说明。
	ExtendInfo  string `json:"extend_info"`  // 扩展信息，json格式。可参考具体内容接入文档中的详细说明。
}

// SyncContentResp 小程序内容接入resp
type SyncContentResp struct {
	ResultData string `json:"result_data"` // 结果数据，json格式，可参考具体内容接入文档中的详细说明。
}

// SyncContent 小程序内容接入
func (s *MiniService) SyncContent(ctx context.Context, biz *SyncContentBiz, opts ...ValueOptions) (*SyncContentResp, error) {
	apiMethod := "alipay.open.mini.content.sync"
	req, err := s.client.NewRequest(apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(SyncContentResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

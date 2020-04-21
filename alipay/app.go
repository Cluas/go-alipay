package alipay

import (
	"context"
)

// AppService 应用服务
// Docs: https://opendocs.alipay.com/apis/api_49
type AppService service

// CreateAppMemberBiz 应用添加成员
type CreateAppMemberBiz struct {
	LogonID string `json:"logon_id"` // 支付宝登录账号ID
	Role    string `json:"role"`     //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

// CreateMember 应用添加成员，目前只支持小程序类型的应用使用
func (s *AppService) CreateMember(ctx context.Context, biz *CreateAppMemberBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.app.members.create"
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

// DeleteMemberBiz 应用删除成员
type DeleteMemberBiz struct {
	LogonID string `json:"logon_id"` // 支付宝登录账号ID
	Role    string `json:"role"`     //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

// DeleteMember 应用删除成员，目前只支持小程序类型的应用使用
func (s *AppService) DeleteMember(ctx context.Context, biz *DeleteMemberBiz, opts ...ValueOptions) error {
	apiMethod := "alipay.open.app.members.delete"
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

// QueryAppMembersBiz 应用查询成员列表
type QueryAppMembersBiz struct {
	Role string `json:"role"` //成员的角色类型，DEVELOPER-开发者，EXPERIENCER-体验者
}

// AppMemberInfo 小程序成员模型
type AppMemberInfo struct {
	UserID    string `json:"user_id"`
	NickName  string `json:"nick_name"`
	Portrait  string `json:"portrait"`
	Status    string `json:"status"`
	GmtJoin   string `json:"gmt_join"`
	LogonID   string `json:"logon_id"`
	GmtInvite string `json:"gmt_invite"`
	Role      string `json:"role"`
}

// QueryAppMembersResp 成员列表
type QueryAppMembersResp struct {
	AppMemberInfoList []*AppMemberInfo `json:"app_member_info_list"`
}

// QueryAppMembers 应用查询成员列表，目前只支持小程序类型的应用
func (s *AppService) QueryAppMembers(ctx context.Context, biz *QueryAppMembersBiz, opts ...ValueOptions) (*QueryAppMembersResp, error) {
	apiMethod := "alipay.open.app.members.query"
	req, err := s.client.NewRequest("GET", apiMethod, biz, opts...)
	if err != nil {
		return nil, err
	}
	resp := new(QueryAppMembersResp)
	_, err = s.client.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

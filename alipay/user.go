package alipay

import (
	"context"
)

// UserService 会员API
//
//docs: https://opendocs.alipay.com/apis/api_2
type UserService service

// UserInfoShare 支付宝会员授权信息
type UserInfoShare struct {
	Avatar             string `json:"avatar,omitempty"`
	City               string `json:"city,omitempty"`
	Gender             string `json:"gender,omitempty"`
	IsCertified        string `json:"is_certified,omitempty"`
	IsStudentCertified string `json:"is_student_certified,omitempty"`
	NickName           string `json:"nick_name,omitempty"`
	Province           string `json:"province,omitempty"`
	UserID             string `json:"user_id,omitempty"`
	UserStatus         string `json:"user_status,omitempty"`
	UserType           string `json:"user_type,omitempty"`
}

// InfoShare 支付宝会员授权信息查询接口
func (s UserService) InfoShare(ctx context.Context, authToken string) (*UserInfoShare, error) {
	apiMethod := "alipay.user.info.share"
	req, err := s.client.NewRequest("POST", apiMethod, nil, AuthToken(authToken))
	if err != nil {
		return nil, err
	}
	userInfoShare := new(UserInfoShare)
	_, err = s.client.Do(ctx, req, userInfoShare)
	if err != nil {
		return nil, err
	}

	return userInfoShare, nil
}

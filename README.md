# go-alipay #
Go 支付宝小程序SDK

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Cluas/go-alipay)
[![GoReport](https://goreportcard.com/badge/github.com/Cluas/go-alipay)](https://goreportcard.com/report/github.com/Cluas/go-alipay)
[![Test Status](https://github.com/Cluas/go-alipay/workflows/tests/badge.svg)](https://github.com/Cluas/go-alipay/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/Cluas/go-alipay/branch/master/graph/badge.svg)](https://codecov.io/gh/Cluas/go-alipay)


微信开放平台小程序第三方接口支持

### 简单示例
```go
package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/Cluas/go-alipay/alipay"
)

func main() {
	// 小程序代码上架示例
	encodedKey, _ := base64.StdEncoding.DecodeString("your_private_key")
	privateKey, _ := x509.ParsePKCS1PrivateKey(encodedKey)
	publicKey, _ := base64.StdEncoding.DecodeString("your_public_key")
	pub, _ := x509.ParsePKIXPublicKey(publicKey)
	client := alipay.NewClient(nil, privateKey, pub.(*rsa.PublicKey))

	biz := alipay.OnlineVersionBiz{
		AppVersion: "v0.0.1",
		BundleID:   "com.alipay.alipaywallet",
	}
	if err := client.Mini.OnlineVersion(context.Background(), &biz); err != nil {
		fmt.Printf("支付宝小程序代码上架失败: %s", err)
	}
	// 如果是第三方代开发
	// client.Mini.OnlineVersion(context.Background(), &biz, alipay.AppAuthToken(token))

}
```
### 目前已对接的接口
- [x] alipay.open.app.members.create 创建应用成员
- [x] alipay.open.app.members.delete 删除应用成员
- [x] alipay.open.app.members.query  查询应用成员列表

- [x] alipay.open.mini.experience.create 生成小程序体验版
- [x] alipay.open.mini.experience.query  小程序体验版状态查询
- [x] alipay.open.mini.experience.cancel 小程序取消体验版

- [x] alipay.open.mini.version.delete 小程序删除版本
- [x] alipay.open.mini.version.audit.apply 小程序提交审核
- [x] alipay.open.mini.version.audit.cancel 小程序撤销审核
- [x] alipay.open.mini.version.audited.cancel 小程序退回开发
- [x] alipay.open.mini.version.online 小程序上架
- [x] alipay.open.mini.version.offline 小程序下架
- [x] alipay.open.mini.version.roolback 小程序回滚
- [x] alipay.open.mini.version.gray.online 小程序灰度上架
- [x] alipay.open.mini.version.gray.cancel 小程序结束灰度
- [x] alipay.open.mini.version.upload 小程序基于模板上传版本
- [x] alipay.open.mini.version.detail.query 小程序版本详情查询



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
	client := alipay.NewClient(nil, privateKey, pub.(*rsa.PublicKey), alipay.AppID("your_app_id"))

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
### 支持所有已公布的小程序API
文档地址: https://opendocs.alipay.com/apis/api_49/



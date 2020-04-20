package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/Cluas/go-alipay/alipay"
)

func fetchUserShareInfo(authToken string) (*alipay.UserInfoShare, error) {
	key := "MIICWwIBAAKBgQC4UcQm06Kz9OH8Q6l2wxSOt9BdObuuC1hJQrQNbkqHU7SM1aI4g156fbAoaEZdb7k2bQSyf6PNWYNS+cl9LPsggbYZ1ZapbqgEt39N4sMKOPUEwMco4P9ZQL6C2+1YfqUc4zZKCqiocgXy0tuV3kKWYleOM/Y+J/2PfAUtKF2p3wIDAQABAoGANAQnRgnNzdla+TUjGvf80jX/oH+NfpWHCc3AQFYSxFQUDPaxPB+exxS3ZP/gc7f23ewwOiuZT3dmf0Es4p2SFOQypacVFyzi4Dj/cvJGxze8Ek047jS5wc6tZiQHjPcmPB0i2/wAJt9ThINdBnSzKrjRhfWy1aRay7fNk1BTmAECQQDvuYRR9yGDifc4T8at2xvUbPKavDFNUx2SNq233A2+DESFa9w3ZirVjiKzLR4/d60Gt/n9j5PssP4syECrGIwBAkEAxNVKNLO44+e8otUPc//s+Uhwzp3ASNT2JkVv4kFO+mkaGErkGnySWmWSbvjziK3TFkYOAGFUzH2+6MPETv+13wJAJKIl/VyVq4NG2z0dsG2+V/z6Kfk+U4GzECf47hLbqsI3KmhsM68SNqZM2TK435wLPe6Zbk0lntMBVJiZgUv0AQJAO/BLgZL9CYHHArro0sUrb5nsqC6HoGYhcvQQJxEGMOESjjU4Ewy+MILfvaVX29Y7AnxgxSLehMsB+LWssPXTdwJAJjRaoDllB2eO5wXAuKZNqYzpI6T3tK7tNG51SDlwkv3WMzuihwkv/tys/pWcFtwJFimbL34e/4dpWB1sHxtA1Q=="
	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB"
	encodedKey, _ := base64.StdEncoding.DecodeString(key)
	privateKey, _ := x509.ParsePKCS1PrivateKey(encodedKey)
	public, _ := base64.StdEncoding.DecodeString(publicKey)
	pub, _ := x509.ParsePKIXPublicKey(public)
	client := alipay.NewClient(nil, privateKey, pub.(*rsa.PublicKey), alipay.AppID("2016020101133396"), alipay.SignType("RSA"), alipay.Charset("utf-8"), alipay.Format("JSON"))
	userInfo, err := client.User.InfoShare(context.Background(), authToken)
	return userInfo, err
}

func main() {
	authToken := "authusrB39ed542c73964f9bbb6a77d4383f8E76"

	user, err := fetchUserShareInfo(authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%v", user)
}

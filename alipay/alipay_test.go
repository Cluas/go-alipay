package alipay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, tearDown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	client = NewClient(nil, nil, nil)
	client.BaseURL, _ = url.Parse(server.URL)
	serverURL = server.URL
	tearDown = server.Close
	return
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestClient_CheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(strings.NewReader(`
{
    "alipay_user_info_share_response": {
        "code": "40002",
        "msg": "Invalid Arguments",
        "sub_code": "isv.invalid-timestamp",
        "sub_msg": "非法的时间戳参数"
    }
}
`)),
	}
	client := NewClient(nil, nil, nil)
	err := client.CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Errorf("Expected error response.")
	}

	want := &ErrorResponse{
		Response: res,
		Msg:      "Invalid Arguments",
		Code:     "40002",
		SubCode:  "isv.invalid-timestamp",
		SubMsg:   "非法的时间戳参数",
	}
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Error = %#v, want %#v", err, want)
	}
}

func TestDo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
							"foo_response": {"A":"a", "code": "10000"}
							}`)
	})

	req, _ := client.NewRequest("GET", "foo", nil)
	body := new(foo)
	client.Do(context.Background(), req, body)

	want := &foo{A: "a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}
func TestDo_nilContext(t *testing.T) {
	client, _, _, tearDown := setup()
	defer tearDown()

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(nil, req, nil)

	if !reflect.DeepEqual(err, errors.New("context must be non-nil")) {
		t.Errorf("Expected context must be non-nil error")
	}
}
func TestNewClint(t *testing.T) {
	c := NewClient(nil, nil, nil)
	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL.String(), defaultBaseURL)
	}
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", c.UserAgent, userAgent)
	}

	cNew := NewClient(nil, nil, nil)
	if c.client == cNew.client {
		t.Error("NewClient returned same http.Client, but they should differ.")
	}

}

func TestClient_Sign(t *testing.T) {
	key := `MIICWwIBAAKBgQC4UcQm06Kz9OH8Q6l2wxSOt9BdObuuC1hJQrQNbkqHU7SM1aI4g156fbAoaEZdb7k2bQSyf6PNWYNS+cl9LPsggbYZ1ZapbqgEt39N4sMKOPUEwMco4P9ZQL6C2+1YfqUc4zZKCqiocgXy0tuV3kKWYleOM/Y+J/2PfAUtKF2p3wIDAQABAoGANAQnRgnNzdla+TUjGvf80jX/oH+NfpWHCc3AQFYSxFQUDPaxPB+exxS3ZP/gc7f23ewwOiuZT3dmf0Es4p2SFOQypacVFyzi4Dj/cvJGxze8Ek047jS5wc6tZiQHjPcmPB0i2/wAJt9ThINdBnSzKrjRhfWy1aRay7fNk1BTmAECQQDvuYRR9yGDifc4T8at2xvUbPKavDFNUx2SNq233A2+DESFa9w3ZirVjiKzLR4/d60Gt/n9j5PssP4syECrGIwBAkEAxNVKNLO44+e8otUPc//s+Uhwzp3ASNT2JkVv4kFO+mkaGErkGnySWmWSbvjziK3TFkYOAGFUzH2+6MPETv+13wJAJKIl/VyVq4NG2z0dsG2+V/z6Kfk+U4GzECf47hLbqsI3KmhsM68SNqZM2TK435wLPe6Zbk0lntMBVJiZgUv0AQJAO/BLgZL9CYHHArro0sUrb5nsqC6HoGYhcvQQJxEGMOESjjU4Ewy+MILfvaVX29Y7AnxgxSLehMsB+LWssPXTdwJAJjRaoDllB2eO5wXAuKZNqYzpI6T3tK7tNG51SDlwkv3WMzuihwkv/tys/pWcFtwJFimbL34e/4dpWB1sHxtA1Q==`
	want := "mUZy1ZICyburtoau/gk4VpD50uutt46d0JJoJOB/WRaia0UeWcf6ERwFJmxo+Urf/z410YCcmlX84SuNDzP+erqKzewpWNZyEsYs8QxonZqV1lzy5Yo3NJRKG0Xejp0Zs2Wek96Egg2VeGABmEfsMBx3IoQzKqU3ZeSw/JCPaeA="

	encodedKey, _ := base64.StdEncoding.DecodeString(key)
	privateKey, _ := x509.ParsePKCS1PrivateKey(encodedKey)

	c := NewClient(nil, privateKey, nil, AppID("2016091100484533"), SignType("RSA"), Charset("UTF-8"))
	values := url.Values{}
	values.Set("method", "alipay.trade.create")
	values.Set("app_id", c.o.AppID)
	values.Set("sign_type", c.o.SignType)
	values.Set("biz_content", `{"out_trade_no":"202004191441122314312","total_amount":"88.88","subject":"iPhone+Xs+Max+256G","buyer_id":"2088102175953034","goods_detail":[],"seller_id":"2088102175107499"}`)
	values.Set("charset", c.o.Charset)
	values.Set("timestamp", "2020-04-19 14:41:12")
	values.Set("version", "1.0")

	got, err := c.Sign(values)
	if err != nil {
		t.Errorf("Client.Sign returned unexcept err: %v", err)
	}
	if got != want {
		t.Errorf("Client.Sign got %v, want %v", got, want)
	}

}

func TestClient_VerifySign(t *testing.T) {
	publicKey := `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB`
	sign := `SnAwVaV8YCy3CYiJymlP+kEwxru6uzS4Ym5VDrUZPWgkqnElkq2zRnyigU2k9hXJHdYXZ2JPIyyGx7jufDn4I6IJrTqqhCOC4q254acls9KM/dXjBYYB0QYORG7DqCYXgEz5TGXHkE56gVac4PbqIt/9wnY9o6LBTpfc2fh96+I=`
	public, _ := base64.StdEncoding.DecodeString(publicKey)
	pub, _ := x509.ParsePKIXPublicKey(public)

	c := NewClient(nil, nil, pub.(*rsa.PublicKey), SignType("RSA"))
	err := c.VerifySign([]byte(`{"code":"10000","msg":"Success","avatar":"https:\/\/tfs.alipayobjects.com\/images\/partner\/T15ABtXk8bXXXXXXXX","city":"杭州市","gender":"m","is_certified":"T","is_student_certified":"F","nick_name":"WinWen","province":"浙江省","user_id":"2088912161915762","user_status":"T","user_type":"2"}`), sign)
	if err != nil {
		t.Errorf("Client.SiVerifySigngn returned unexcept err: %v", err)
	}
}

func TestAppAuthToken(t *testing.T) {
	v := url.Values{}

	setter := AppAuthToken("test")
	setter(v)
	got := v.Get("app_auth_token")
	want := "test"

	if got != want {
		t.Errorf("AppAuthToken got %v, want %v", got, want)
	}
}

func TestAppID(t *testing.T) {
	o := Options{}

	setter := AppID("test")
	setter(&o)
	got := o.AppID
	want := "test"

	if got != want {
		t.Errorf("AppID got %v, want %v", got, want)
	}
}

func TestAuthToken(t *testing.T) {
	v := url.Values{}

	setter := AuthToken("test")
	setter(v)
	got := v.Get("auth_token")
	want := "test"

	if got != want {
		t.Errorf("AppAuthToken got %v, want %v", got, want)
	}
}

func TestCharset(t *testing.T) {
	o := Options{}

	setter := Charset("utf-8")
	setter(&o)
	got := o.Charset
	want := "utf-8"

	if got != want {
		t.Errorf("Charset got %v, want %v", got, want)
	}
}

func TestFormat(t *testing.T) {
	o := Options{}

	setter := Format("JSON")
	setter(&o)
	got := o.Format
	want := "JSON"

	if got != want {
		t.Errorf("Charset got %v, want %v", got, want)
	}
}

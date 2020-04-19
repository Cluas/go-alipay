package alipay

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL = "https://openapi.alipay.com/gateway.do"
	userAgent      = "go-alipay"
)

type Options struct {
	AppID        string
	PrivateKey   string
	PublicKey    string
	Format       string
	Charset      string
	SignType     string
	Version      string
	NotifyURL    string
	AppAuthToken string
}

type Option func(*Options)

func AppID(appID string) Option {
	return func(o *Options) {
		o.AppID = appID
	}
}

func PrivateKey(privateKey string) Option {
	return func(o *Options) {
		o.PrivateKey = privateKey
	}
}

func PublicKey(publicKey string) Option {
	return func(o *Options) {
		o.PublicKey = publicKey
	}
}

func Format(format string) Option {
	return func(o *Options) {
		o.Format = format
	}
}

func Charset(charset string) Option {
	return func(o *Options) {
		o.Charset = charset
	}
}

func SignType(signType string) Option {
	return func(o *Options) {
		o.SignType = signType
	}
}

// A Client manages communication with the Alipay API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Alipay API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	o *Options

	// User agent used when communicating with the Alipay API.
	UserAgent string
}

// NewClient returns a new Alipay API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, setters ...Option) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	o := &Options{
		Format:   "JSON",
		Charset:  "utf8",
		SignType: "RSA2",
	}
	for _, setter := range setters {
		setter(o)
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, o: o}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, apiMethod string, content interface{}) (*http.Request, error) {
	var buf *bytes.Buffer
	if content == nil {
		return nil, errors.New("content is required")
	}
	buf = &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(content)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Set("app_id", c.o.AppID)
	v.Set("method", apiMethod)
	v.Set("format", c.o.Format)
	v.Set("charset", c.o.Charset)
	v.Set("sign_type", c.o.SignType)
	v.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	v.Set("version", "1.0")
	v.Set("biz_content", buf.String())
	if c.o.AppAuthToken != "" {
		v.Set("app_auth_token", c.o.AppAuthToken)
	}
	if c.o.NotifyURL != "" {
		v.Set("notify_url", c.o.NotifyURL)
	}
	sign, err := c.Sign(v.Encode())
	if err != nil {
		return nil, err
	}
	v.Set("sign", sign)

	req, err := http.NewRequest(method, c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	switch method {
	case http.MethodPost:
		req.Form = v
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case http.MethodGet:
		req.URL.RawQuery = v.Encode()
	default:
		return nil, errors.New("unsupported method")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Sign(urlStr string) (string, error) {

	signType := crypto.SHA256
	if c.o.SignType == "RSA" {
		signType = crypto.SHA1
	}
	h := crypto.Hash.New(signType)
	h.Write([]byte(urlStr))
	encodedKey, err := base64.StdEncoding.DecodeString(c.o.PrivateKey)
	if err != nil {
		return "", err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(encodedKey)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, signType, h.Sum(nil))
	if err != nil {
		return "", err
	}

	sign := base64.StdEncoding.EncodeToString(signature)

	return sign, nil

}

// ErrorResponse is common error response.
type ErrorResponse struct {
	*http.Response
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v, %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Msg, r.Code, r.SubCode, r.SubMsg)
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

// Response is a Alipay API response.
type Response struct {
	*http.Response
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = withContext(ctx, req)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	var apiMethod string
	switch req.Method {
	case http.MethodGet:
		apiMethod = req.URL.Query().Get("method")
	default:
		apiMethod = req.Form.Get("method")
	}
	response := &Response{resp}

	err = c.CheckResponse(resp, apiMethod)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

func (c *Client) CheckResponse(r *http.Response, apiMethod string) error {
	var (
		respStr, sign string
		respContent   []byte
	)
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		dataStr := string(data)
		respNode := strings.Replace(apiMethod, ".", "_", -1) + "_response"
		respIdx := strings.Index(dataStr, respNode)
		signIdx := strings.Index(dataStr, "sign")

		if signIdx > respIdx {
			respStr = dataStr[respIdx+len(respNode)+2 : signIdx-2]
			sign = dataStr[signIdx+6:]
		} else {
			respStr = dataStr[respIdx+len(respNode)+2:]
			sign = dataStr[signIdx+6 : respIdx-2]
		}
		respContent = []byte(respStr)
		json.Unmarshal(respContent, errorResponse)
		if err := c.VerifySign(respContent, sign); err != nil {
			return errors.New("invalid signature")
		}
	}

	if errorResponse.Code == "10000" {
		buf := bytes.NewBuffer(respContent)
		r.Body = ioutil.NopCloser(buf)
		return nil
	}
	return errorResponse
}

func (c *Client) VerifySign(content []byte, sign string) error {
	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(c.o.PublicKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}

	signType := crypto.SHA256
	if c.o.SignType == "RSA" {
		signType = crypto.SHA1
	}
	h := crypto.Hash.New(signType)
	h.Write(content)
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), signType, h.Sum(nil), signData)
}

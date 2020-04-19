package alipay

import (
	"net/url"
	"testing"
)

func TestNewClint(t *testing.T) {
	c := NewClient(nil)
	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL.String(), defaultBaseURL)
	}
	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", c.UserAgent, userAgent)
	}

	cNew := NewClient(nil)
	if c.client == cNew.client {
		t.Error("NewClient returned same http.Client, but they should differ.")
	}

}

func TestClient_Sign(t *testing.T) {
	key := `MIIEpAIBAAKCAQEAs1oFT17Tl0pB+TPrUcAi/65Iyu2PzmITHFeLQ63ptZiDi5SOmA579B6t+PqInOHwMIlc1BCxAXujaEkJnTNYM0pJrRYbfjzxNOqT/KzeP1Tbw2+ljetP85izDGO8VStxiwdR86X19cjNYzJCln73De4r9oy3bexFvEroWmsFD2Hl9PtyIbxYImptAjPHSTdiZyT8eV+wCQ92pfK+Soj65DKn7soUCRw+c15eix234YW2MVZohhL5KaUzT0nWrQ/HdHm9YGEycT6GCFahl85HLxgXKVG+zWEynkSnQjQNAemFsnI7AKXcOi/3/8oTK44nNzoJ+Vx27ttmVZMm21qs3wIDAQABAoIBACEhPd3dfGRz+R/ZcVyWi2CDhzrFC01qXHFd0oHb/FHpLFN6S6TW+BquGCDn1ph0O/QbS6R4uNm3RwYah8YcF+zRii6I6Oc2xq+prVB9dkhomnfNhd1jXE95I4nkWoO5FK5D2eWSAI84Wog6XoaCSYyvNDvE3pVVkcQoaidZoYOnmF8xAJ7+IBaQnLNwbdxymLteMKo8hg9FmKcVJwUP3yyofapEXh+m5i3Vw3f1ygTOTt/U7zhpR7hNAZnEE60CVB/NiZwav3mUa/iGKOHBSgkXxss61E5RWE6HPm9on4YhSVkwBoJDslCi7ZymKwtq/8juJyl1CtA6oF3+vFHhWXECgYEA6q2hbKoFF6Sj5w7WBiZH8Wxq4dpuch5l5+QWUOWYZSsSuKNlqQF5cgiuYrzT9OSh8aCvfvfBRzqUNMuz190BGli2qKk1EUNAXfYOPYFT8K9yjtXcsRolfCyJvDQcG1/deq0Ymmc7dmtdUSW1DHYN0ishkSKWgPTUBTvNylzbrr0CgYEAw6WNMy/Cb817rhXwJUSXRxFVkpjpXnyUqp0ycwPGE0ezqFpdAi9T10VCX+tb8rliD80tbKLSGBSbdrylZt9+hHZGwG7X+avGTAAi47ro5XvCdqBPzZoThWMQX0M7yedj+zxNU5PUHj2wY51lbowYb1rjj7pORAXXJw+5Yr2P4csCgYAfEUVhgVCTRLR+DjS/M6oqh9POeoBMk0GiZUufAgYevokH8Hmw2gUd15BeeekFbo8R0dfp0xq/Nz4PjzgLx7IxrrTftybTVjfbOIvCELDwGAnuZtmEiJWVzr+dLrjZ1uMXLnsnVatFjgkYiQli0O4beKJE+HM8Ny0qVDR2KbzbcQKBgQC3PWX7C/87AaSWDtwFu7FIEZcpH2hrY5lGFsb6bA1nX1+IDWJROox28dpRxjkvI6wRiqTBu9m0ThCxa8wqhqaTX8eyUd6ca+LkqbgDvLqGvJwMyOcVdPpVksvxvSHRX5QH6Zmu7qTb0gFTu0YV/Mah/OlnYyIMmb/CSPrj3RFWcwKBgQDdfHKlQjDoonna3hjs6fuHmo/08DbyREa7MoTd0uPDGE5dXSKnhWxmfRyK028EM4AUugx3+YfAhJi7lT9knypEN+9gfkJUDNkV+TipK96CwUSzhHVTfJG7YBafnp4x1tkm5YGPbEtxMqPI4fEEa0reT1/T4q1WKeNHyKs1Bsq+LQ==`
	want := "itrO06cvCtAOl8uxP6KyZmPl/UTAuYC5vdgBgLZryAOcFYGZg4bA2leCjdOpfqSJLYSAPIgLAt3TvFVvGGy/Jo8g/Cwv4rrkqmKuUJEcoBN0hPA93eyR5lyUuBfHypLzpthtH1ZbHWQqrKlBQaSysk+LFr2HG3fI4Q4YT/Hz0eZf7G5gcUHx1ToxUReVqhhGJMKU12yqBanvN7oMfqskqvffDcBAJWKhTL8Z2IkZ1LvgroVqYhdQQrBoZxD+qrBqw3pMKNsw5ux+RbIiOezWWsKyqQRBKJX26/6jR0+KJj/eznwBOtEFGTVq/jPuFDdeMEhUPQ4ZO+JnWrDkSJd+Zg=="
	c := NewClient(nil, PrivateKey(key), AppID("2016091100484533"), SignType("RSA2"), Charset("UTF-8"))
	values := &url.Values{}
	values.Set("method", "alipay.trade.create")
	values.Set("app_id", c.o.AppID)
	values.Set("sign_type", c.o.SignType)
	values.Set("biz_content", `{"out_trade_no":"202004191441122314312","total_amount":"88.88","subject":"iPhone+Xs+Max+256G","buyer_id":"2088102175953034","goods_detail":[],"seller_id":"2088102175107499"}`)
	values.Set("charset", c.o.Charset)
	values.Set("timestamp", "2020-04-19 14:41:12")
	values.Set("version", "1.0")
	v, _ := url.QueryUnescape(values.Encode())
	got, err := c.Sign(v)
	if err != nil {
		t.Errorf("Client.Sign returned unexcept err: %v", err)
	}
	if got != want {
		t.Errorf("Client.Sign got %v, want %v", got, want)
	}

}

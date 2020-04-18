package alipay

import "testing"

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

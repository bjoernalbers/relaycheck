package main

import (
	"net/http/httptest"
	"testing"
)

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		xForwardedFor string
		remoteAddr   string
		want         string
	}{
		{
			"",
			"",
			"",
		},
		{
			"",
			"1.2.3.4",
			"1.2.3.4",
		},
		{
			"",
			"1.2.3.4:56789",
			"1.2.3.4",
		},
		{
			"5.6.7.8",
			"1.2.3.4:12345",
			"5.6.7.8",
		},
		{
			"5.6.7.8",
			"1.2.3.4:12345",
			"5.6.7.8",
		},
		{
			" 5.6.7.8 ",
			"1.2.3.4:12345",
			"5.6.7.8",
		},
		{
			"5.6.7.8, 10.11.12.13",
			"1.2.3.4:12345",
			"5.6.7.8",
		},
		{
			"5.6.7.8,10.11.12.13",
			"1.2.3.4:12345",
			"5.6.7.8",
		},
	}
	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/", nil)
		r.RemoteAddr = tt.remoteAddr
		r.Header.Set("X-Forwarded-For", tt.xForwardedFor)
		if got := getClientIP(r); got != tt.want {
			t.Errorf("getClientIP() = %q, want: %q", got, tt.want)
		}
	}
}

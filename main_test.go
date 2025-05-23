package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		xForwardedFor string
		remoteAddr    string
		want          string
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

func TestRelayCheck(t *testing.T) {
	tests := []struct {
		ip   string
		want response
	}{
		{
			"1.2.3.4",
			response{Relay: false, IP: "1.2.3.4"},
		},
		{
			aRelayIP,
			response{Relay: true, IP: aRelayIP, Location: &location{CountryCode: "DE", RegionCode: "DE-BE", City: "Berlin"}},
		},
	}
	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("X-Forwarded-For", tt.ip)
		rr := httptest.NewRecorder()
		relayCheck(rr, req)
		var resp response
		json.NewDecoder(rr.Body).Decode(&resp)
		if got, want := rr.Header().Get("Content-Type"), "application/json"; got != want {
			t.Fatalf("Content-Type = %q, want: %q", got, want)
		}
		if diff := cmp.Diff(tt.want, resp); diff != "" {
			t.Errorf("response mismatch (-want +got):\n%s", diff)
		}
	}
}

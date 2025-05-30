package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/kmikiy/go-icloud-private-relay/relay"
)

const aRelayIP = "172.225.6.92"

// version gets set via ldflags.
var version = "unset"

// response is returned to the client.
type response struct {
	Relay    bool      `json:"relay"`
	IP       string    `json:"ip"`
	Location *location `json:"location,omitempty"`
}

// location represents an iCloud Private Relay location .
type location struct {
	CountryCode string `json:"country_code"`
	RegionCode  string `json:"region_code"`
	City        string `json:"city"`
}

func init() {
	log.SetFlags(0)
	flag.Usage = usage
}

func main() {
	addr := flag.String("addr", ":8080", "Address to listen to.")
	flag.Parse()
	warmUpCache()
	http.HandleFunc("/", relayCheck)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// usage prints usage instructions.
func usage() {
	header := fmt.Sprintf(`relaycheck - Simple HTTP API to detect iCloud Private Relay clients (version: %s)

Usage: relaycheck [options]

Options:`, version)
	fmt.Fprintln(flag.CommandLine.Output(), header)
	flag.PrintDefaults()
}

// warmUpCache performs a sample query to fetch the relay list from Apple.
func warmUpCache() {
	relay.IsICloudPrivateRelayAddress(aRelayIP)
}

// relayCheck handles the actual requests.
func relayCheck(w http.ResponseWriter, req *http.Request) {
	ip := getClientIP(req)
	resp := response{IP: ip}
	if loc, err := relay.ICloudPrivateRelay(ip); err != nil {
		resp.Relay = false
	} else {
		resp.Relay = true
		resp.Location = &location{
			CountryCode: loc.CountryCode,
			RegionCode:  loc.State,
			City:        loc.City,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}

// getClientIP returns the client IP address from the HTTP request.
// If the request has been forwarded by a reverse proxy, the address is
// extracted from the "X-Forwared-For" header.
// Otherwise the remote address will be returned.
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ip, _, _ := strings.Cut(xff, ",")
		return strings.TrimSpace(ip)
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return ip
	}
	return r.RemoteAddr
}

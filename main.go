// relaycheck
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

// version gets set via ldflags
var version = "unset"

func init() {
	log.SetFlags(0)
	flag.Usage = Usage
}

func main() {
	addr := flag.String("addr", ":8080", "Address to listen to.")
	flag.Parse()
	warmUpCache()
	http.HandleFunc("/", relayCheck)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// Usage prints usage instructions.
func Usage() {
	header := fmt.Sprintf(`relaycheck - Simple HTTP API to detect iCloud Private Relay clients (version: %s)

Usage: relaycheck [options]

Options:`, version)
	fmt.Fprintln(flag.CommandLine.Output(), header)
	flag.PrintDefaults()
}

// warmUpCache performs a sample query to fetch the relay list from Apple.
func warmUpCache() {
	isRelay(aRelayIP)
}

// relayCheck handles the actual requests.
func relayCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := getClientIP(req)
	json.NewEncoder(w).Encode(&response{
		Relay: isRelay(ip),
		IP:    ip,
	})
}

// response is returned to the client.
type response struct {
	Relay bool   `json:"relay"`
	IP    string `json:"ip"`
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

// isRelay returns true if the IP is an iCloud Private Relay, otherwise false.
func isRelay(ip string) bool {
	return relay.IsICloudPrivateRelayAddress(ip)
}

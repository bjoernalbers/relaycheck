// relaycheck
package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/kmikiy/go-icloud-private-relay/relay"
)

func init() {
	log.SetFlags(0)
}

func main() {
	http.HandleFunc("/", relayCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// relayCheck handles the actual requests.
func relayCheck(w http.ResponseWriter, req *http.Request) {
	var resp response
	resp.Relay = relay.IsICloudPrivateRelayAddress(getClientIP(req))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// response is returned to the client.
type response struct {
	Relay bool `json:"relay"`
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

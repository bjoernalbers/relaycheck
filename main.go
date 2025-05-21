// relaycheck
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("----- New HTTP Request -----")

		// Request-Line
		fmt.Printf("%s %s %s\n", r.Method, r.URL.RequestURI(), r.Proto)

		// Headers
		fmt.Println("Headers:")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Printf("  %s: %s\n", name, value)
			}
		}

		// RemoteAddr
		fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
		fmt.Printf("Client IP: %s\n", getClientIP(r))
		fmt.Println("----------------------------")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getClientIP returns the client IP address from the http request.
// If the request has been forwared by a reverse proxy, the address is
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

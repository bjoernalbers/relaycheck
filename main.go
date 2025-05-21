// relaycheck
package main

import (
	"fmt"
	"log"
	"net/http"
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

		fmt.Println("----------------------------")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

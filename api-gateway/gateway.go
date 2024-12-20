// api-gateway/gateway.go
package main

import (
	"net/http"
)

func main() {
	// Proxy requests to specific services
	http.HandleFunc("/api/booking", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("http://localhost:8080/booking")
		w.WriteHeader(resp.StatusCode)
	})

	http.ListenAndServe(":9000", nil)
}
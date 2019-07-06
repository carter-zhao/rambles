package server

import (
	"io"
	"net/http"
)

//http.HandleFunc("/health-check", HealthCheckHandler)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.URL.EscapedPath() != "/health-check" {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

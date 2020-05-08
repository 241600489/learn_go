package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.Handle("/metrix", promhttp.Handler())
	_ = http.ListenAndServe(":9001", server)

}

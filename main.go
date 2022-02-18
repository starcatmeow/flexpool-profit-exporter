package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	prometheus.MustRegister(&Collector{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal(err)
	}
}

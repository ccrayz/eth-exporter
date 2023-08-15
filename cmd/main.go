package main

import (
	"fmt"
	"log"
	"net/http"

	"ccrayz/eth-exporter/internal/ethereum"
	"ccrayz/eth-exporter/internal/prometheus"

	"ccrayz/eth-exporter/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	var ethAccounts = config.LoadConfig("config.yml")

	// Ethereum metrics collection
	rpcEndpoint := "https://api.sepolia.kroma.network"
	ethMetrics := ethereum.NewMetricsCollector(rpcEndpoint)

	exporter := prometheus.NewExporter(ethMetrics, ethAccounts)

	prometheus.RegisterExporter(exporter)

	// Start an HTTP server to expose metrics for Prometheus scraping
	fmt.Println("Starting server on port 8080")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func LoadConfig(s string) {
	panic("unimplemented")
}

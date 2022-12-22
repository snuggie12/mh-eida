package server

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (server *Server) startMetrics() {
	logger := server.Logger
	metricsConf := server.AdminConfig.MetricsConfig
	metricsMux := http.NewServeMux()
	metricsMux.Handle(metricsConf.Path, promhttp.Handler())
	logger.Infow("Starting metrics server.",
		"path", metricsConf.Path,
		"port", metricsConf.Port,
	)

	err := http.ListenAndServe(fmt.Sprintf(":%v", metricsConf.Port), metricsMux)
	if err != nil {
		logger.Error("Failed to start metrics endpoint")
	}
}

func (server *Server) addMetricsToAdmin(mux *http.ServeMux) {
  logger := server.Logger
	logger.Info("Placeholder for addMetricsToAdmin")
}

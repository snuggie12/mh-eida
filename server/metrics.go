package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	MetricsPath = "/metrics"
)

func (server *Server) addMetricsToAdmin(mux *http.ServeMux) {
	mux.Handle(MetricsPath, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
}

package metrics

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	metricsTypes "snuggie12/eida/pkg/types/metrics"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	MetricsPath = "/metrics"
)

var (
	defaultReceiverLabels = []string{
		"receiver_name",
		"receiver_port",
		"status_code",
	}
	receiverRequestsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "eida_receiver_requests_total",
		Help: "HTTP requests incoming to a receiver.",
	}, defaultReceiverLabels)

	receiverRequestsHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "eida_receiver_requests_duration_seconds",
			Help:    "HTTP requests incoming to a receiver performance.",
			Buckets: []float64{0.1, 0.25, .5, 1, 2, 4, 8, 16},
		}, defaultReceiverLabels,
	)
)

//MetricsServer is an aggregate of specific component metrics servers
type MetricsServer struct {
	Path                  string
	ReceiverMetricsServer *ReceiverMetricsServer
	Registry              *prometheus.Registry
	SenderMetricsServer   *SenderMetricsServer
}

//ReceiverMetricsServer holds all metrics and channels that receive info for labels
type ReceiverMetricsServer struct {
	RequestsCounter       *prometheus.CounterVec
	RequestsCounterChan   chan metricsTypes.ReceiverRequestsCounterInfo
	RequestsHistogram     *prometheus.HistogramVec
	RequestsHistogramChan chan metricsTypes.ReceiverRequestsHistogramInfo
}

type SenderMetricsServer struct{}

func NewMetricsServer(mux *http.ServeMux) (*MetricsServer, error) {
	reg := prometheus.NewRegistry()
	mux.Handle(MetricsPath, promhttp.HandlerFor(prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}, promhttp.HandlerOpts{}))

	reg.MustRegister(receiverRequestsCounter)
	reg.MustRegister(receiverRequestsHistogram)

	requestsCounterChan := make(chan metricsTypes.ReceiverRequestsCounterInfo)
	requestsHistogramChan := make(chan metricsTypes.ReceiverRequestsHistogramInfo)

	return &MetricsServer{
		Path: MetricsPath,
		ReceiverMetricsServer: &ReceiverMetricsServer{
			RequestsCounter:       receiverRequestsCounter,
			RequestsCounterChan:   requestsCounterChan,
			RequestsHistogram:     receiverRequestsHistogram,
			RequestsHistogramChan: requestsHistogramChan,
		},
		Registry:            reg,
		SenderMetricsServer: &SenderMetricsServer{},
	}, nil
}

var invalidPromLabelChars = regexp.MustCompile(`[^a-zA-Z0-9_]`)

func normalizeLabels(prefix string, appLabels []string) []string {
	normLabels := []string{}
	for _, label := range appLabels {
		//prometheus labels don't accept dash in their name
		curr := invalidPromLabelChars.ReplaceAllString(label, "-")
		normLabel := fmt.Sprintf("%s_%s", prefix, curr)
		normLabels = append(normLabels, normLabel)
	}
	return normLabels
}

func (metricsSrv *ReceiverMetricsServer) ProcessMetricsInfo(ctx context.Context, componentErrorChan chan error, logger *zap.SugaredLogger) {
	for {
		select {
		case counterInfo := <-metricsSrv.RequestsCounterChan:
			metricsSrv.RequestsCounter.WithLabelValues(counterInfo.ParseInfo()).Inc()
		case histogramInfo := <-metricsSrv.RequestsHistogramChan:
			metricsSrv.RequestsHistogram.WithLabelValues(histogramInfo.ParseInfo()).Observe(histogramInfo.Duration.Seconds())
		case <-componentErrorChan:
			logger.Info("Stopping metrics receiver")
			return
		}
	}
}

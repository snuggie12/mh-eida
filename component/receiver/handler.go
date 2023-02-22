package receiver

import (
	"fmt"
	"net/http"
	"time"

	metricsTypes "snuggie12/eida/pkg/types/metrics"
)

func (httpReceiver *httpReceiver) handleWebhook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST accepted", http.StatusNotImplemented)
		httpReceiver.sendMetrics(http.StatusNotImplemented, 0)
		return
	}

	t := time.Now()
	w.WriteHeader(http.StatusOK)

	//Pass the data to the parser
	

	elapsed := time.Since(t)

	httpReceiver.sendMetrics(http.StatusOK, elapsed)

	response := fmt.Sprintf(
		"%v - Hello from receiver %v on port %v and path %v\n",
		http.StatusOK,
		httpReceiver.Name,
		httpReceiver.HttpPort,
		httpReceiver.Path,
	)
	
	w.Write([]byte(response))
}

func (httpReceiver *httpReceiver) sendMetrics(statusCode int, duration time.Duration) {
	// Send counter info to the metrics channel
	reqCounterChan := httpReceiver.MetricsServer.ReceiverMetricsServer.RequestsCounterChan
	reqCounterChan <- *metricsTypes.NewReceiverRequestsCounterInfo(
		httpReceiver.Name,
		httpReceiver.HttpPort,
		statusCode,
	)

	// Send histogram info to the metrics channel
	recHistogramChan := httpReceiver.MetricsServer.ReceiverMetricsServer.RequestsHistogramChan
	recHistogramChan <- *metricsTypes.NewReceiverRequestsHistogramInfo(
		httpReceiver.Name,
		httpReceiver.HttpPort,
		statusCode,
		duration,
	)
}

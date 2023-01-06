package receiver

import (
	"fmt"
	"net/http"
	"snuggie12/eida/config"
	"snuggie12/eida/server/metrics"
	metricsTypes "snuggie12/eida/pkg/types/metrics"
	"time"

	"go.uber.org/zap"
)

type httpReceiver struct {
	Address       string
	Name          string
	Path          string
	Port          string
	MetricsServer *metrics.MetricsServer
}

func NewHttpReceiver(receiverConf *config.ReceiverConfig, metricsServer *metrics.MetricsServer) *httpReceiver {
	return newHttpReceiver(receiverConf, metricsServer)
}

func newHttpReceiver(receiverConf *config.ReceiverConfig, metricsServer *metrics.MetricsServer) *httpReceiver {
	return &httpReceiver{
		Address:       receiverConf.Address,
		Name:          receiverConf.Name,
		Path:          receiverConf.Path,
		Port:          receiverConf.Port,
		MetricsServer: metricsServer,
	}
}

func startHttpReceiver(receiverConf *config.ReceiverConfig, logger *zap.SugaredLogger, metricsServer *metrics.MetricsServer) error {
	httpReceiver := newHttpReceiver(receiverConf, metricsServer)

	recAddress := httpReceiver.Address
	recName := httpReceiver.Name
	recPath := httpReceiver.Path
	recPort := httpReceiver.Port

	logger.Debugw("startHttpReceiver configs",
		"address", recAddress,
		"path", recPath,
		"port", recPort,
	)

	logger.Debug("before vars")
	var (
		recCounterInfo   metricsTypes.ReceiverRequestsCounterInfo
		recHistogramInfo metricsTypes.ReceiverRequestsHistogramInfo
	)

	logger.Debug("after vars, before metrics server")
	recMetricsServe := metricsServer.ReceiverMetricsServer

	logger.Debug("before chans")
	recCounterChan := recMetricsServe.RequestsCounterChan
	recHistogramChan := recMetricsServe.RequestsHistogramChan
	logger.Debug("after chans")

	logger.Debug("before info")

	recCounterInfo.ReceiverName, recHistogramInfo.ReceiverName = recName, recName
	recCounterInfo.ReceiverPort, recHistogramInfo.ReceiverPort = recPort, recPort

	twoExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusOK)
		recCounterInfo.StatusCode, recHistogramInfo.StatusCode = "200", "200"
		recHistogramInfo.Duration = time.Since(t)
		recCounterChan <- recCounterInfo
		recHistogramChan <- recHistogramInfo

		w.Write([]byte("200 - Hello from example application.\n"))
	})
	fourExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusNotFound)
		recCounterInfo.StatusCode, recHistogramInfo.StatusCode = "404", "404"
		recHistogramInfo.Duration = time.Since(t)
		recCounterChan <- recCounterInfo
		recHistogramChan <- recHistogramInfo

		w.Write([]byte("404 - Not Found\n"))
	})
	fiveExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusInternalServerError)
		recCounterInfo.StatusCode, recHistogramInfo.StatusCode = "503", "503"
		recHistogramInfo.Duration = time.Since(t)
		recCounterChan <- recCounterInfo
		recHistogramChan <- recHistogramInfo

		w.Write([]byte("503 - Bad Gateway\n"))
	})

	mux := http.NewServeMux()
	mux.Handle("/", twoExExHandler)
	mux.Handle("/err", fourExExHandler)
	mux.Handle("/internal-err", fiveExExHandler)

	var srv *http.Server
	srv = &http.Server{Addr: fmt.Sprintf("%v:%v", httpReceiver.Address, httpReceiver.Port), Handler: mux}
	go logger.Fatal(srv.ListenAndServe())
	logger.Debug("donezo")
	return nil
}

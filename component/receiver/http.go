package receiver
/* 
import (
	"fmt"
	"net/http"
	"snuggie12/eida/config"
	"snuggie12/eida/server/metrics"
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
		"metrics-server", httpReceiver.MetricsServer,
	)

	receiverMetricsServe := metricsServer.ReceiverMetricsServers[recName]

	twoExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusOK)
		receiverMetricsServe.IncRequestsHistogram(recName, recPort, "200", time.Since(t))
		receiverMetricsServe.IncRequestsCounter(recName, recPort, "200")
		w.Write([]byte("Hello from example application."))
	})
	fourExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusNotFound)
		receiverMetricsServe.IncRequestsHistogram(recName, recPort, "404", time.Since(t))
		receiverMetricsServe.IncRequestsCounter(recName, recPort, "404")
	})
	fiveExExHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusInternalServerError)
		receiverMetricsServe.IncRequestsHistogram(recName, recPort, "503", time.Since(t))
		receiverMetricsServe.IncRequestsCounter(recName, recPort, "503")

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
 */

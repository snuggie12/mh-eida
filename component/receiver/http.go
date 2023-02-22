package receiver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"snuggie12/eida/config"

	"snuggie12/eida/server/metrics"

	"go.uber.org/zap"
)

type serverAddress string
var serverAddrCtxKey serverAddress = "serverAddr"

type httpReceiver struct {
	HttpAddress   string
	Name          string
	Path          string
	HttpPort      string
	MetricsServer *metrics.MetricsServer
}

func NewHttpReceiver(receiverConf *config.ReceiverConfig, metricsServer *metrics.MetricsServer) *httpReceiver {
	return &httpReceiver{
		HttpAddress:   receiverConf.HttpConfig.HttpAddress,
		Name:          receiverConf.Name,
		Path:          receiverConf.HttpConfig.Path,
		HttpPort:      receiverConf.HttpConfig.HttpPort,
		MetricsServer: metricsServer,
	}
}

func (httpReceiver *httpReceiver) start(logger *zap.SugaredLogger) error {
	recAddress := httpReceiver.HttpAddress
	recPath := httpReceiver.Path
	recPort := httpReceiver.HttpPort

	logger.Debugw("startHttpReceiver configs",
		"address", recAddress,
		"path", recPath,
		"port", recPort,
	)

	ctx, cancel := context.WithCancel(context.Background())
	mux := http.NewServeMux()
	mux.Handle(recPath, http.HandlerFunc(httpReceiver.handleWebhook))

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", recAddress, recPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverAddrCtxKey, l.Addr().String())
			return ctx
		},
	}

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

	cancel()
	logger.Debug("donezo")
	return nil
}

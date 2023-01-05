package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"snuggie12/eida/config"
	metricsTypes "snuggie12/eida/pkg/types/metrics"
	"snuggie12/eida/server/metrics"
	"snuggie12/eida/util"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Logger *zap.SugaredLogger
}

func NewServer(conf *config.Config, logger *zap.SugaredLogger) *Server {
	return newServer(conf, logger)
}

func newServer(conf *config.Config, logger *zap.SugaredLogger) *Server {
	//adminConf := config.NewAdminConfig(conf.AdminConfigOptions)
	return &Server{
		Config: conf,
		Logger: logger,
	}
}

func (server *Server) StartAdminServer() {
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger := server.Logger

	adminServer := newAdminServer(server.Config)

	logger.Debugw("Admin Server Configuration",
		"address", util.FriendlyAddress(adminServer.config.Address),
		"port", adminServer.config.Port,
	)

	componentsConfig := server.Config.ParseFullConfig()
	logger.Debugw("Components config",
		"receivers", config.FriendlyReceiverConfigs(componentsConfig.ReceiverConfigs),
	)

	metricsServer, err := metrics.NewMetricsServer(adminServer.mux)
	if err != nil {
		logger.Fatalw("Error occurred while registering creating metrics server", "error", err)
	}

	recCounterChan := metricsServer.ReceiverMetricsServer.RequestsCounterChan
	recHistogramChan := metricsServer.ReceiverMetricsServer.RequestsHistogramChan

	metricsDone := make(chan bool)

	go metricsServer.ReceiverMetricsServer.ProcessMetricsInfo(metricsDone, logger)

	var recCounterInfo metricsTypes.ReceiverRequestsCounterInfo

	recCounterInfo.ReceiverName = "joebob"
	recCounterInfo.ReceiverPort = "5555"
	recCounterInfo.StatusCode = "200"

	recCounterChan <- recCounterInfo

	recCounterInfo.StatusCode = "569"

	recCounterChan <- recCounterInfo

	var recHistogramInfo metricsTypes.ReceiverRequestsHistogramInfo

	recHistogramInfo.ReceiverName = "joebob"
	recHistogramInfo.ReceiverPort = "5555"
	recHistogramInfo.StatusCode = "503"
	recHistogramInfo.Duration = time.Duration(time.Duration(10)*time.Second)

	recHistogramChan <- recHistogramInfo

	recHistogramInfo.Duration = time.Duration(time.Duration(200)*time.Millisecond)

	recHistogramChan <- recHistogramInfo

	//Start up the metrics server
	go http.ListenAndServe(fmt.Sprintf("%v:%v",
		adminServer.config.Address,
		adminServer.config.Port,
	), adminServer.mux)

	//receiver.InitializeReceivers(doneChan, componentsConfig.ReceiverConfigs, logger, adminServer.config.StrictLoadingEnabled, *metricsServer)

	<-doneChan

	logger.Info("Signal Received. Stopping the Server")
}

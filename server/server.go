package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"snuggie12/eida/component"
	"snuggie12/eida/config"
	"snuggie12/eida/server/metrics"
	"snuggie12/eida/util"
	"syscall"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	osSignalChan := make(chan os.Signal, 1)
	componentErrorChan := make(chan error)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger := server.Logger

	adminServer := newAdminServer(server.Config)

	logger.Debugw("Admin Server Configuration",
		"address", util.FriendlyAddress(adminServer.config.HttpAddress),
		"port", adminServer.config.HttpPort,
	)

	metricsServer, err := metrics.NewMetricsServer(adminServer.mux)
	if err != nil {
		logger.Fatalw("Error occurred while registering creating metrics server", "error", err)
	}

	go metricsServer.ReceiverMetricsServer.ProcessMetricsInfo(ctx, componentErrorChan, logger)

	//Start up the admin server
	go http.ListenAndServe(fmt.Sprintf("%v:%v",
		adminServer.config.HttpAddress,
		adminServer.config.HttpPort,
	), adminServer.mux)

	componentsConfig, err := server.Config.ParseFullConfig()
	if err != nil {
		logger.Errorw("Failed to parse component config", "error", err)
	}

	logger.Debugw("Components config",
		"receivers", config.FriendlyReceiverConfigs(componentsConfig),
	)

	components := component.NewComponents(componentsConfig, logger, componentErrorChan)

	go components.Senders.Start()

	go components.Parsers.Start()

	go components.Receivers.Start()

	for {
		select {
		case err := <-componentErrorChan:
			logger.Errorw("Component has failed.", "error", zap.Error(err))
			return
		case signal := <-osSignalChan:
			logger.Infow("Received signal. Shutting down", "signal", signal)
			return
		case <-ctx.Done():
			logger.Infow("Context done. Shutting down", "context error", zap.Error(ctx.Err()))
		}
	}

}

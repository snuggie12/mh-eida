package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"snuggie12/eida/component/receiver"
	"snuggie12/eida/config"
	"snuggie12/eida/util"
	"syscall"

	"go.uber.org/zap"
)

const (
	LivenessProbePath string = "/livez"
)

type Server struct {
	AdminConfig *config.AdminConfig
	FullConfig  *config.Config
	Logger      *zap.SugaredLogger
}

func NewServer(conf *config.Config, logger *zap.SugaredLogger) *Server {
	return newServer(conf, logger)
}

func newServer(conf *config.Config, logger *zap.SugaredLogger) *Server {
	adminConf := config.NewAdminConfig(conf.AdminConfigOptions)
	server := &Server{
		AdminConfig: adminConf,
		FullConfig:  conf,
		Logger:      logger,
	}

	return server
}

func (server *Server) StartAdminServer() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger := server.Logger
	logger.Info("Parsing Admin Server Config")

	err := server.AdminConfig.ParseAdminConfig()
	if err != nil {
		logger.Errorf("Unable to parse admin config: %v", err)
	}

	logger.Infow("Admin Server Configuration",
		"address", util.FriendlyAddress(server.AdminConfig.Address),
		"port", server.AdminConfig.Port,
	)

	adminMux := http.NewServeMux()

	server.addMetricsToAdmin(adminMux)
	server.addPprofToAdmin(adminMux)
	server.addHealthToAdmin(adminMux)

	go server.startAdmin(done, adminMux)

	componentsConfig := server.FullConfig.ParseFullConfig()
	logger.Debugw("Components config",
		"receivers", config.FriendlyReceiverConfigs(componentsConfig.ReceiverConfigs),
	)

	go receiver.StartReceivers(done, componentsConfig.ReceiverConfigs, logger, server.AdminConfig.StrictLoadingEnabled)

	<-done
	logger.Info("Signal Received. Stopping the Server")
}

func (server *Server) startAdmin(doneChan chan os.Signal, mux *http.ServeMux) {
	logger := server.Logger

	err := http.ListenAndServe(fmt.Sprintf("%v:%v", server.AdminConfig.Address, server.AdminConfig.Port), mux)
	if err != nil {
		logger.Errorf("Failed to start admin endpoint: %v", err)
		doneChan <- syscall.SIGILL
	}

}

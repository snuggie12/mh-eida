package server

import (
	//"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	//"snuggie12/eida/config"
)

type Server struct {
	AdminConfig *AdminConfig
	Logger      *zap.SugaredLogger
}

func NewServer(adminConf *AdminConfig, logger *zap.SugaredLogger) *Server {
	return newServer(adminConf, logger)
}

func newServer(adminConf *AdminConfig, logger *zap.SugaredLogger) *Server {
	server := &Server{
		AdminConfig: adminConf,
		Logger:      logger,
	}

	return server
}

func (server *Server) StartAdminServer() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger := server.Logger
	logger.Info("Parsing Admin Server Config")

	server.parseAdminConfig()

	logger.Infow("Admin Server Configuration",
		"host", server.AdminConfig.AdminHost,
		"port", server.AdminConfig.AdminPort,
	)
	//adminMux := http.NewServeMux()

	if server.AdminConfig.MergeMetricsToAdmin == true {
		logger.Debugw("Starting metrics on the same port as the admin API", "port", server.AdminConfig.AdminPort)
	}

	<-done
	logger.Info("Signal Received. Stopping the Server")
}

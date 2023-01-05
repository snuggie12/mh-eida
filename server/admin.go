package server

import (
	"net/http"
	"snuggie12/eida/config"
)

type adminServer struct {
	config *config.AdminConfig
	mux    *http.ServeMux
}

func newAdminServer(conf *config.Config) *adminServer {
	adminConf := config.NewAdminConfig(conf.AdminConfigOptions)

	adminMux := http.NewServeMux()

	addPprofToAdmin(adminMux)
	addHealthToAdmin(adminMux)

	return &adminServer{
		config: adminConf,
		mux:    adminMux,
	}
}

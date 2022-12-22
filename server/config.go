package server

import (
	"snuggie12/eida/util"
)

type AdminConfig struct {
	AdminHost   string `mapstructure:"host" json:"AdminHost"`
	AdminPort   string `mapstructure:"port" json:"AdminPort"`
	MergeMetricsToAdmin bool `json:"MergeMetricsToAdmin" default:"false"`
	MetricsConfig MetricsConfig `mapstructure:"metrics" json:"MetricsConfig"`
}

type MetricsConfig struct {
	Path string `mapstructure:"path" json:"MetricsPath"`
	Port string `mapstructure:"port" json:"MetricsPort"`
}

func (server *Server) parseAdminConfig() {
	conf := server.AdminConfig
	if util.PortsMatch(conf.AdminPort, conf.MetricsConfig.Port) == true {
		server.AdminConfig.MergeMetricsToAdmin = true
	}
	
	// TODO: Check we have legit host and metrics path
}



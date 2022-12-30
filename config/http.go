package config

import "snuggie12/eida/util"

type HttpConfigOptions struct {
	ListenLocal bool   `mapStructure:"listenLocal"`
	Port        string `mapStructure:"port"`
}

type HttpConfig struct {
	Address string
	Port    string
}

func NewHttpConfig(httpConfOpts *HttpConfigOptions) HttpConfig {
	var address string
	if httpConfOpts.ListenLocal == true {
		address = "127.0.0.1"
	}

	return HttpConfig{
		Address: address,
		Port:    httpConfOpts.Port,
	}
}

func (httpConfig *HttpConfig) parseHttpConfig() (error) {
	//TODO: make sure port is valid, IPs, etc
	return nil
}

func FriendlyHttpConfigs(httpConfs []HttpConfig) ([]HttpConfig) {
	for _, httpConf := range(httpConfs) {
		httpConf.Address = util.FriendlyAddress(httpConf.Address)
	}
	return httpConfs
}

package config

import (
	"snuggie12/eida/util"
	"strconv"
)

type HttpConfigOptions struct {
	ListenLocal bool   `mapstructure:"listenLocal"`
	Path        string `mapstructure:"path"`
	HttpPort    int    `mapstructure:"port"`
}

type HttpConfig struct {
	HttpAddress string `yaml:"address"`
	Path        string `yaml:"path"`
	HttpPort    string `yaml:"port"`
}

func NewHttpConfig(httpConfOpts *HttpConfigOptions) HttpConfig {
	var address string
	if httpConfOpts.ListenLocal == true {
		address = "127.0.0.1"
	}

	return HttpConfig{
		HttpAddress: address,
		Path:        httpConfOpts.Path,
		HttpPort:    strconv.Itoa(httpConfOpts.HttpPort),
	}
}

func (httpConfig *HttpConfig) parseHttpConfig() error {
	//TODO: make sure port is valid, IPs, etc
	return nil
}

func FriendlyHttpConfigs(httpConfs []HttpConfig) []HttpConfig {
	for _, httpConf := range httpConfs {
		httpConf.HttpAddress = util.FriendlyAddress(httpConf.HttpAddress)
	}
	return httpConfs
}

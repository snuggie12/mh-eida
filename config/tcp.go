package config

import "strconv"

type TcpConfigOptions struct {
	ListenLocal bool `mapstructure:"listenLocal"`
	TcpPort     int  `mapstructure:"port,omitempty"`
}

type TcpConfig struct {
	TcpAddress string `yaml:"address,omitempty" json:"address,omitempty"`
	TcpPort    string `yaml:"port,omitempty" json:"port,omitempty"`
}

func NewTcpConfig(tcpConfOpts *TcpConfigOptions) TcpConfig {
	var address string
	if tcpConfOpts.ListenLocal == true {
		address = "127.0.0.1"
	}

	return TcpConfig{
		TcpAddress: address,
		TcpPort:    strconv.Itoa(tcpConfOpts.TcpPort),
	}
}

package config

import (
	"go.uber.org/zap"
)

type Config struct {
	AdminConfigOptions *AdminConfigOptions `mapstructure:"admin,omitempty"`
	InputConfig        *InputConfig
	Logger             *zap.SugaredLogger
	ReceiverConfigsOptions []*ReceiverConfigOptions `mapstructure:"receivers"`
}

type ComponentsConfig struct {
	ReceiverConfigs []ReceiverConfig
}

func NewConfig(addr string, port string, logger *zap.SugaredLogger) *Config {
	return newConfig(addr, port, logger)
}

func newConfig(addr string, port string, logger *zap.SugaredLogger) *Config {
	inputConf := newInputConfig(addr, port)
	conf := &Config{
		InputConfig: inputConf,
	}

	return conf
}

func (conf *Config) ParseFullConfig() (*ComponentsConfig) {
	ComponentsConf := ComponentsConfig{
		ReceiverConfigs: ParseRecieverConfigs(conf.ReceiverConfigsOptions),
	}
	return &ComponentsConf
}

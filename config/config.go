package config

import (
	"go.uber.org/zap"
)

type Config struct {
	AdminConfigOptions     *AdminConfigOptions     `mapstructure:"admin,omitempty"`
	ComponentConfigOptions ComponentConfigOptions `mapstructure:",squash"`
	Logger                 *zap.SugaredLogger
}

func (conf *Config) ParseFullConfig() (*ComponentConfigs, error) {
	compConfOpts := conf.ComponentConfigOptions

	receiverConfigs := make(map[string]*ReceiverConfig)
	for receiverName, receiverConfOpts := range compConfOpts.ReceiverConfigOptions {
		receiverConf, err := receiverConfOpts.Parse(receiverName)
		if err != nil {
			return nil, err
		}

		receiverConfigs[receiverName] = receiverConf
	}

	ComponentsConf := ComponentConfigs{
		ReceiverConfigs: receiverConfigs,
	}
	return &ComponentsConf, nil
}

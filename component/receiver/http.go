package receiver

import (
	"snuggie12/eida/config"

	"go.uber.org/zap"
)

func NewHttpReceiver(receiverConf *config.ReceiverConfig) *HttpReceiver {
	return newHttpReceiver(receiverConf)
}

func newHttpReceiver(receiverConf *config.ReceiverConfig) *HttpReceiver {
	metricsServer := newMetricsServer()
	return &HttpReceiver{
		Address:       receiverConf.Address,
		Path:          receiverConf.Path,
		Port:          receiverConf.Port,
		MetricsServer: metricsServer,
	}
}

func startHttpReceiver(receiverConf *config.ReceiverConfig, logger *zap.SugaredLogger) error {
	httpReceiver := newHttpReceiver(receiverConf)
	logger.Debugw("startHttpReceiver configs",
		"address", httpReceiver.Address,
		"path", httpReceiver.Path,
		"port", httpReceiver.Port,
		"metrics-server", httpReceiver.MetricsServer,
	)

	logger.Debug("donezo")
	return nil
}

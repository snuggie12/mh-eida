package receiver

import (
	"context"
	"fmt"
	"snuggie12/eida/config"
	"snuggie12/eida/server/metrics"

	"go.uber.org/zap"
)

type OldReceiver struct {
	Config config.ReceiverConfig
}

func OldInitializeReceivers(
	componentErrorChan chan error,
	ctx context.Context,
	logger *zap.SugaredLogger,
	metricsServer *metrics.MetricsServer,
	receiverConfs []config.ReceiverConfig,
) {

	// Loop through all receiver configs
	for _, receiverConf := range receiverConfs {
		receiverType := config.GetGenericType(receiverConf.Type)

		// Attempt to start a receiver
		go OldstartReceiver(componentErrorChan, receiverType, receiverConf, logger, metricsServer)

		// If strict loading is on and there's an error then exit ASAP
/* 		if err == nil {
			logger.Errorw("Error while loading receiver and strict loading enabled",
				"receiver-name", receiverConf.Name,
				"error", err,
			)
			componentErrorChan <- err
		} */
	}
	logger.Info("Outside of startReceiver")
	for {
		select {
		case err := <- componentErrorChan:
			logger.Errorw("after start receiver loop", "error", err)
		}
	}
}

func OldstartReceiver(compErrChan chan error, recType string, recConf config.ReceiverConfig, logger *zap.SugaredLogger, metricsServer *metrics.MetricsServer) error {
	var err error

	switch recType {
	case "http":
		logger.Infof("Starting receiver on port %v", recConf.HttpPort)
		go startHttpReceiver(&recConf, logger, metricsServer)
/* 		if err != nil {
			logger.Errorf("Problem starting http receiver: %v", err)
		} */
	case "none":
		if recConf.Name == "" {
			recConf.Name = "missing-name"
		}
		err = fmt.Errorf("Receiver config missing type. Receiver: %v", recConf.Name)
	default:
		err = fmt.Errorf("Not quite sure how you got here. Turn back now. Unknown error")
	}

	return err

}

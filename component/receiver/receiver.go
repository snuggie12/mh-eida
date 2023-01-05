package receiver

/*
import (
	"fmt"
	"os"
	"snuggie12/eida/config"
	"snuggie12/eida/server/metrics"

	"go.uber.org/zap"
)

type Receiver struct {
	Config config.ReceiverConfig
}

func InitializeReceivers(
	doneChan chan os.Signal,
	receiverConfs []config.ReceiverConfig,
	logger *zap.SugaredLogger, isStrict bool,
	metricsServer metrics.MetricsServer) {
	var (
		errorCount int
	)

	// Loop through all receiver configs
	for _, receiverConf := range receiverConfs {
		receiverType := config.GetGenericType(receiverConf.Type)

		// Attempt to start a receiver
		go startReceiver(receiverType, receiverConf, logger, &metricsServer)

		/* 		// If strict loading is on and there's an error then exit ASAP
			if err != nil && isStrict {
				logger.Errorw("Error while loading receiver and strict loading enabled",
					"receiver-name", receiverConf.Name,
					"error", err,
				)
				doneChan <- syscall.SIGILL
			}

			// Strict loading is off and we got errors. We'll simply warn that the receiver will not load
			if err != nil {
				errorCount++
				logger.Warnw("Error adding receiver",
					"receiver-name", receiverConf.Name,
					"error", err,
				)
			}
		}

		if errorCount >= len(receiverConfs) {
			logger.Errorf("All receivers failed to load. Exiting...")
			doneChan <- syscall.SIGILL
	}
}

func startReceiver(recType string, recConf config.ReceiverConfig, logger *zap.SugaredLogger, metricsServer *metrics.MetricsServer) error {
	var err error

	switch recType {
	case "http":
		logger.Infof("Starting receiver on port %v", recConf.Port)
		err = startHttpReceiver(&recConf, logger, metricsServer)
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
 */

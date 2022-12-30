package receiver

import (
	"fmt"
	"os"
	"snuggie12/eida/config"
	"syscall"

	"go.uber.org/zap"
)

var typeToGenericType = map[string]string{
	"argocd":         "http",
	"launchDarkly":   "http",
	"genericWebhook": "http",
	"kubernetes":     "pull",
}

func StartReceivers(doneChan chan os.Signal, receiverConfs []config.ReceiverConfig, logger *zap.SugaredLogger, isStrict bool) {
	var (
		errorCount int
	)

	// Loop through all receiver configs
	for _, receiverConf := range receiverConfs {
		receiverType := typeToGenericType[receiverConf.Type]
		if receiverType == "" {
			receiverType = "none"
		}

		// Attempt to start a receiver
		err := startReceiver(receiverType, receiverConf, logger)

		// If strict loading is on and there's an error then exit ASAP
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

func startReceiver(recType string, recConf config.ReceiverConfig, logger *zap.SugaredLogger) error {
	var err error

	switch recType {
	case "http":
		err = startHttpReceiver(&recConf, logger)
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

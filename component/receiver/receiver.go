package receiver

import (
	"snuggie12/eida/component/common"
	"snuggie12/eida/config"
)

type Receivers []*Receiver
/* struct {
	Receivers []*Receiver
} */

type Receiver struct {
	ComponentCommon common.ComponentCommon
	ReceiverName string
	ReceiverConfig config.ReceiverConfig
}

func InitializeReceivers(compCommon *common.ComponentCommon, receiverConfs map[string]*config.ReceiverConfig) (Receivers, error) {
	logger := compCommon.Logger
	receivers := make([]*Receiver, 0)
	for receiverName, receiverConf := range(receiverConfs) {
		receiver, err := initializeReceiver(compCommon, receiverName, receiverConf)
		if err != nil {
			logger.Error("Error while initializing receivers.", "receiver", receiverName)
			compCommon.ComponentErrorChan <- err
			return nil, err
		}

		logger.Debugw("Initialized receiver and appending to receivers", "receiver", receiver.ReceiverConfig)
		receivers = append(receivers, receiver)
	}

	return receivers, nil
}

func initializeReceiver(compCommon *common.ComponentCommon, receiverName string, receiverConf *config.ReceiverConfig) (*Receiver, error) {
	err := validateReceiverConfig(receiverConf)
	if err != nil {
		compCommon.ComponentErrorChan <- err
		return nil, err
	}

	return &Receiver{
		ComponentCommon: *compCommon,
		ReceiverConfig: *receiverConf,
	}, nil
}

func validateReceiverConfig(receiverConf *config.ReceiverConfig) error {
	//TODO: actually validate something
	return nil
}

func (receivers Receivers) Start() {
	for _, receiver := range receivers {
		receiver.start()
	}
}

func (receiver Receiver) start() {
	logger := receiver.ComponentCommon.Logger
	logger.Info("inside of receiver.start()")
}

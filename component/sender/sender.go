package sender

import (
	"snuggie12/eida/component/common"
	"snuggie12/eida/config"
)

type Senders []*Sender

type Sender struct {
	ComponentCommon common.ComponentCommon
	SenderName string
	SenderConfig config.SenderConfig
}

func InitializeSenders(compCommon *common.ComponentCommon, senderConfs map[string]*config.SenderConfig) (Senders, error) {
	logger := compCommon.Logger
	senders := make([]*Sender, 0)
	for senderName, senderConf := range(senderConfs) {
		sender, err := initializeSender(compCommon, senderName, senderConf)
		if err != nil {
			logger.Error("Error while initializing senders.", "sender", senderName)
			compCommon.ComponentErrorChan <- err
			return nil, err
		}

		logger.Debugw("Initialized sender and appending to senders", "sender", sender)
		senders = append(senders, sender)
	}

	return senders, nil
}

func initializeSender(compCommon *common.ComponentCommon, senderName string, senderConf *config.SenderConfig) (*Sender, error) {
	err := validateSenderConfig(senderConf)
	if err != nil {
		compCommon.ComponentErrorChan <- err
		return nil, err
	}

	return &Sender{
		ComponentCommon: *compCommon,
		SenderConfig: *senderConf,
	}, nil
}

func validateSenderConfig(senderConf *config.SenderConfig) error {
	//TODO: actually validate something
	return nil
}

func (senders *Senders) Start() {
	return
}

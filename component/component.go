package component

import (
	"snuggie12/eida/component/common"
	"snuggie12/eida/component/parser"
	"snuggie12/eida/component/receiver"
	"snuggie12/eida/component/sender"
	"snuggie12/eida/config"

	"go.uber.org/zap"
)

type ComponentsInt interface {
	Start()
}

type Components struct {
	Receivers *receiver.Receivers
	Parsers   *parser.Parsers
	Senders   *sender.Senders
}

func NewComponents(componentsConfig *config.ComponentConfigs, logger *zap.SugaredLogger, compErrChan chan error) *Components {
	compCommon := common.NewComponentCommon(logger, compErrChan)

	receivers, err := receiver.InitializeReceivers(compCommon, componentsConfig.ReceiverConfigs)
	if err != nil {
		logger.Errorw("Error initializing receivers.", "error", err)
	}

	parsers, err := parser.InitializeParsers(compCommon, componentsConfig.ParserConfigs)
	if err != nil {
		logger.Errorw("Error initializing parsers.", "error", err)
	}

	senders, err := sender.InitializeSenders(compCommon, componentsConfig.SenderConfigs)
	if err != nil {
		logger.Errorw("Error initializing senders.", "error", err)
	}

	return &Components{
		Receivers: &receivers,
		Parsers:   &parsers,
		Senders:   &senders,
	}
}

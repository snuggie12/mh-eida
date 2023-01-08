package parser

import (
	"snuggie12/eida/component/common"
	"snuggie12/eida/config"
)

type Parsers []*Parser

type Parser struct {
	ComponentCommon common.ComponentCommon
	ParserName string
	ParserConfig config.ParserConfig
}

func InitializeParsers(compCommon *common.ComponentCommon, parserConfs map[string]*config.ParserConfig) (Parsers, error) {
	logger := compCommon.Logger
	parsers := make([]*Parser, 0)
	for parserName, parserConf := range(parserConfs) {
		parser, err := initializeParser(compCommon, parserName, parserConf)
		if err != nil {
			logger.Errorw("Error while initializing parser.", "parser", parserName)
			compCommon.ComponentErrorChan <- err
			return nil, err
		}

		logger.Debugw("Initialized parser and appending to parsers", "parser", parser)
		parsers = append(parsers, parser)
	}

	return parsers, nil
}

func initializeParser(compCommon *common.ComponentCommon, parserName string, parserConf *config.ParserConfig) (*Parser, error) {
	err := validateParserConfig(parserConf)
	if err != nil {
		compCommon.ComponentErrorChan <- err
		return nil, err
	}

	return &Parser{
		ComponentCommon: *compCommon,
		ParserConfig: *parserConf,
	}, nil
}

func validateParserConfig(parserConf *config.ParserConfig) error {
	//TODO: actually validate something
	return nil
}

func (parsers *Parsers) Start() {
	return
}

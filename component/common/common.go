package common

import "go.uber.org/zap"

type ComponentCommon struct {
	ComponentErrorChan chan error
	Logger *zap.SugaredLogger
}

func NewComponentCommon(logger *zap.SugaredLogger, compErrChan chan error) *ComponentCommon {
	return &ComponentCommon{
		ComponentErrorChan: compErrChan,
		Logger: logger,
	}
}

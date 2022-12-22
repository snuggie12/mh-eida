package util

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

type LoggingConfig struct {
	LogLevel string `mapstructure:"log-level" json:"LogLevel"`
}

func Logger(level string) *zap.SugaredLogger {
	atomicLevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		fmt.Printf("Cannot parse logging level: %v", err)
		os.Exit(1)
	}

	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(atomicLevel.Level())

	logger := zap.Must(cfg.Build())

	defer logger.Sync()

	return logger.Sugar()
}

package cmd

import(
	"snuggie12/eida/config"
	cmdutil "snuggie12/eida/cmd/util"
)

type CmdConfig struct {
	Config config.Config `mapstructure:",squash"`
	LoggingConfig cmdutil.LoggingConfig `mapstructure:"logging"`
}

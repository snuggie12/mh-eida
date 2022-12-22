package cmd

import(
	"snuggie12/eida/server"
	cmdutil "snuggie12/eida/cmd/util"
)

type Config struct {
	AdminConfig server.AdminConfig `mapstructure:"admin"`
	LoggingConfig cmdutil.LoggingConfig `mapstructure:"logging"`
}

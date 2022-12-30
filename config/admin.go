package config

type AdminConfigOptions struct {
	HttpConfigOptions  `mapstructure:",squash"`
	StrictLoadingEnabled bool `mapstructure:"strictLoadingEnabled"`
}

type AdminConfig struct {
	HttpConfig
	StrictLoadingEnabled bool
}

func NewAdminConfig(adminConfOpts *AdminConfigOptions) *AdminConfig {
	return newAdminConfig(adminConfOpts)
}

func newAdminConfig(adminConfOpts *AdminConfigOptions) *AdminConfig {
	adminHttpConf := NewHttpConfig(&adminConfOpts.HttpConfigOptions)
	return &AdminConfig{
		HttpConfig: adminHttpConf,
		StrictLoadingEnabled: adminConfOpts.StrictLoadingEnabled,
	}
}

func (conf *AdminConfig) ParseAdminConfig() error {
	return nil
}

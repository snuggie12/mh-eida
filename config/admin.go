package config

type AdminConfigOptions struct {
	HttpConfigOptions  `mapstructure:",squash"`
}

type AdminConfig struct {
	HttpConfig
}

func NewAdminConfig(adminConfOpts *AdminConfigOptions) *AdminConfig {
	return newAdminConfig(adminConfOpts)
}

func newAdminConfig(adminConfOpts *AdminConfigOptions) *AdminConfig {
	adminHttpConf := NewHttpConfig(&adminConfOpts.HttpConfigOptions)
	return &AdminConfig{
		HttpConfig: adminHttpConf,
	}
}

package config

type AdminConfig struct {
	HttpConfig
}

func NewAdminConfig(adminConfOpts *HttpConfigOptions) *AdminConfig {
	return newAdminConfig(adminConfOpts)
}

func newAdminConfig(adminConfOpts *HttpConfigOptions) *AdminConfig {
	adminConf := NewHttpConfig(adminConfOpts)
	return &AdminConfig{
		HttpConfig: adminConf,
	}
}

func (conf *AdminConfig) ParseAdminConfig() error {
	return nil
}

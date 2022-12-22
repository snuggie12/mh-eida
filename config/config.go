package config

type Config struct {
	InputConfig *InputConfig
}

func NewConfig(addr string, port string) *Config {
	return newConfig(addr, port)
}

func newConfig(addr string, port string) *Config {
	inputConf := newInputConfig(addr, port)
	conf := &Config{
		InputConfig: inputConf,
	}

	return conf
}

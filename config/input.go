package config

type InputConfig struct {
	Address string
	Port string
}

func newInputConfig(addr string, port string) *InputConfig {
	ic := &InputConfig{
		Address: addr,
		Port: port,
	}

	return ic
}

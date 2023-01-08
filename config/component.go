package config

type componentConfigOptions interface{
	Parse() struct{}
}

type ComponentConfigOptions struct {
	ReceiverConfigOptions map[string]*ReceiverConfigOptions `mapstructure:"receivers"`
	ParserConfigOptions   map[string]*ParserConfigOptions   `mapstructure:"parsers"`
	SenderConfigOptions   map[string]*SenderConfigOptions   `mapstructure:"senders"`
}

type ComponentConfigs struct {
	ReceiverConfigs map[string]*ReceiverConfig `yaml:"receivers"`
	ParserConfigs   map[string]*ParserConfig   `yaml:"parsers"`
	SenderConfigs   map[string]*SenderConfig   `yaml:"senders"`
}

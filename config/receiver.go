package config

type ReceiverConfig struct {
	HttpConfig
}

func NewReceiverConfig(receiverConfOpts *HttpConfigOptions) (*ReceiverConfig, error) {
	return newReceiverConfig(receiverConfOpts)
}

func newReceiverConfig(receiverConfOpts *HttpConfigOptions) (*ReceiverConfig, error) {
	receiverConf := NewHttpConfig(receiverConfOpts)
	return &ReceiverConfig{
			HttpConfig: receiverConf,
	},
	nil

}

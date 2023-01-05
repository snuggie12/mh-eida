package config

type ReceiverConfigOptions struct {
	HttpConfigOptions `mapstructure:",squash"`
	Name              string `mapstructure:"name"`
	Path              string `mapstructure:"path"`
	Type              string `mapstructure:"type"`
}

type ReceiverConfig struct {
	HttpConfig
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Type string `yaml:"type"`
}

func NewReceiverConfig(receiverConfOpts *ReceiverConfigOptions) (*ReceiverConfig, error) {
	return newReceiverConfig(receiverConfOpts)
}

func newReceiverConfig(receiverConfOpts *ReceiverConfigOptions) (*ReceiverConfig, error) {
	receiverHttpConf := NewHttpConfig(&receiverConfOpts.HttpConfigOptions)
	return &ReceiverConfig{
			HttpConfig: receiverHttpConf,
			Name:       receiverConfOpts.Name,
			Path:       receiverConfOpts.Path,
			Type:       receiverConfOpts.Type,
		},
		nil

}

func ParseRecieverConfigs(receiverConfsOpts []*ReceiverConfigOptions) ([]ReceiverConfig) {
	receiverConfs := make([]ReceiverConfig, 0)
	for _, receiverConfOpts := range(receiverConfsOpts) {
		receiverConf, err := NewReceiverConfig(receiverConfOpts)
		if err != nil {
			continue
		}

		receiverConfs = append(receiverConfs, *receiverConf)
	}
	
	return receiverConfs
}

func FriendlyReceiverConfigs(receiverConfs []ReceiverConfig) []ReceiverConfig {
	friendlyReceiverConfs := make([]ReceiverConfig, 0)
	for _, receiverConf := range receiverConfs {
		if len(receiverConf.Address) == 0 {
			receiverConf.Address = "0.0.0.0"
		}
		friendlyReceiverConfs = append(friendlyReceiverConfs, receiverConf)
	}

	return friendlyReceiverConfs
}

var typeToGenericType = map[string]string{
	"argocd":         "http",
	"launchDarkly":   "http",
	"genericWebhook": "http",
	"kubernetes":     "pull",
}

func GetGenericType(specificType string) (string) {

	receiverType := typeToGenericType[specificType]
	if receiverType == "" {
		receiverType = "none"
	}

	return receiverType
}

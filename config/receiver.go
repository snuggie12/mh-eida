package config

import "fmt"

type ReceiverConfigOptions struct {
	HttpConfigOptions `mapstructure:"http,omitempty" json:"http,omitempty"`
	TcpConfigOptions  `mapstructure:"tcp,omitempty" json:"tcp,omitempty"`
	Type              string `mapstructure:"type"`
}

type ReceiverConfig struct {
	HttpConfig `yaml:"http,omitempty" json:"http,omitempty"`
	Name       string `yaml:"name"`
	TcpConfig  `yaml:"tcp,omitempty" json:"tcp,omitempty"`
	Type       string `yaml:"type" json:"type"`
}

//Implements component interfaces parse method
func (recConfOpts *ReceiverConfigOptions) Parse(compName string) (*ReceiverConfig, error) {
	err := recConfOpts.validate()
	if err != nil {
		return nil, err
	}

	if recConfOpts.HttpConfigOptions.HttpPort != 0 {
		return &ReceiverConfig{
			HttpConfig: NewHttpConfig(&recConfOpts.HttpConfigOptions),
			Name:       compName,
			TcpConfig:  TcpConfig{},
			Type:       recConfOpts.Type,
		}, nil
	}

	if recConfOpts.TcpConfigOptions.TcpPort != 0 {
		return &ReceiverConfig{
			HttpConfig: HttpConfig{},
			Name:       compName,
			TcpConfig:  NewTcpConfig(&recConfOpts.TcpConfigOptions),
			Type:       recConfOpts.Type,
		}, nil
	}

	return nil, fmt.Errorf("Specific receiver type not configured such as http or tcp")
}

func (recConfOpts *ReceiverConfigOptions) validate() error {
	//TODO: actually validate things like a single receiver can't be both http and tcp
	return nil
}

//Makes it easier to
func FriendlyReceiverConfigs(ComponentConfs *ComponentConfigs) map[string]ReceiverConfig {
	resultReceiverConfs := make(map[string]ReceiverConfig, 0)
	for receiverName, receiverConf := range ComponentConfs.ReceiverConfigs {

		//Friendly http address if empty
		if receiverConf.HttpConfig.HttpPort != "" {
			if receiverConf.HttpAddress == "" {
				receiverConf.HttpConfig.HttpAddress = "0.0.0.0"
			}
		}

		//Friendly tcp address if empty
		if receiverConf.TcpConfig.TcpPort != "" {
			if receiverConf.TcpAddress == "" {
				receiverConf.TcpConfig.TcpAddress = "0.0.0.0"
			}
		}

		resultReceiverConfs[receiverName] = *receiverConf
	}
	return resultReceiverConfs
}

var typeToGenericType = map[string]string{
	"argocd":         "http",
	"launchDarkly":   "http",
	"genericWebhook": "http",
	"kubernetes":     "pull",
}

func GetGenericType(specificType string) string {

	receiverType := typeToGenericType[specificType]
	if receiverType == "" {
		receiverType = "none"
	}

	return receiverType
}

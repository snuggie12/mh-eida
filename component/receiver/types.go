package receiver

import "snuggie12/eida/config"

type Receiver struct {
	Config config.ReceiverConfig
}

type KubernetesReceiver struct {
	//TODO: Watch kubernetes events
}

type PullReceiver struct {
	//TODO: Generic receiver that fetches things instead of actually receiving them
}

type HttpReceiver struct {
	Address string
	Path string
	Port string
	MetricsServer *MetricsServer
}

type MetricsServer struct {}

type ParsedHttpReceiverConfig struct {
	Address string
	PathsToFuncs []PathToFunc
	Port string
}

type PathToFunc struct {
	Path string
	Function func()
}

package metrics

import "time"

//ReceiverRequestsCommonInfo contains shared labels for receiver requests
type ReceiverRequestsCommonInfo struct {
	ReceiverName string
	ReceiverPort string
	StatusCode   string
}

//ReceiverRequestsCounterInfo contains labels for the receiver requests counter
type ReceiverRequestsCounterInfo struct {
	ReceiverRequestsCommonInfo
}

//ReceiverRequestsHistogramInfo contains labels for the receiver requests histogram
type ReceiverRequestsHistogramInfo struct {
	ReceiverRequestsCommonInfo
	Duration time.Duration
}

func (info ReceiverRequestsCounterInfo) ParseInfo() (string, string, string) {
	return info.ReceiverName, info.ReceiverPort, info.StatusCode
}

func (info ReceiverRequestsHistogramInfo) ParseInfo() (string, string, string) {
	return info.ReceiverName, info.ReceiverPort, info.StatusCode
}

package metrics

import (
	"strconv"
	"time"
)

// ReceiverRequestsCommonInfo contains shared labels for receiver requests
type ReceiverRequestsCommonInfo struct {
	ReceiverName string
	ReceiverPort string
	StatusCode   string
}

// ReceiverRequestsCounterInfo contains labels for the receiver requests counter
type ReceiverRequestsCounterInfo struct {
	ReceiverRequestsCommonInfo
}

// ReceiverRequestsHistogramInfo contains labels for the receiver requests histogram
type ReceiverRequestsHistogramInfo struct {
	ReceiverRequestsCommonInfo
	Duration time.Duration
}

func NewReceiverRequestsCounterInfo(name string, port string, statusCode int) *ReceiverRequestsCounterInfo {
	return &ReceiverRequestsCounterInfo{
		ReceiverRequestsCommonInfo{
			ReceiverName: name,
			ReceiverPort: port,
			StatusCode:   strconv.Itoa(statusCode),
		},
	}
}

func NewReceiverRequestsHistogramInfo(name string, port string, statusCode int, duration time.Duration) *ReceiverRequestsHistogramInfo {
	return &ReceiverRequestsHistogramInfo{
		ReceiverRequestsCommonInfo{
			ReceiverName: name,
			ReceiverPort: port,
			StatusCode:   strconv.Itoa(statusCode),
		},
		duration,
	}
}

func (info ReceiverRequestsCounterInfo) ParseInfo() (string, string, string) {
	return info.ReceiverName, info.ReceiverPort, info.StatusCode
}

func (info ReceiverRequestsHistogramInfo) ParseInfo() (string, string, string) {
	return info.ReceiverName, info.ReceiverPort, info.StatusCode
}

package event

type UnparsedData struct {
	ReceiverName string
	ReceiverType string
}

type EventParser interface {
	ParseIntoEvent() *Event
}

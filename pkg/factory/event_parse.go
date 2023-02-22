package factory

import "context"

// Starter allows for the composition of various ways to start a component
type EventParser interface {
	ParseEvent(context.Context) error
}

// ParseEventFunction is how most components should interface with the factory package
// when it comes to a parser translating to an event
type ParseEventFunction func(context.Context) error

// A parser should create a function which returns their version of the ParseEvent function and
// then call ParseEvent on it to implement the interface
func (f ParseEventFunction) ParseEvent(ctx context.Context) error {
	if f == nil {
		return nil
	}
	return f(ctx)
}

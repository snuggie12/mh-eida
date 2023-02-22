package factory

import "context"

// Starter allows for the composition of various ways to start a component
type Starter interface {
	Start(ctx context.Context) error
}

// StartFunction is how most components should interface with the factory package
// when it comes to starting a component
type StartFunction func(context.Context) error

// A component should create a function which returns their version of the start function and
// then call Start on it to implement the interface
func (f StartFunction) Start(ctx context.Context) error {
	if f(ctx) == nil {
		return nil
	}
	return f(ctx)
}

// Stopper allows for the composition of various ways to stop a component
type Stopper interface {
	Stop() error
}

// StopFunction is how most components should interface with the factory package
// when it comes to stoping a component
type StopFunction func(context.Context) error

// A component should create a function which returns their version of the stop function and
// then call Stop on it to implement the interface
func (f StopFunction) Stop(ctx context.Context) error {
	if f(ctx) == nil {
		return nil
	}
	return f(ctx)
}

// A composition of Starter and Stopper
type StartStopper interface {
	Starter
	Stopper
}

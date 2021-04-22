package eventmanager

import (
	"github.com/Confialink/wallet-notifications/internal/service"
)

type Dispatcher bool

// DispatcherContract is dispatcher interface
type DispatcherContract interface {
	Dispatch(eventName string, data service.CallData, subscribers []service.Subscriber)
}

// Dispatch dispatches the event in sequential
func (d *Dispatcher) Dispatch(eventName string, data service.CallData, subscribers []service.Subscriber) {
	dispatch(subscribers, eventName, data)
}

// dispatch dispatches the event across all the subscribers
func dispatch(subscribers []service.Subscriber, eventName string, data service.CallData) {
	for _, subscriber := range subscribers {
		subscriber.Call(&data)
	}
}

// NewDispatcher is factory method for the event manager dispatcher
func NewDispatcher() *Dispatcher {
	return new(Dispatcher)
}

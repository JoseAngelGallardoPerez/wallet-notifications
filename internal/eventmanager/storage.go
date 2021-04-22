package eventmanager

import "github.com/Confialink/wallet-notifications/internal/service"

// Storage the storage interface
type Storage interface {
	Subscribers(string) []service.Subscriber
	Attach(string, service.Subscriber)
	Detach(string, service.Subscriber)
}

// Memory storage structure
type Memory struct {
	storage map[string][]service.Subscriber
}

// Attach attaches a subscriber for an event
func (m *Memory) Attach(eventName string, subscriber service.Subscriber) {
	subscribers := m.storage[eventName]
	m.storage[eventName] = append(subscribers, subscriber)
}

// Subscribers retrieve all subscribers of a given event
func (m *Memory) Subscribers(eventName string) []service.Subscriber {
	sr := m.storage[eventName]
	return sr
}

// Detach de attaches a subscriber
func (m *Memory) Detach(eventName string, subscriber service.Subscriber) {
	var key int
	for k, v := range m.storage[eventName] {
		if subscriber == v {
			key = k
			break
		}
	}
	m.storage[eventName] = append(m.storage[eventName][:key], m.storage[eventName][1+key:]...)
}

// NewMemoryStorage is factory method for the storage structure
func NewMemoryStorage() *Memory {
	m := make(map[string][]service.Subscriber)
	return &Memory{
		m,
	}
}

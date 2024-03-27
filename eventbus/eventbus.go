package eventbus

import (
	"errors"
	"sync"
)

type (
	Topic = string

	Subscriber[T any] struct {
		Name   string
		Notify func([]T)
	}

	EventBus[T any] struct {
		subscribers map[Topic]map[string]Subscriber[T]
		mu          sync.Mutex
	}
)

func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{
		subscribers: make(map[Topic]map[string]Subscriber[T]),
	}
}

func (eb *EventBus[T]) Subscribe(topic Topic, subscribers ...Subscriber[T]) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if eb.subscribers[topic] == nil {
		eb.subscribers[topic] = make(map[string]Subscriber[T], 0)
	}
	for _, subscriber := range subscribers {
		eb.subscribers[topic][subscriber.Name] = subscriber
	}
}

func (eb *EventBus[T]) Unsubscribe(topic Topic, subscriberNames ...string) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	subscribers, ok := eb.subscribers[topic]
	if !ok {
		return errors.New("topic not found")
	}

	for _, subscriberName := range subscriberNames {
		for name := range subscribers {
			if name == subscriberName {
				delete(subscribers, name)
			}
		}
	}
	return nil
}

func (eb *EventBus[T]) UnsubscribeAll() {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.subscribers = make(map[Topic]map[string]Subscriber[T])
}

// Publish sends an event to all subscribers of a specific event type.
func (eb *EventBus[T]) Publish(topic Topic, events []T) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	subscribers, ok := eb.subscribers[topic]
	if !ok {
		return errors.New("topic not found")
	}
	for _, subscriber := range subscribers {
		subscriber.Notify(events)
	}
	return nil
}

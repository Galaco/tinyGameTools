package tinygametools

import (
	"github.com/pkg/errors"
	"time"
)

var (
	// ErrorAlreadyUnsubscribed
	ErrorAlreadyUnsubscribed = errors.New("already unsubscribed, or never subscribed")
)

// EventName
type EventName string

// Event
type Event interface {
	// Type returns the event identifier
	Type() EventName
	// Message returns the payload of the event. Its up the the receiver to know
	// the intended format
	Message() interface{}
}

// EventManager provides a lightweight pub/sub utility
// for passing events around.
type EventManager struct {
	listeners   map[EventName][]*func(event Event)
	queuedEvent map[time.Duration]Event
}

// Subscribe takes a callback function and bind it with an EventName.
// When an event of that EventName is published, the callback will be executed.
func (service *EventManager) Subscribe(eventName EventName, callback *func(event Event)) error {
	if service.listeners[eventName] == nil {
		service.listeners[eventName] = []*func(event Event){callback}
	} else {
		service.listeners[eventName] = append(service.listeners[eventName], callback)
	}

	return nil
}

// Unsubscribe unassociated a particular callback from an EventName
func (service *EventManager) Unsubscribe(eventName EventName, callback *func(event Event)) error {
	if service.listeners[eventName] == nil {
		return ErrorAlreadyUnsubscribed
	}

	for idx, c := range service.listeners[eventName] {
		if c == callback {
			if idx == 0 {
				if len(service.listeners[eventName]) == 1 {
					service.listeners[eventName] = []*func(event Event){}
				} else {
					service.listeners[eventName] = service.listeners[eventName][1:]
				}
			} else {
				service.listeners[eventName] = append(service.listeners[eventName][:idx], service.listeners[eventName][idx+1:]...)
			}

			return nil
		}
	}

	return ErrorAlreadyUnsubscribed
}

// Publish takes an Event and executes all callbacks associated with
// the passed event.Type()
func (service *EventManager) Publish(event Event) error {
	if service.listeners[event.Type()] == nil {
		return nil
	}
	for _, callback := range service.listeners[event.Type()] {
		(*callback)(event)
	}

	return nil
}

// NewEventManager returns a new EventManager
func NewEventManager() *EventManager {
	return &EventManager{
		listeners: map[EventName][]*func(event Event){},
	}
}

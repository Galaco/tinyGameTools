package tinygametools

import (
	"github.com/pkg/errors"
)

var (
	// ErrorAlreadyUnsubscribed
	ErrorAlreadyUnsubscribed = errors.New("already unsubscribed, or never subscribed")
)

// EventName
type EventName string

type eventListener struct {
	owner interface{}
	callback func (Event)
}

func newEventListener(owner interface{}, callback func(Event)) eventListener {
	return eventListener{
		owner: owner,
		callback: callback,
	}
}

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
	listeners   map[EventName][]eventListener
}

// Subscribe takes a callback function and bind it with an EventName.
// When an event of that EventName is published, the callback will be executed.
func (service *EventManager) Subscribe(eventName EventName, callback func(event Event), owner interface{}) error {
	if service.listeners[eventName] == nil {
		service.listeners[eventName] = []eventListener{newEventListener(owner, callback)}
	} else {
		service.listeners[eventName] = append(service.listeners[eventName], newEventListener(owner, callback))
	}

	return nil
}

// Unsubscribe unassociated a particular callback from an EventName
func (service *EventManager) Unsubscribe(eventName EventName, owner interface{}) error {
	if service.listeners[eventName] == nil {
		return ErrorAlreadyUnsubscribed
	}

	for idx, listener := range service.listeners[eventName] {
		if listener.owner == owner {
			if idx == 0 {
				if len(service.listeners[eventName]) == 1 {
					service.listeners[eventName] = make([]eventListener, 0)
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
	for _, listener := range service.listeners[event.Type()] {
		listener.callback(event)
	}

	return nil
}

// NewEventManager returns a new EventManager
func NewEventManager() *EventManager {
	return &EventManager{
		listeners: map[EventName][]eventListener{},
	}
}

package tinygametools

import "testing"

type testEvent struct{}

func (evt *testEvent) Type() EventName {
	return EventName("foo")
}

func (evt *testEvent) Message() interface{} {
	return nil
}

func TestEventManager_Publish(t *testing.T) {
	sut := NewEventManager()
	evtTest := &testEvent{}
	var called bool
	evtCallback := func(event Event) {
		called = true
	}
	_ = sut.Subscribe(evtTest.Type(), evtCallback, evtTest)
	err := sut.Publish(evtTest)
	if err != nil {
		t.Error(err)
	}
	if called != true {
		t.Error("subscribed callback was never called")
	}
}

func TestEventManager_Subscribe(t *testing.T) {
	sut := NewEventManager()
	evtName := EventName("foo")
	evtCallback := func(event Event) {}
	err := sut.Subscribe(evtName, evtCallback, evtName)
	if err != nil {
		t.Error(err)
	}
	if sut.listeners[evtName] == nil {
		t.Error("Expected listener in eventManager, but no listener found")
	}
}

func TestEventManager_Unsubscribe(t *testing.T) {
	sut := NewEventManager()
	evtName := EventName("foo")
	evtCallback := func(event Event) {}
	err := sut.Subscribe(evtName, evtCallback, evtName)
	if err != nil {
		t.Error(err)
	}
	err = sut.Unsubscribe(evtName, evtName)
	if err != nil {
		t.Error(err)
	}
	if len(sut.listeners[evtName]) > 0 {
		t.Errorf("Expected 0 listeners for %s, but %d found", evtName, len(sut.listeners[evtName]))
	}
}

func TestNewEventManager(t *testing.T) {
	sut := NewEventManager()

	if sut == nil {
		t.Errorf("nil returned, but expected object")
	}
}

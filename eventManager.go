package crump8

// EventManager Handles execution events on the chip8
type EventManager struct {
	Stop   chan struct{}
	Pause  chan struct{}
	Resume chan struct{}
}

// NewEventManager returns a new event manager
func NewEventManager() *EventManager {
	return &EventManager{
		Stop:   make(chan struct{}),
		Pause:  make(chan struct{}),
		Resume: make(chan struct{}),
	}
}

package broadcast

type Broadcast struct {
	input chan interface{}
	reg   chan chan<- interface{}
	unreg chan chan<- interface{}

	outputs map[chan<- interface{}]bool
}

// The Broadcaster interface describes the main entry points to
// broadcasters.
type Broadcaster interface {
	// Register a new channel to receive broadcasts
	Register(chan<- interface{})
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- interface{})
	// Shut this broadcaster down.
	Close() error
	// Submit a new object to all subscribers
	Submit(interface{})
	// Try Submit a new object to all subscribers return false if input chan is fill
	TrySubmit(interface{}) bool
}

func (b *Broadcast) broadcast(m interface{}) {
	for ch := range b.outputs {
		ch <- m
	}
}

func (b *Broadcast) run() {
	for {
		select {
		case m := <-b.input:
			b.broadcast(m)
		case ch, ok := <-b.reg:
			if ok {
				b.outputs[ch] = true
			} else {
				return
			}
		case ch := <-b.unreg:
			delete(b.outputs, ch)
		}
	}
}

// NewBroadcaster creates a new Broadcast with the given input
// channel buffer length.
func NewBroadcaster(buflen int) Broadcaster {
	b := &Broadcast{
		input:   make(chan interface{}, buflen),
		reg:     make(chan chan<- interface{}),
		unreg:   make(chan chan<- interface{}),
		outputs: make(map[chan<- interface{}]bool),
	}

	go b.run()

	return b
}

func (b *Broadcast) Register(newch chan<- interface{}) {
	b.reg <- newch
}

func (b *Broadcast) Unregister(newch chan<- interface{}) {
	b.unreg <- newch
}

func (b *Broadcast) Close() error {
	close(b.reg)
	close(b.unreg)
	return nil
}

// Submit an item to be broadcast to all listeners.
func (b *Broadcast) Submit(m interface{}) {
	if b != nil {
		b.input <- m
	}
}

// TrySubmit attempts to submit an item to be broadcast, returning
// true iff it the item was broadcast, else false.
func (b *Broadcast) TrySubmit(m interface{}) bool {
	if b == nil {
		return false
	}
	select {
	case b.input <- m:
		return true
	default:
		return false
	}
}
package concurrency

import "sync"

type Signaler struct {
	mu sync.Mutex
	ch chan struct{}
}

func NewSignaler() *Signaler {
	return &Signaler{
		ch: make(chan struct{}),
	}
}

func (s *Signaler) Signal() {
	s.mu.Lock()
	defer s.mu.Unlock()

	close(s.ch)
	s.ch = make(chan struct{})
}

func (s *Signaler) Wait() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.ch
}

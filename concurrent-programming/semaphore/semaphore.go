package semaphore

import "sync"

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Semaphore) Acquire() {
	s.cond.L.Lock()

	for s.permits <= 0 {
		s.cond.Wait()
	}
	s.permits--

	s.cond.L.Unlock()
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()

	s.permits++
	s.cond.Signal()

	s.cond.L.Unlock()
}

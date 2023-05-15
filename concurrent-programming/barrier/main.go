package barrier

import "sync"

type Barrier struct {
	size        int
	waitCounter int
	cond        *sync.Cond
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		size:        size,
		waitCounter: 0,
		cond:        sync.NewCond(&sync.Mutex{}),
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCounter++
	if b.waitCounter == b.size {
		b.waitCounter = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.cond.L.Unlock()
}

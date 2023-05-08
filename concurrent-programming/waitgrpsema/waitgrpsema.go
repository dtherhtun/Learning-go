package waitgrpsema

import "github.com/dtherhtun/Learning-go/concurrent-programming/semaphore"

type WaitGrp struct {
	sema *semaphore.Semaphore
}

func NewWaitGrp(size int) *WaitGrp {
	return &WaitGrp{sema: semaphore.NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	wg.sema.Release()
}

package tryreadwritemutex

import "sync"

type ReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) TryWriteLock() bool {
	return rw.globalLock.TryLock()
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	if rw.readersLock.TryLock() {
		globalSuccess := true
		if rw.readersCounter == 0 {
			globalSuccess = rw.globalLock.TryLock()
		}
		if globalSuccess {
			rw.readersCounter++
		}
		rw.readersLock.Unlock()
		return globalSuccess
	} else {
		return false
	}
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

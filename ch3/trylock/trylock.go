package trylock

// 使用 channel 实现一个 trylock

type MyLock struct {
	ch chan struct{}
}

func NewMyLock() *MyLock {
	lock := &MyLock{
		ch: make(chan struct{}, 1),
	}
	lock.ch <- struct{}{}
	return lock
}

func (l *MyLock) Lock() {
	<-l.ch
}

func (l *MyLock) Unlock() {
	select {
	case <-l.ch:
	default:
		panic("unlock unlocked MyLock")
	}
}

func (l *MyLock) TryLock() bool {
	select {
	case <-l.ch:
		return true
	default:
		return false
	}
}

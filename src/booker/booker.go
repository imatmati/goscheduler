package booker

import (
	"sync"
)

type Booker interface {
	Book(pos int64) func()
	//Release(pos int64)
}

type BookerImpl struct {
	booking     map[int64]sync.Locker
	bookinglock sync.Locker
}

func (b *BookerImpl) Book(pos int64) func() {
	b.bookinglock.Lock()
	var (
		lock sync.Locker
		ok   bool
	)
	if lock, ok = b.booking[pos]; !ok {
		lock = &sync.Mutex{}
		b.booking[pos] = lock
	}
	lock.Lock()
	b.bookinglock.Unlock()
	return b.makeReleaseFunc(pos)
}

func (b *BookerImpl) makeReleaseFunc(pos int64) func() {
	return func() {
		b.bookinglock.Lock()
		b.booking[pos].Unlock()
		b.booking[pos] = nil
		b.bookinglock.Unlock()
	}
}

/*func (b *BookerImpl) Release(pos int64) {
	b.bookinglock.Lock()
	var (
		lock sync.Locker
		ok   bool
	)
	if lock, ok = b.booking[pos]; ok {
		lock.Unlock()
		b.booking[pos] = nil
	}
	b.bookinglock.Unlock()
}*/

func New() Booker {
	return &BookerImpl{make(map[int64]sync.Locker), &sync.Mutex{}}
}

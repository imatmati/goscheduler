package booker

import (
	"sync"
)

type Booker interface {
	Channel() (chan Book, chan Release)
	Start()
}

type Book struct {
	pos      int64
	response chan bool
}

type Release Book

type Impl struct {
	booking map[int64]sync.Locker
	book    chan Book
	release chan Release
}

func (b *Impl) Channel() (chan Book, chan Release) {
	return b.book, b.release
}

func (b *Impl) Start() {
	go func() {

		for {
			select {
			case request := <-b.book:
				newLock := &sync.Mutex{}
				newLock.Lock()
				var (
					lock sync.Locker
					ok   bool
				)
				if lock, ok = b.booking[request.pos]; ok {
					lock.Lock()
				}

				b.booking[request.pos] = newLock
				request.response <- true

			case release := <-b.release:
				var (
					lock sync.Locker
					ok   bool
				)
				if lock, ok = b.booking[release.pos]; ok {
					lock.Unlock()
					b.booking[release.pos] = nil
				}
				release.response <- ok

			}
		}
	}()
}

func New() Booker {
	return &Impl{make(map[int64]sync.Locker), make(chan Book, 10), make(chan Release, 10)}
}

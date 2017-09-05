package booker

import (
	"fmt"
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
	fmt.Println("BookerImpl lock access ")
	b.bookinglock.Lock()
	var (
		lock sync.Locker
		ok   bool
	)
	fmt.Println("BookerImpl after lock access")
	if lock, ok = b.booking[pos]; !ok {
		lock = &sync.Mutex{}
		b.booking[pos] = lock
	}
	fmt.Printf("BookerImpl before lock of node %v\n", lock)
	lock.Lock()
	fmt.Println("BookerImpl after lock of node")
	b.bookinglock.Unlock()
	fmt.Println("BookerImpl unlock access")
	return b.makeReleaseFunc(pos)
}

func (b *BookerImpl) makeReleaseFunc(pos int64) func() {
	return func() {
		fmt.Println("BookerImpl stage 1")
		b.bookinglock.Lock()
		fmt.Println("BookerImpl stage 2")
		b.booking[pos].Unlock()
		fmt.Println("BookerImpl stage 3")
		delete(b.booking, pos)
		fmt.Println("BookerImpl stage 4")
		b.bookinglock.Unlock()
		fmt.Println("BookerImpl stage 5")
	}
}

func New() Booker {
	return &BookerImpl{make(map[int64]sync.Locker), &sync.Mutex{}}
}

package booker

import (
	"sync"
)

type Booker struct {
	booking       map[int64]*sync.Mutex
	bookingChan   chan Book
	releasingChan chan Release
}
type Book struct {
	pos      int64
	response chan *sync.Mutex
}

type Release struct {
	pos      int64
	response chan bool
}

func (b Booker) book(pos int64) {
	response := make(chan *sync.Mutex)
	b.bookingChan <- Book{pos, response}
	lock := <-response
	lock.Lock()

}

func (b Booker) release(pos int64) bool {
	response := make(chan bool)
	b.releasingChan <- Release{pos, response}
	return <-response
}

func (b Booker) Start() {
	go func() {

		for {
			select {
			case request := <-b.bookingChan:
				var (
					lock *sync.Mutex
					ok   bool
				)
				if lock, ok = b.booking[request.pos]; !ok || lock == nil {
					lock = new(sync.Mutex)
					b.booking[request.pos] = lock
				}
				request.response <- lock
			case release := <-b.releasingChan:
				var (
					lock *sync.Mutex
					ok   bool
				)
				if lock, ok = b.booking[release.pos]; ok && lock != nil {

					lock.Unlock()
					b.booking[release.pos] = nil
				}
				release.response <- true

			}
		}
	}()
}

func New() *Booker {
	return &Booker{make(map[int64]*sync.Mutex),
		make(chan Book),
		make(chan Release)}
}

func NewStarted() *Booker {
	booker := &Booker{make(map[int64]*sync.Mutex),
		make(chan Book),
		make(chan Release)}
	booker.Start()
	return booker
}

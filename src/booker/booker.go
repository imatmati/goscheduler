package booker

import (
	"fmt"
	"sync"
)

type Booker struct {
	booking       map[int64]sync.Locker
	bookingChan   chan Book
	releasingChan chan Release
}
type Book struct {
	pos      int64
	response chan bool
	id       string
}

type Release Book

func (b Booker) book(pos int64, id string) bool {
	response := make(chan bool)
	fmt.Printf("Sending book request for %s\n", id)
	b.bookingChan <- Book{pos, response, id}
	fmt.Printf("Book request sent for %s\n", id)
	return <-response
}

func (b Booker) release(pos int64, id string) bool {
	response := make(chan bool)
	fmt.Printf("Sending release request for %s\n", id)
	b.releasingChan <- Release{pos, response, id}
	fmt.Printf("Release request sent for %s\n", id)
	return <-response
}

func (b Booker) Start() {
	go func() {

		for {
			select {
			case request := <-b.bookingChan:
				fmt.Printf("Request from %s\n", request.id)
				if lock, ok := b.booking[request.pos]; ok && lock != nil {
					//Someone acquired the node, waiting for my turn
					fmt.Println("Waiting for lock")
					lock.Lock()
					fmt.Println("Lock acquired")
				} else {
					// Node is accessible, save it from others.
					newLock := &sync.Mutex{}
					fmt.Println("Creating new lock")
					newLock.Lock()
					b.booking[request.pos] = newLock
					fmt.Println("Lock put")
				}
				request.response <- true
			case release := <-b.releasingChan:
				fmt.Printf("Release from %s\n", release.id)
				if lock, ok := b.booking[release.pos]; ok {
					fmt.Println("Lock present")
					lock.Unlock()
					fmt.Println("Lock acquired")
					b.booking[release.pos] = nil
					fmt.Println("Lock deleted")
				}
				release.response <- true

			}
		}
	}()
}

func New() *Booker {
	return &Booker{make(map[int64]sync.Locker),
		make(chan Book),
		make(chan Release)}
}

func NewStarted() *Booker {
	booker := &Booker{make(map[int64]sync.Locker),
		make(chan Book),
		make(chan Release)}
	booker.Start()
	return booker
}

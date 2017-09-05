package booker

import (
	"fmt"
	"sync"
	"testing"
)

func TestBookAPosition(t *testing.T) {
	booker := &BookerImpl{make(map[int64]sync.Locker), &sync.Mutex{}}
	release := booker.Book(0)
	if booker.booking[0] == nil {
		t.Error("No booking realised for position 0")
	}
	release()
	if booker.booking[0] != nil {
		t.Error("Release failed for position 0")
	}
}

func TestConcurrentBookAPosition(t *testing.T) {
	booker := &BookerImpl{make(map[int64]sync.Locker), &sync.Mutex{}}
	synchro := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	var a int

	go func() {
		fmt.Println("Goroutine 2 : wait for goroutine 1")
		<-synchro
		releaseConcurrent := booker.Book(0)
		fmt.Println("Goroutine 2 : got 0 lock")
		a = -1
		releaseConcurrent()
		fmt.Println("Goroutine 2 : release node after changing a")
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine 1")
		releaseConcurrent := booker.Book(0)
		fmt.Println("Goroutine 1 : got 0 lock")
		synchro <- 0
		fmt.Println("Goroutine 1 : signaled other goroutine")
		a = 1
		releaseConcurrent()
		fmt.Println("Goroutine 1 : released node after changing a")
		wg.Done()
	}()

	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
	if a != -1 {
		t.Errorf("Wrong concurrent access to same position")
	}

}

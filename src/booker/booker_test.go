package booker

import (
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
	syncro := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	var a int
	go func() {
		releaseConcurrent := booker.Book(0)
		syncro.Signal()
		a = 1
		releaseConcurrent()
		wg.Done()
	}()

	go func() {
		syncro.Wait()
		releaseConcurrent := booker.Book(0)
		a = -1
		releaseConcurrent()
		wg.Done()
	}()
	wg.Wait()
	if a != -1 {
		t.Errorf("Wrong concurrent access to same position")
	}

}

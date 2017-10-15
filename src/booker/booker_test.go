package booker

import (
	"testing"
)

func TestBookAPosition(t *testing.T) {
	booker := NewStarted()
	booker.book(0)

	if booker.booking[0] == nil {
		t.Error("No booking realised for position 0")
	}

}

func TestReleaseAPosition(t *testing.T) {
	booker := New()
	booker.Start()
	booker.book(0)
	if ok := booker.release(0); !ok {
		t.Error("No release realised for position 0")
	}
}

func TestConcurrentNodeAccess2(t *testing.T) {
	booker := New()
	booker.Start()
	var (
		com = make(chan bool)
		end = make(chan bool)
	)
	a := 0
	go func() {

		// Waiting for func 2 to acquire node 0
		<-com
		//Trying to acquire node 0
		booker.book(0)

		// Node 0 acquired after func 2 released it, a set to 1
		a = 2
		//Modifying  a after func 2
		booker.release(0)
		//Signaling main goroutine to proceed.
		end <- true
	}()

	go func() {
		//Acquiring node 0
		booker.book(0)
		// Signal func 1 node 0 acquired
		com <- true
		// Modifying a
		a = 1
		// Releasing node for func 1
		booker.release(0)

	}()

	<-end
	if a != 2 {
		t.Error("No concurrent serialisation realised for position 0")
	}
}

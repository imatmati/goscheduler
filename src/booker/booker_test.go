package booker

import (
	"fmt"
	"sync"
	"testing"
)

func TestBookAPosition(t *testing.T) {
	booker := New()
	booker.Start()
	booker.book(0, "TestBookAPosition")

	if booker.booking[0] == nil {
		t.Error("No booking realised for position 0")
	}

}

func TestReleaseAPosition(t *testing.T) {
	booker := New()
	booker.Start()
	booker.book(0, "TestReleaseAPosition")
	if ok := booker.release(0, "TestReleaseAPosition"); !ok {
		t.Error("No release realised for position 0")
	}
}

func TestConcurrentNodeAccess(t *testing.T) {
	booker := New()
	booker.Start()
	sig := sync.NewCond(&sync.Mutex{})
	end := sync.NewCond(&sync.Mutex{})
	com := make(chan bool)
	a := 0
	go func() {
		fmt.Println("func 1 : Locking wait")
		sig.L.Lock()
		fmt.Println("func 1 : Waiting")
		com <- true
		sig.Wait()

		fmt.Println("func 1 : unocking wait")
		//sig.L.Unlock()
		fmt.Println("func 1 : Unlocking")
		booker.book(0, "func 1")
		fmt.Println("func 1 : Acquired node 0")
		a = 2
		booker.release(0, "func 1")
		fmt.Println("func 1 : Released node 0")
		//end.L.Lock()
		end.Signal()
		//end.L.Lock()
		fmt.Println("func 1 : finished")
	}()

	go func() {
		//time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("func 2")
		booker.book(0, "func 2")
		fmt.Println("func 2 : Acquired node 0 ")
		//sig.L.Lock()
		fmt.Println("func 2 : waiting for com ")

		fmt.Println("func 2 : com received ")
		<-com
		sig.Signal()

		fmt.Println("func 2 : signal sent")
		//sig.L.Unlock()
		a = 1
		booker.release(0, "func 2")
		fmt.Println("func 2: Released node 0")
		fmt.Println("func 2 finished")
	}()

	end.L.Lock()
	end.Wait()
	end.L.Unlock()
	if a != 2 {
		t.Error("No concurrent serialisation realised for position 0")
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
		fmt.Println("1")
		// Waiting for func 2 to acquire node 0
		<-com
		fmt.Println("2")
		//Trying to acquire node 0
		booker.book(0, "func 1")
		fmt.Println("3")
		// Node 0 acquired after func 2 released it, a set to 1
		a = 2
		//Modifying  a after func 2
		booker.release(0, "func 1")
		fmt.Println("4")
		//Signaling main goroutine to proceed.
		end <- true
		fmt.Println("5")
	}()

	go func() {
		//Acquiring node 0
		fmt.Println("a")
		booker.book(0, "func 2")
		fmt.Println("b")
		// Signal func 1 node 0 acquired
		com <- true
		fmt.Println("c")
		// Modifying a
		a = 1
		// Releasing node for func 1
		booker.release(0, "func 2")
		fmt.Println("d")
	}()

	<-end
	if a != 2 {
		t.Error("No concurrent serialisation realised for position 0")
	}
}

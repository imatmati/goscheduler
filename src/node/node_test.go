package node

import (
	"fmt"
	"testing"
)

func f(v struct{ a int }) {
	fmt.Println(v.a)
}

func TestComparisonGreaterPriority(t *testing.T) {

	n := New("ok", 1)
	n2 := New("ko", 0)
	if !n.Greater(n2) {
		t.FailNow()
	}

}

func TestComparisonLessPriority(t *testing.T) {

	n := New("ok", 1)
	n2 := New("ko", 0)
	if !n2.Less(n) {
		t.FailNow()
	}

}

func TestComparisonEqualsPriority(t *testing.T) {

	n := New("ok", 0)
	n2 := New("ko", 0)
	if !n.Equals(n2) {
		t.FailNow()
	}

}

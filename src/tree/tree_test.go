package tree

import (
	"fmt"
	"log"
	"testing"
	"node"
)

func TestLengthOfFreshTree(t *testing.T) {

	length := NewDefaultTree().length()
	if length != 0 {
		t.Error(fmt.Sprintf("La longueur attendue est 0 mais n'est que %d\n", length))
	}

}

func TestCapacityOfFreshTree(t *testing.T) {

	capacity := NewDefaultTree().capacity()
	if capacity != 0 {
		t.Error(fmt.Sprintf("La capacité attendue est 0 mais n'est que %d\n", columnLength, capacity))
	}

}

func TestCapacityAfterAllocation(t *testing.T) {
	tree := NewDefaultTree()
	if err := tree.allocateNewRow(); err != nil {
		log.Print(err.Error())
		t.FailNow()
	}
	capacity := tree.capacity()
	if capacity != rowLength {
		t.Error(fmt.Sprintf("La capacité attendue est %d mais n'est que %d\n", rowLength, capacity))
	}
}

func TestCapacityOverflow(t *testing.T) {
	tree := NewTree(0)
	if err := tree.allocateNewRow(); err == nil {
		t.FailNow()
	}
}

func TestRemainderAboveAndMultipleColumnLength(t *testing.T) {
	col, row := getColumnRow(1200)
	if col != 12 || row != 0 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n",col, row))
		t.FailNow()
	}
}

func TestRemainderAboveAndNotMultipleColumnLength(t *testing.T) {
	col, row := getColumnRow(1207)
	if col != 12 || row != 7 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n",col, row))
		t.FailNow()
	}
}

func TestRemainderLessThanColumnLength(t *testing.T) {
	col, row := getColumnRow(87)
	if col != 0 || row != 87 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n",col, row))
		t.FailNow()
	}
}

func TestSetNode0 (t *testing.T) {
	tree := NewDefaultTree()
	nodetoinsert := node.New()
	err := tree.setNode(&nodetoinsert,0)
	if err != nil {
			log.Print(err.Error())
	}
	if tree.store[0][0] != &nodetoinsert {
		log.Printf("expected address %p, actual address %p",&nodetoinsert,tree.store[0][0])
		t.FailNow()
	}

}
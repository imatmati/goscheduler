package tree

import (
	"fmt"
	"log"
	"node"
	"testing"
)

func TestLengthOfFreshTree(t *testing.T) {

	length := NewDefaultTree().length()
	if length != 0 {
		t.Errorf("La longueur attendue est 0 mais n'est que %d\n", length)
	}

}

func TestCapacityOfFreshTree(t *testing.T) {

	capacity := NewDefaultTree().capacity()
	if capacity != 0 {
		t.Errorf("La capacité attendue est 0 mais n'est que %d\n", capacity)
	}

}

func TestCapacityAfterAllocation(t *testing.T) {
	tree := NewDefaultTree()
	if err := tree.allocateNewRow(); err != nil {
		log.Print(err.Error())
		t.FailNow()
	}
	capacity := tree.capacity()
	if capacity != columnLength {
		t.Errorf("La capacité attendue est %d mais n'est que %d\n", rowLength, capacity)
	}
}

func TestCapacityOverflow(t *testing.T) {
	tree := NewTree(0)
	if err := tree.allocateNewRow(); err == nil {
		t.FailNow()
	}
}

func TestRemainderAboveAndMultipleColumnLength(t *testing.T) {
	row, col := getColumnRow(120012)
	if col != 12 || row != 12 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n", col, row))
		t.FailNow()
	}
}

func TestRemainderAboveAndNotMultipleColumnLength(t *testing.T) {
	row, col := getColumnRow(100007)
	if col != 7 || row != 10 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n", col, row))
		t.FailNow()
	}
}

func TestRemainderLessThanColumnLength(t *testing.T) {
	row, col := getColumnRow(87)
	if col != 87 || row != 0 {
		log.Print(fmt.Sprintf("col : %d, row : %d\n", col, row))
		t.FailNow()
	}
}

func TestSetNode0(t *testing.T) {
	tree := NewDefaultTree()
	nodetoinsert := node.New("whatever", 0)

	if err := tree.setNode(&nodetoinsert, 0); err != nil {
		log.Print(err.Error())
		t.FailNow()
	}
	if tree.store[0][0] != &nodetoinsert {
		log.Printf("expected address %p, actual address %p", &nodetoinsert, tree.store[0][0])
		t.FailNow()
	}

}

func TestUpWithoutChange(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 0)
	child := node.New("child", 1)
	tree.push(&parent)
	tree.push(&child)
	tree.up(1)
	if allegedParent := tree.getNode(0); allegedParent != &parent {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, parent)
	}
	if allegedChild := tree.getNode(1); allegedChild != &child {
		t.Errorf("Child %v is uncorrect, %v expected", allegedChild, parent)
	}
}

func TestUpWitTwoNodes(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 2)
	child := node.New("child", 1)
	if err := tree.push(&parent); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if err := tree.push(&child); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if err := tree.up(1); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if allegedParent := tree.getNode(0); allegedParent != &child {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, child)
	}
	if allegedChild := tree.getNode(1); allegedChild != &parent {
		t.Errorf("Child %v is uncorrect, %v expected", allegedChild, parent)
	}
}

func TestUpWitThreeNodes(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 4)
	child := node.New("child", 1)
	sibling := node.New("sibling", 5)
	tree.push(&parent)
	tree.push(&sibling)
	tree.push(&child)
	tree.up(2)
	if allegedParent := tree.getNode(0); allegedParent != &child {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, parent)
	}
	if allegedFirstChild := tree.getNode(1); allegedFirstChild != &sibling {
		t.Errorf("Child %v is uncorrect, %v expected", allegedFirstChild, sibling)
	}
	if allegedSecondChild := tree.getNode(2); allegedSecondChild != &parent {
		t.Errorf("Child %v is uncorrect, %v expected", allegedSecondChild, parent)
	}

}

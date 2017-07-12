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

	if err := tree.setNode(nodetoinsert, 0); err != nil {
		log.Print(err.Error())
		t.FailNow()
	}
	if tree.store[0][0] != nodetoinsert {
		log.Printf("expected address %p, actual address %p", nodetoinsert, tree.store[0][0])
		t.FailNow()
	}

}

func TestUpWithoutChange(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 0)
	child := node.New("child", 1)
	tree.push(parent)
	tree.push(child)
	tree.up(1)
	if allegedParent, err := tree.getNode(0); allegedParent != parent || err != nil {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, parent)
	}
	if allegedChild, err := tree.getNode(1); allegedChild != child || err != nil {
		t.Errorf("Child %v is uncorrect, %v expected", allegedChild, parent)
	}
}

func TestUpWitTwoNodes(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 2)
	child := node.New("child", 1)
	if err := tree.push(parent); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if err := tree.push(child); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if err := tree.up(1); err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	if allegedParent, err := tree.getNode(0); allegedParent != child || err != nil {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, child)
	}
	if allegedChild, err := tree.getNode(1); allegedChild != parent || err != nil {
		t.Errorf("Child %v is uncorrect, %v expected", allegedChild, parent)
	}
}

func TestUpWitThreeNodes(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 4)
	child := node.New("child", 1)
	sibling := node.New("sibling", 5)
	tree.push(parent)
	tree.push(sibling)
	tree.push(child)
	tree.up(2)
	if allegedParent, err := tree.getNode(0); allegedParent != child || err != nil {
		t.Errorf("Parent %v is uncorrect, %v expected", allegedParent, parent)
	}
	if allegedFirstChild, err := tree.getNode(1); allegedFirstChild != sibling || err != nil {
		t.Errorf("Child %v is uncorrect, %v expected", allegedFirstChild, sibling)
	}
	if allegedSecondChild, err := tree.getNode(2); allegedSecondChild != parent || err != nil {
		t.Errorf("Child %v is uncorrect, %v expected", allegedSecondChild, parent)
	}

}

func TestDownWithOneNodesWithPop(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	tree.push(parent)
	var (
		popednode *node.Node
		err       error
	)

	if popednode, err = tree.pop(); err != nil {
		t.Error(err.Error())
	}
	if popednode != parent {
		t.Errorf("Poped node is not parent")
	}
	if topnode, err := tree.getNode(0); parent != topnode || err != nil {
		t.Errorf("Top node %v is incorrect, %v expected", topnode, parent)
	}
}

func TestDownWithTwoNodesWithDown(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	child := node.New("child", 4)
	tree.push(parent)
	tree.push(child)
	if err := tree.down(); err != nil {
		t.Error(err.Error())
	}
	if allegedParent, err := tree.getNode(0); parent != allegedParent || err != nil {
		t.Errorf("Parent %v is incorrect, %v expected", allegedParent, parent)
	}
}

func TestDownWithTwoNodesWithPop(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	child := node.New("child", 4)
	tree.push(parent)
	tree.push(child)
	var (
		popednode *node.Node
		err       error
	)

	if popednode, err = tree.pop(); err != nil {
		t.Error(err.Error())
	}
	if popednode != parent {
		t.Errorf("Poped node is not parent")
	}
	if topnode, err := tree.getNode(0); child != topnode || err != nil {
		t.Errorf("Top node %v is incorrect, %v expected", topnode, child)
	}
}

func TestDownWithThreeNodesWithPopAndInversion(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	child := node.New("child", 4)
	child2 := node.New("child2", 6)
	tree.push(parent)
	tree.push(child)
	tree.push(child2)
	var (
		popednode *node.Node
		err       error
	)

	if popednode, err = tree.pop(); err != nil {
		t.Error(err.Error())
	}
	if popednode != parent {
		t.Errorf("Poped node is not parent")
	}
	if topnode, err := tree.getNode(0); child != topnode || err != nil {
		t.Errorf("Top node %v is incorrect, %v expected", topnode, child)
	}
}

func TestDownWithThreeNodesWithPopWithoutInversion(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	child := node.New("child", 6)
	child2 := node.New("child2", 4)
	tree.push(parent)
	tree.push(child)
	tree.push(child2)
	var (
		popednode *node.Node
		err       error
	)

	if popednode, err = tree.pop(); err != nil {
		t.Error(err.Error())
	}
	if popednode != parent {
		t.Errorf("Poped node is not parent")
	}
	if topnode, err := tree.getNode(0); child2 != topnode || err != nil {
		t.Errorf("Top node %v is incorrect, %v expected", topnode, child)
	}
}

func TestDownWithFourNodesWithPop(t *testing.T) {
	tree := NewDefaultTree()
	parent := node.New("parent", 1)
	child := node.New("child", 6)
	child2 := node.New("child2", 4)
	child3 := node.New("child3", 8)
	tree.push(parent)
	tree.push(child)
	tree.push(child2)
	tree.push(child3)
	var (
		popednode *node.Node
		err       error
	)

	if popednode, err = tree.pop(); err != nil {
		t.Error(err.Error())
	}
	if popednode != parent {
		t.Errorf("Poped node is not parent")
	}
	if topnode, err := tree.getNode(0); child2 != topnode || err != nil {
		t.Errorf("Top node %v is incorrect, %v expected", topnode, child)
	}
}

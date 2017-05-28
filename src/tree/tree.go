package tree

import (
	"errors"
	"fmt"
	"node"
)

const (
	columnLength int = 100
	rowLength    int = 10000
)

func NewDefaultTree() *tree {
	store := make([][]*node.Node, 0, columnLength)
	return &tree{store,0}
}

func NewTree(length int) *tree {
	store := make([][]*node.Node, 0, length)
	return &tree{store,0}
}

var DefaultTree *tree = NewDefaultTree()

type tree struct {
	store [][]*node.Node
	top int
}

func (t *tree) push(node *node.Node) error {
	return t.setNode(node,t.top)
}

func (t tree) length() int {
	return len(t.store)
}

func (t tree) capacity() int {
	var capacity int
	for i := range t.store {
		capacityColumn := cap(t.store[i])
		if capacityColumn == 0 {
			break
		}
		capacity += capacityColumn
	}
	return capacity
}

func (t *tree) setNode(node *node.Node, pos int) error {
	col, row := getColumnRow(pos)
	colindex, _ := getColumnRow(t.top)
	// Beyond top and last line
	if col > colindex {
		return errors.New(fmt.Sprintf("Out of range insertion: asked %d but length is %d",))
	}
	// Row may be already allocated or not
	if col == colindex  && row == 0 {
			t.allocateNewRow() 
	} 

	t.store[col][row] = node
	return nil
}

func getColumnRow(pos int) (col , row  int) {
	col = int(pos/columnLength)
	row = pos - col * columnLength 
	return 
}

func (t tree) getNode(pos int) *node.Node {
	col, row := getColumnRow(pos)
	return t.store[col][row]
}

func (t *tree) allocateNewRow() error {

	length := len(t.store)
	capacity := cap(t.store)
	if length == capacity {
		return errors.New(fmt.Sprintf("Max capacity %d reached", capacity))
	}
	t.store = append(t.store,make([]*node.Node, rowLength))
	return nil
}

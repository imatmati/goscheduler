package tree

import (
	"fmt"
	"node"
)

const (
	columnLength int = 100
	rowLength    int = 10000
)

//NewDefaultTree creates a defaulted heap container.
func NewDefaultTree() *tree {
	store := make([][]*node.Node, 0, columnLength)
	return &tree{store, 0}
}

//NewTree creates a heap container of length specified.
func NewTree(length int) *tree {
	store := make([][]*node.Node, 0, length)
	return &tree{store, 0}
}

//DefaultTree is a provided default tree with default length.
var DefaultTree *tree = NewDefaultTree()

type tree struct {
	store [][]*node.Node
	top   int
}

func (t *tree) push(node *node.Node) error {

	if err := t.setNode(node, t.top); err != nil {
		return err
	}
	defer func() { t.top++ }()
	if err := t.up(t.top); err != nil {
		return err
	}
	return nil
}

func (t *tree) up(nodefrom int) error {

	return nil
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
		return fmt.Errorf("Out of range insertion: asked %d but length is %d", pos, t.top)
	}
	// Row may be already allocated or not
	if col == colindex && row == 0 {
		t.allocateNewRow()
	}

	t.store[col][row] = node
	return nil
}

func getColumnRow(pos int) (col, row int) {
	col = int(pos / columnLength)
	row = pos - col*columnLength
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
		return fmt.Errorf("Max capacity %d reached", capacity)
	}
	t.store = append(t.store, make([]*node.Node, rowLength))
	return nil
}

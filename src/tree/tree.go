package tree

import (
	"fmt"
	"node"
)

const (
	rowLength    int = 100
	columnLength int = 10000
)

//NewDefaultTree creates a defaulted heap container.
func NewDefaultTree() *tree {
	store := make([][]*node.Node, 0, rowLength)
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
	//top points to the next available bucket in 2d heap array.
	top int
}

func (t *tree) push(node *node.Node) error {
	// When I push, I try to insert at top position.
	if err := t.setNode(node, t.top); err != nil {
		return err
	}
	defer func() { t.top++ }()
	if err := t.up(t.top); err != nil {
		return err
	}
	return nil
}

func (t *tree) up(idnode int) error {
	l := t.length()

	if idnode >= l {
		return fmt.Errorf("Out of bound: %d requested, but length of %d", idnode, l)
	}
	nodeToUp := t.getNode(idnode)

	for {

		isRight := idnode%2 == 0
		idparent := -1
		if isRight {
			idparent = (idnode / 2) - 1
		} else {
			idparent = (idnode - 1) / 2
		}
		if idparent < 0 {
			break
		}
		nodeParent := t.getNode(idparent)
		if nodeParent.Priority > nodeToUp.Priority {

			t.setNode(nodeToUp, idparent)
			t.setNode(nodeParent, idnode)
			idnode = idparent

		} else {
			break
		}
	}
	return nil
}

func (t tree) length() int {
	return len(t.store) * columnLength
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
	// Row and column position in store 2d array of heap.
	row, col := getColumnRow(pos)
	//fmt.Printf("pos %d, col %d, row %d, top %d\n", pos, col, row, t.top)
	// Get the row and column of the nex available bucket in 2d heap array.
	rowtop, coltop := getColumnRow(t.top)
	// Inserting beyond top is forbidden. Can't store beyond next available bucket.
	// And when push is done, I simply set the node at top position.
	if row == rowtop && col > coltop || row > rowtop {
		return fmt.Errorf("Out of range insertion: asked %d but length is %d", pos, t.top)
	}
	// Row may be already allocated or not
	// As a row is fully reserved, there's room for top position as long as
	// it doesn't start a new line.
	if coltop == 0 {
		t.allocateNewRow()
	}
	t.store[row][col] = node
	return nil
}

func getColumnRow(pos int) (row, col int) {
	col = pos % columnLength
	row = int(pos / columnLength)
	return
}

func (t tree) getNode(pos int) *node.Node {
	row, col := getColumnRow(pos)
	return t.store[row][col]
}

func (t *tree) allocateNewRow() error {

	length := len(t.store)
	capacity := cap(t.store)
	if length == capacity {
		return fmt.Errorf("Max capacity %d reached", capacity)
	}
	t.store = append(t.store, make([]*node.Node, columnLength))
	return nil
}

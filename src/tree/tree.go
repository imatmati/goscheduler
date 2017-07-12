package tree

import (
	"errors"
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

// Push insert the node at the next available room in the 2d heap array.
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

func (t *tree) pop() (*node.Node, error) {
	if t.length() == 0 {
		return nil, errors.New("tree is empty")
	}
	topNode, err := t.getNode(0)
	if err != nil {
		return topNode, err
	}

	last, err := t.getNode(t.top - 1)
	if err != nil {
		return last, err
	}
	t.setNode(last, 0)

	t.down()

	return topNode, nil
}

func (t *tree) down() error {
	if t.length() <= 1 {
		return nil
	}

	topnodeID := 0
	var (
		topnode, leftnode, rightnode *node.Node
		leftID, rightID              int
		err                          error
	)
	for topnodeID > -1 {
		if topnode, err = t.getNode(topnodeID); err != nil {
			return err
		}
		if leftnode, leftID, err = t.getLeft(topnodeID); err != nil {
			return err
		}
		if rightnode, rightID, err = t.getRight(topnodeID); err != nil {
			return err
		}

		if rightnode != nil && topnode.Priority > rightnode.Priority && rightnode.Priority < leftnode.Priority {
			t.setNode(topnode, rightID)
			t.setNode(rightnode, topnodeID)
			topnodeID = rightID
		} else if leftnode != nil && topnode.Priority > leftnode.Priority {
			t.setNode(topnode, leftID)
			t.setNode(leftnode, topnodeID)
			topnodeID = leftID
		} else {
			topnodeID = -1
		}
	}
	return nil
}

func nodeMinPriority(leftnode *node.Node, leftID int, rightnode *node.Node, rightID int) (*node.Node, int) {
	if leftnode.Priority <= rightnode.Priority {
		return leftnode, leftID
	}
	return rightnode, rightID
}

// Up maintains the condition of a heap that is to say it makes any transformations needed to maintain the lowest priority node at top.
func (t *tree) up(pos int) error {
	if err := t.preconditionGet(pos); err != nil {
		return err
	}

	nodeToUp, err := t.getNode(pos)
	if err != nil {
		return err
	}

	for {

		isRight := pos%2 == 0
		idparent := -1
		if isRight {
			idparent = (pos / 2) - 1
		} else {
			idparent = (pos - 1) / 2
		}
		if idparent < 0 {
			break
		}
		nodeParent, err := t.getNode(idparent)
		if err != nil {
			return err
		}
		if nodeParent.Priority > nodeToUp.Priority {

			t.setNode(nodeToUp, idparent)
			t.setNode(nodeParent, pos)
			pos = idparent

		} else {
			break
		}
	}
	return nil
}

// Length returns the lenght of reserved space for the 2d heap array.
// By nature, it's a multiple of columLength.
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

// SetNode tries to insert the provided node at the zero indexed position in
// the 2d heap array. It can insert in any position below and equal to top index
// and takes charge of allocating new row if needed if top has gone beyond the current row of data.
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

func (t tree) preconditionGet(pos int) error {
	if pos < 0 {
		return fmt.Errorf("index %d of node is negative", pos)
	}
	if pos >= t.length() {
		return fmt.Errorf("index %d of node is beyond length of tree", pos)
	}
	return nil
}

func (t tree) getNode(pos int) (*node.Node, error) {
	if err := t.preconditionGet(pos); err != nil {
		return nil, err
	}
	row, col := getColumnRow(pos)
	return t.store[row][col], nil
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

func (t tree) getLeft(pos int) (*node.Node, int, error) {
	if err := t.preconditionGet(pos); err != nil {
		return nil, -1, err
	}
	leftID := 2*pos + 1
	if leftID >= t.length() {
		return nil, -1, nil
	}
	node, err := t.getNode(leftID)
	return node, leftID, err
}

func (t tree) getRight(pos int) (*node.Node, int, error) {
	if err := t.preconditionGet(pos); err != nil {
		return nil, -1, err
	}
	rightID := 2 * (pos + 1)
	if rightID >= t.length() {
		return nil, -1, nil
	}
	node, err := t.getNode(rightID)
	return node, rightID, err

}

package tree

import (
	"errors"
	"fmt"
	"node"
	"sync"
	"sync/atomic"
)

const (
	rowLength    int32 = 100
	columnLength int32 = 10000
)

// Tree implements the heap as a 2D array.
type Tree struct {
	store [][]*node.Node
	//top points to the next available bucket in 2d heap array.
	top *topmanager
}

//NewDefaultTree creates a defaulted heap container.
func NewDefaultTree() *Tree {
	store := make([][]*node.Node, 0, rowLength)
	return &Tree{store, &topmanager{0, &sync.Mutex{}}}
}

//NewTree creates a heap container of length specified.
func NewTree(length int) *Tree {
	store := make([][]*node.Node, 0, length)
	return &Tree{store, &topmanager{0, &sync.Mutex{}}}
}

//DefaultTree is a provided default tree with default length.
var DefaultTree *Tree = NewDefaultTree()

type topmanager struct {
	top     int64
	topLock sync.Locker
}

func (t *topmanager) acquireTop() int64 {
	t.topLock.Lock()
	defer func() { t.top++; t.topLock.Unlock() }()
	return t.top

}

func (t topmanager) getTop() int64 {
	return atomic.LoadInt64(&t.top)
}

func (t *topmanager) setTop(previousid, id int64) {
	atomic.CompareAndSwapInt64(&t.top, previousid, id)
}

func (t *Tree) setTop(node *node.Node) (top int64, err error) {
	top = t.top.acquireTop()
	if err = t.setNode(node, top); err != nil {
		return
	}
	return
}

// Push insert the node at the next available room in the 2d heap array.
func (t *Tree) Push(node *node.Node) error {
	// When I push, I try to insert at top position.
	var (
		top int64
		err error
	)
	if top, err = t.setTop(node); err != nil {
		return err
	}

	err = t.up(top)
	return err

}

// Pop extracts the root element of the heap with the less Priority
func (t *Tree) Pop() (*node.Node, error) {
	if t.length() == 0 {
		return nil, errors.New("tree is empty")
	}
	topNode, err := t.getNode(0)
	if err != nil {
		return topNode, err
	}
	last, err := t.getNode(t.top.getTop() - 1)
	if err != nil {
		return last, err
	}
	t.setNode(last, 0)

	t.down()

	return topNode, nil
}

// Down operates the descent of the replacing node of root after its extraction by Pop.
func (t *Tree) down() error {
	if t.length() <= 1 {
		return nil
	}

	var topnodeID int64 = 0
	var (
		topnode, leftnode, rightnode *node.Node
		leftID, rightID              int64
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

// Up maintains the condition of a heap that is to say it makes any transformations needed to maintain the lowest priority node at top.
func (t *Tree) up(pos int64) error {
	if err := t.preconditionGet(pos); err != nil {
		return err
	}

	nodeToUp, err := t.getNode(pos)
	if err != nil {
		return err
	}

	for {

		isRight := pos%2 == 0
		var idparent int64 = -1
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
func (t Tree) length() int64 {
	return int64(len(t.store)) * int64(columnLength)
}

// SetNode tries to insert the provided node at the zero indexed position in
// the 2d heap array. It can insert in any position below and equal to top index
// and takes charge of allocating new row if needed if top has gone beyond the current row of data.
func (t *Tree) setNode(node *node.Node, pos int64) error {
	fmt.Println("pos", pos)
	// Row and column position in store 2d array of heap.
	row, col := getColumnRow(pos)
	//fmt.Printf("pos %d, col %d, row %d, top %d\n", pos, col, row, t.top)
	// Get the row and column of the nex available bucket in 2d heap array.
	rowtop, coltop := getColumnRow(t.top.getTop())
	// Inserting beyond top is forbidden. Can't store beyond next available bucket.
	// And when push is done, I simply set the node at top position.
	if row == rowtop && col > coltop || row > rowtop {
		return fmt.Errorf("Out of range insertion: asked %d but length is %d", pos, t.top)
	}
	fmt.Println("coltop", coltop)
	// Row may be already allocated or not
	// As a row is fully reserved, there's room for top position as long as
	// it doesn't start a new line.
	if coltop == 0 {
		t.allocateNewRow()
	}
	fmt.Print("col ", col, " row ", row, "\n")
	t.store[row][col] = node
	return nil
}

func getColumnRow(pos int64) (row, col int64) {
	col = pos % int64(columnLength)
	row = pos / int64(columnLength)
	return
}

func (t Tree) preconditionGet(pos int64) error {
	if pos < 0 {
		return fmt.Errorf("index %d of node is negative", pos)
	}
	if pos >= t.length() {
		return fmt.Errorf("index %d of node is beyond length of tree", pos)
	}
	return nil
}

func (t Tree) getNode(pos int64) (*node.Node, error) {
	if err := t.preconditionGet(pos); err != nil {
		return nil, err
	}
	row, col := getColumnRow(pos)
	return t.store[row][col], nil
}

func (t *Tree) allocateNewRow() error {

	length := len(t.store)
	capacity := cap(t.store)
	if length == capacity {
		return fmt.Errorf("Max capacity %d reached", capacity)
	}
	t.store = append(t.store, make([]*node.Node, columnLength))
	return nil
}

func (t Tree) getChild(pos int64, f func(int64) int64) (*node.Node, int64, error) {
	if err := t.preconditionGet(pos); err != nil {
		return nil, -1, err
	}
	rightID := f(pos)
	if rightID >= t.length() {
		return nil, -1, nil
	}
	node, err := t.getNode(rightID)
	return node, rightID, err

}

func (t Tree) getLeft(pos int64) (*node.Node, int64, error) {
	return t.getChild(pos, func(p int64) int64 { return 2*p + 1 })
}

func (t Tree) getRight(pos int64) (*node.Node, int64, error) {
	return t.getChild(pos, func(p int64) int64 { return 2 * (p + 1) })
}

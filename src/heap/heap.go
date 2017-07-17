package heap

import (
	"node"
	t "tree"
)

//HeapManager manage the push and pop of nodes into the heap.
type Manager interface {
	Pop() (*node.Node, error)
	Push(node *node.Node) error
}

type heapManagerImpl struct {
	Manager
	tree *t.Tree
}

func (hm heapManagerImpl) Pop() (*node.Node, error) {
	return hm.tree.Pop()
}

func (hm heapManagerImpl) Push(node *node.Node) error {
	return hm.tree.Push(node)
}

//New creates a new head manager
func New() Manager {
	return heapManagerImpl{tree: t.DefaultTree}
}

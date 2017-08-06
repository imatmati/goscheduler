package node

import (
	"fmt"
)

// Node is the Implementation of node elements in the Heap.
type Node struct {
	Priority uint
	Load     interface{}
}

func (n Node) String() string {
	return fmt.Sprintf("{ Priority : %d, Load : %+v}", n.Priority, n.Load)
}

//NewEmptyNode creates an empty Node.
func NewEmptyNode() Node {
	return Node{}
}

//New creates a Node with load and priority.
func New(load interface{}, priority uint) *Node {
	return &Node{Load: load, Priority: priority}
}

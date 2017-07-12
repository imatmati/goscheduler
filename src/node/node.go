package node

import (
	"fmt"
)

// Node is the interface of nodes.
type Comparer interface {
	Less(other Node) bool
	Greater(other Node) bool
	LessOrEquals(other Node) bool
	GreaterOrEquals(other Node) bool
	Equals(other Node) bool
}

// Node is the Implementation of node elements in the Heap.
type Node struct {
	Comparer
	Priority uint
	Load     interface{}
}

// SetPriority set priority of node.
func (n Node) SetPriority(priority uint) {
	n.Priority = priority
}

// GetPriority set priority of node.
func (n Node) GetPriority() uint {
	return n.Priority
}

//Less returns true if priority of current node is less than priority of other.
func (n Node) Less(other Node) bool {
	return n.Priority < other.Priority
}

//Greater returns true if priority of current node is greater than priority of other.
func (n Node) Greater(other Node) bool {
	return n.Priority > other.Priority
}

//Equals returns true if priority of current node is equals to priority of other.
func (n Node) Equals(other Node) bool {
	return n.Priority == other.Priority
}

//LessOrEquals returns true if priority of current node is less or equals to priority of other.
func (n Node) LessOrEquals(other Node) bool {
	return n.Less(other) || n.Equals(other)
}

//GreaterOrEquals returns true if priority of current node is greater or equals to priority of other.
func (n Node) GreaterOrEquals(other Node) bool {
	return n.Greater(other) || n.Equals(other)
}

func (n Node) String() string {
	return fmt.Sprintf("{ Priority : %d, Load : %+v}", n.Priority, n.Load)
}

//NewEmptyNode creates an empty Node.
func NewEmptyNode() *Node {
	return &Node{}
}

//New creates a Node with load and priority.
func New(load interface{}, priority uint) *Node {
	return &Node{Load: load, Priority: priority}
}

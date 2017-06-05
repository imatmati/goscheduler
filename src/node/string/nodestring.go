package string

import (
	"fmt"
	"node"
)

type _node node.Node

//NodeString is a specialisation of Node for a string load.
type NodeString struct {
	Load string
	_node
}

func (t NodeString) String() string {
	return fmt.Sprintf("Priority : %d, Value : %s", t.Priority, t.Load)
}

//New creates a new Node for string load.
func New(load string, priority uint) *NodeString {
	return &NodeString{Load: load, _node: _node{Priority: priority}}

}

package string

import (
	"fmt"
	"node"
)

//NodeString is a specialisation of Node for a string load.
type NodeString struct {
	node.Comparer
	Priority uint
	Load     string
}

func (t NodeString) String() string {
	return fmt.Sprintf("Priority : %d, Value : %s", t.Priority, t.Load)
}

//New creates a new Node for string load.
func New(load string, priority uint) NodeString {
	return NodeString{Load: load, Priority: priority}

}

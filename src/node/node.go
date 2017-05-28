package node

// Comparer is the interface of comparison used by Nodes.
type Comparer interface {
	Less(other Node) bool
	Greater(other Node) bool
	Equals(other Node) bool
}

// Loader is the interface allowing getting the load part of a Node.
type Loader interface {
	SetLoad(interface{})
	GetLoad(interface{})
}

// Node is the interface of elements in the Heap.
type Node struct {
	Comparer
	Loader
}

func New() Node{
	return Node{}
}

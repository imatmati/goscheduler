package node

// Node is the interface of nodes.
type Node interface {
	Less(other Node) bool
	Greater(other Node) bool
	LessOrEquals(other Node) bool
	GreaterOrEquals(other Node) bool
	Equals(other Node) bool
	GetPriority() uint
	SetPriority(uint)
}

// Impl is the Implementation of node elements in the Heap.
type Impl struct {
	Node
	Priority uint
	Load     interface{}
}

// SetPriority set priority of node.
func (n Impl) SetPriority(priority uint) {
	n.Priority = priority
}

// GetPriority set priority of node.
func (n Impl) GetPriority() uint {
	return n.Priority
}

//Less returns true if priority of current node is less than priority of other.
func (n Impl) Less(other Node) bool {
	return n.Priority < other.GetPriority()
}

//Greater returns true if priority of current node is greater than priority of other.
func (n Impl) Greater(other Node) bool {
	return n.Priority > other.GetPriority()
}

//Equals returns true if priority of current node is equals to priority of other.
func (n Impl) Equals(other Node) bool {
	return n.Priority == other.GetPriority()
}

//LessOrEquals returns true if priority of current node is less or equals to priority of other.
func (n Impl) LessOrEquals(other Node) bool {
	return n.Less(other) || n.Equals(other)
}

//GreaterOrEquals returns true if priority of current node is greater or equals to priority of other.
func (n Impl) GreaterOrEquals(other Node) bool {
	return n.Greater(other) || n.Equals(other)
}

//NewEmptyNode creates an empty Node.
func NewEmptyNode() Node {
	return &Impl{}
}

//New creates a Node with load and priority.
func New(load interface{}, priority uint) Node {
	return Impl{Load: load, Priority: priority}
}

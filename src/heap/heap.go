package heap

import (
	"node"
	"errors"
	"sync"
)

type HeapManagerItf interface {
	Pop() (node.Node,  error)
	Push(node node.Node) error
	GetLength() uint32
}

type HeapManager struct {
	HeapManagerItf
	mut sync.Mutex
}

func(hm *HeapManager) Pop() node.Node {
	return hm.getNode(hm.GetLength())
}

func (hm *HeapManager) Push(node node.Node) error{
	// la concurrence est à traiter plus tard.
	hm.mut.Lock()
	defer hm.mut.Unlock()
	return hm.setNode(node, GetLength())
	
}

func (hm *HeapManager) setNode(node node.Node, pos uint32) error{
	

}

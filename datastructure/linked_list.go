package datastructure

type Node struct {
	value interface{}
	next *Node
}


type LinkedList interface {
	AddFirst(value interface{})
	AddLast(value interface{})
	Add(value interface{}, position int)
	Get(position int)
	Remove(position int)
	Size() int
	Has(value interface{}) bool
	getLastPtr() *Node
	increment()
}

type LinkedListImpl struct {
	first *Node
	size int
}

func New() *LinkedListImpl {
	return &LinkedListImpl{nil, 0 }
}

func (ll *LinkedListImpl) increment(){
	ll.size++
}

func (ll *LinkedListImpl) AddFirst(value interface{}) {
	node := Node{value, ll.first}
	ll.first = &node
	ll.increment()
}

func (ll *LinkedListImpl) AddLast(value interface{}){
	node := &Node{value, nil}
	ll.increment()

	if ll.first == nil {
		ll.first = node
		return
	}

	ll.getLastPtr().next = node
}

func (ll LinkedListImpl) Get(position int) {
	
}

func (ll LinkedListImpl) getLastPtr() *Node {
	ptr := ll.first
	if ptr == nil { return ptr }
	for ptr.next != nil { ptr = ptr.next }
	return ptr
}











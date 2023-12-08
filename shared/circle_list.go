package shared

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type CircleList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func NewCircleList[T any]() *CircleList[T] {
	return &CircleList[T]{head: nil, tail: nil}
}

func (cll *CircleList[T]) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircleList[T]) AddNode(data T) {
	newNode := &Node[T]{Data: data, Next: nil}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.Next = cll.head
	} else {
		cll.tail.Next = newNode
		cll.tail = newNode
		newNode.Next = cll.head
	}
}
func (cll *CircleList[T]) Pop() T {
	r := cll.head
	cll.head = cll.head.Next
	cll.tail = cll.head.Next
	return r.Data
}

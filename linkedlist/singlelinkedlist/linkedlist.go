package singlelinkedlist

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T] // 시작 Node
	tail *Node[T] // 끝 Node
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}

	// root 가 nil 인 경우는 tail 도 nil임을 보장
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	node := l.root

	cnt := 0
	for ; node != nil; node = node.next {
		cnt++
	}

	return cnt
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	// idx out
	if idx >= l.Count() {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	// work same
	// originNext := node.next
	// node.next = newNode
	// newNode.next = originNext
	node.next, newNode.next = newNode, node.next

	if node == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	targetNode := l.root
	for ; targetNode != nil; targetNode = targetNode.next {
		if targetNode == node {
			return true
		}
	}

	return false
}

// insert before 같은 경우는 기준노드의 전노드를 찾아야 하기에 고려해줘야 하는 부분들이 많다.
// 1. 이전노드가 root인 경우
// 2. 이전노드 찾기
func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	// 1. 이전노드가 root인 경우
	// 맨처음에 넣는거기에 pushFront와 동일하다
	if node == l.root {
		l.PushFront(value)
		return
	}
	// 2. 이전노드 찾기
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	prevNode.next, newNode.next = newNode, node
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	targetNode := l.root

	for ; targetNode != nil; targetNode = targetNode.next {
		if targetNode.next == node {
			return targetNode
		}
	}

	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}

	// 두번째가 첫번째가 되면 된다.
	l.root.next, l.root = nil, l.root.next

	if l.root == nil {
		l.tail = nil
	}
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}

	prev.next = node.next
	node.next = nil

	if node == l.tail {
		l.tail = prev
	}
}

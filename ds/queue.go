package ds

// LinkedQueue implements a queue using a LinkedList
type LinkedQueue struct {
	top    *LinkedListNode
	bottom *LinkedListNode
}

// Empty returns true if the queue is considered as empty
func (l *LinkedQueue) Empty() bool {
	return l.bottom == nil
}

// Enqueue enquueue a nodes into the list
func (l *LinkedQueue) Enqueue(nodes ...*LinkedListNode) {
	for _, node := range nodes {
		if l.top == nil {
			l.top, l.bottom = node, node
			node.next = nil
			continue
		}

		l.top.next = node
		l.top = node
	}

}

// Dequeue remove the first in node
func (l *LinkedQueue) Dequeue() (*LinkedListNode, bool) {
	if l.bottom == nil {
		return nil, false
	}

	node := l.bottom

	if node.next == nil {
		l.top, l.bottom = nil, nil
	} else {
		l.bottom = node.next
	}

	return node, true
}

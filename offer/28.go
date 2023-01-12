package offer

//面试题28：展平多级双向链表

type Node *struct {
	val   int
	prev  Node
	next  Node
	child Node
}

func NewNode(value int) Node {
	return &struct {
		val   int
		prev  Node
		next  Node
		child Node
	}{value, nil, nil, nil}
}

func Flatten(head Node) Node {
	flattenRecur(head)
	return head
}

func flattenRecur(head Node) Node {
	var prev Node
	node := head
	for node != nil {
		next := node.next
		if node.child != nil {
			tail := flattenRecur(node.child)
			node.next = node.child
			node.child.prev = node
			node.child = nil

			tail.next = next
			if next != nil {
				next.prev = tail
			}
			prev = tail
		} else {
			prev = node
		}
		node = next
	}
	return prev
}

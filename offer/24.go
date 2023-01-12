package offer

//面试题24：反转链表
//题目：定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。例如，把图4.8（a）中的链表反转之后得到的链表如图4.8（b）所示。

func RevertList(head ListNode) ListNode {
	var (
		prev ListNode
		cur  = head
	)
	for cur != nil {
		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	return prev
}

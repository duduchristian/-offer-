package offer

//试题25：链表中的数字相加
//题目：给定两个表示非负整数的单向链表，请问如何实现这两个整数的相加并且把它们的和仍然用单向链表表示？
//链表中的每个节点表示整数十进制的一位，并且头节点对应整数的最高位数而尾节点对应整数的个位数。
//例如，在图4.10（a）和图4.10（b）中，两个链表分别表示整数123和531，它们的和为654，对应的链表如图4.10（c）所示。

func AddTwoNumbers(head1, head2 ListNode) ListNode {
	head1 = RevertList(head1)
	head2 = RevertList(head2)
	return RevertList(addReversed(head1, head2))
}

func addReversed(head1, head2 ListNode) ListNode {
	dummy := NewListNode(0)
	sumNode := dummy
	carry := 0
	for head1 != nil || head2 != nil {
		sum := carry
		if head1 != nil {
			sum += head1.val
		}
		if head2 != nil {
			sum += head2.val
		}
		carry = sum / 10
		sum %= 10
		newNode := NewListNode(sum)

		sumNode.next = newNode
		sumNode = sumNode.next

		if head1 != nil {
			head1 = head1.next
		}
		if head2 != nil {
			head2 = head2.next
		}
	}
	if carry > 0 {
		sumNode.next = NewListNode(carry)
	}
	return dummy.next
}

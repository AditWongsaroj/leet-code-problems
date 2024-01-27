package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	carry := 0

	for l1 != nil || l2 != nil || carry == 1 {
		l1v := 0
		l2v := 0

		if l1 != nil {
			l1v = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
		}

		sum := l1v + l2v + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum %= 10
		}

		cur.Val = sum

		if l1 != nil || l2 != nil || carry == 1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return result
}

package removeNthNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lnArr := LnToArray(head)
	nPos := len(lnArr) - n
	
	lnArr = append(lnArr[:nPos], lnArr[nPos+1:]... )
	return Ln_builder(lnArr)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Ln_builder(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	ln := &ListNode{}
	cur := ln

	for i, num := range array {
		cur.Val = num
		if i != len(array)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return ln
}

func LnToArray(ln *ListNode) []int {
	arr := []int{}
	for ln != nil {
		arr = append(arr, ln.Val)
		ln = ln.Next
	}
	return arr
}
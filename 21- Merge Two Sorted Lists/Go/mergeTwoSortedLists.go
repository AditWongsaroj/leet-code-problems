package mergeTwoSortedLists

func oneNilList(cur, list *ListNode) *ListNode {
	for list != nil {
		cur.Val = list.Val
		list = list.Next
		if list != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return cur
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result

	for {
		if list1 == nil && list2 == nil {
			var temp *ListNode
			return temp
		}
		if list1 == nil {
			oneNilList(cur, list2)
			break
		}
		if list2 == nil {
			oneNilList(cur, list1)
			break
		}

		x := list1.Val
		y := list2.Val

		if x <= y {
			cur.Val = x
			list1 = list1.Next
		} else {
			cur.Val = y
			list2 = list2.Next
		}

		cur.Next = &ListNode{}
		cur = cur.Next
	}
	return result
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
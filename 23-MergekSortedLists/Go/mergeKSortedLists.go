package mergeKSortedLists

import (
	"sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
	r := []int{}

	for _, listnode := range lists {
		r = append(r, LnToArray(listnode)...)
	}
	sort.Ints(r)
    return Ln_builder(r)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Ln_builder(array []int) *ListNode {
    if len(array) == 0 {return nil}
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
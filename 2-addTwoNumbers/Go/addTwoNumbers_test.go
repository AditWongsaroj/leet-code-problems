package addTwoNumbers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func ln_builder(array []int) *ListNode {
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

func lnToArray(ln *ListNode) []int {
	arr := []int{}
	for ln != nil {
		arr = append(arr, ln.Val)
		ln = ln.Next
	}
	return arr
}

func TestArray(t *testing.T) {
	for _, c := range []struct {
		l1, l2, want []int
	}{
		{[]int{0}, []int{0}, []int{0}},
		{[]int{0}, []int{2, 4, 3}, []int{2, 4, 3}},
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{5, 6, 4}, []int{9, 9, 9}, []int{4, 6, 4, 1}},
		{[]int{9, 9, 9}, []int{9, 9, 9, 9, 9}, []int{8, 9, 9, 0, 0, 1}},
	} {
		got := AddTwoNumbers(ln_builder(c.l1), ln_builder(c.l2))
		if !cmp.Equal(got, ln_builder(c.want)) {
			t.Errorf("TwoSum(%v , %v) got == %v, want %v", c.l1, c.l2, lnToArray(got), c.want)
		}
	}
}

package mergeKSortedLists

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		array  [][]int
		want []int
	}{
		{[][]int{{1,4,5},{1,3,4},{2,6}}, []int{1,1,2,3,4,4,5,6}},
		{[][]int{}, []int{}},
	} {
		lists := []*ListNode{}
		for _, arr := range c.array {
			lists = append(lists, Ln_builder(arr))
			
		} 
			got := LnToArray(mergeKLists(lists))
			if !cmp.Equal(got, c.want) {
				t.Errorf("mergeKLists() got == %v, want %v", got, c.want)
			

		}
	}
}


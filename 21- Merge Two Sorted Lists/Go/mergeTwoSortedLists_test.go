package mergeTwoSortedLists

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		array  []int
		array2  []int
		want []int
	}{
	//	{[]int{1,2,4}, []int{1,3,4}, []int{1,1,2,3,4,4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	} {
			got := LnToArray(mergeTwoLists(Ln_builder(c.array), Ln_builder(c.array2)))
			if !cmp.Equal(got, c.want) {
				t.Errorf("mergeKLists() got == %v, want %v", got, c.want)
			

		}
	}
}


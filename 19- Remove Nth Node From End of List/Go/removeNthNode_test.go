package removeNthNode

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		array  []int
		n int
		want []int
	}{
		{[]int{1,2,3,4,5}, 2, []int{1,2,3,5}},
		{[]int{1}, 1, []int{}},
		{[]int{1,2}, 1, []int{1}},
	} {
			got := LnToArray(removeNthFromEnd(Ln_builder(c.array), c.n))
			if !cmp.Equal(got, c.want) {
				t.Errorf("mergeKLists() got == %v, want %v", got, c.want)
			

		}
	}
}


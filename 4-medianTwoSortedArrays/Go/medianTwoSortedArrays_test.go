package medianTwoSortedArrays

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T){
	for _, c := range[]struct{
		arr1, arr2 []int
		want float32
	}{
		{[]int{0}, []int{1,2}, 1 },
	}{
		got := medianTwoSortedArrays(c.arr1, c.arr2)
		if !cmp.Equal(got, float64(c.want)) {
			t.Errorf("medianTwoSortedArrays(%v, %v) got == %v, want %v", c.arr1, c.arr2, got, c.want)
		}
	}
}
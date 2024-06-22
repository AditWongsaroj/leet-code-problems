package threeSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1,0,1,2,-1,-4}, [][]int{{-1,-1,2},{-1,0,1}}},
		{[]int{0,1,1}, [][]int{}},
		{[]int{0,0,0}, [][]int{{0,0,0}}},
		
	} {
		got := threeSum(c.nums)
		if !cmp.Equal(got, c.want) {
			t.Errorf("got == %v, want %v", got, c.want)
		}
	}
}


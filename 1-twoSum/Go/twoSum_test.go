package twoSum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T){
	for _, c := range []struct{
		inTarget int
		inArray, want []int 
	}{
		{9, []int{2,7,12,4}, []int{0,1}},
		{6, []int{2,7,12,4}, []int{0,3}},
		{19, []int{2,7,12,4}, []int{1,2}},
		{0, []int{0,0}, []int{0,1}},
	}{
		got := TwoSum(c.inArray, c.inTarget)
		if !cmp.Equal( got, c.want){
			t.Errorf("TwoSum(%v , %v) == %v, want %v", c.inArray, c.inTarget, got, c.want)
		}
	}
}
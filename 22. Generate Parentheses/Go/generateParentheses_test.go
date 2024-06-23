package generateParentheses

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		num  int
		want []string
	}{
	//	{[]int{1,2,4}, []int{1,3,4}, []int{1,1,2,3,4,4}},
	//	{ 1, []string{"()"}},
		{ 3, []string{"((()))","(()())","(())()","()(())","()()()"}},
	} {
			got := generateParentheses(c.num)
			if !cmp.Equal(got, c.want) {
				t.Errorf("mergeKLists() got == %v, want %v", got, c.want)
			

		}
	}
}
package validParentheses

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		str  []string
		want bool
	}{
		{[]string{"()", "{}", "[]"}, true},
		{[]string{"]", ")", "}"}, false},
		{[]string{"(", "[", "{"}, false},
		{[]string{"()[]{}"}, true},
		{[]string{"([{}])", "[]({})", "[()]{}"}, true},
		{[]string{"(]", "{)"}, false},
		{[]string{"([)]"}, false},
	} {
		for _, s := range c.str {
			got := (isValid(s))
			if !cmp.Equal(got, c.want) {
				t.Errorf("IsValid( %v ) got == %v, want %v", s, got, c.want)
			}

		}
	}

}

package stringToInteger

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArray(t *testing.T) {
	for _, c := range []struct {
		str  []string
		want int
	}{
		{[]string{"42", " 42", " +42", "042a2"}, 42},
		{[]string{"w", "0", "0-1"}, 0},
		{[]string{"-42", " -042"}, -42},
		{[]string{"-2147483649"}, -2147483648},
		{[]string{"2147483649"}, 2147483647},
		{[]string{""}, 0},
	} {
		for _, s := range c.str {
			got := (myAtoi(s))
			if !cmp.Equal(got, c.want) {
				t.Errorf("IsValid( %v ) got == %v, want %v", s, got, c.want)
			}

		}
	}
}
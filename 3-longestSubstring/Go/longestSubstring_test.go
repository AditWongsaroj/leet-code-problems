package longestsubstring

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)


func TestArray(t *testing.T){
	for _, c := range []struct {
		str []string
		want int
	}{
		{[]string{""}, 0},
		{[]string{"a","aa", " ", "111", "?"}, 1}, 
		{[]string{"ab", " 2", "/1", "1/1"}, 2},
		{[]string{"abc", "abcabcbba", "dvds", "abccba"}, 3},
		{[]string{"abcd","abcdabcd", "abcabcd", "1234432111", "^%$?"}, 4}, 
	} {
		for _, s := range c.str{
			got := LengthOfLongestSubstring(s)
			if !cmp.Equal(got, c.want) {
				t.Errorf("LengthOfLongestSubstring(%v) got == %v, want %v", c.str, got, c.want)
			}

		}
	}
	
}
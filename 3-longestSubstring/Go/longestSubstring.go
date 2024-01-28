package longestsubstring

func LengthOfLongestSubstring(s string) int {
	charmap := []rune{}
	longest := 0
	charcount := 0
	for _, ch := range s {

		if longest < charcount {
			longest = charcount
		}
		if idx, ok := rInString(ch, charmap); ok {
			charmap = charmap[idx+1:]
			charcount = len(charmap)
		}
		charcount++
		charmap = append(charmap, ch)
	}
	if longest < charcount {
		longest = charcount
	}
	return longest
}

func rInString(r rune, str []rune) (int, bool) {
	for i, b := range str {
		if b == r {
			return i, true
		}
	}
	return 0, false
}

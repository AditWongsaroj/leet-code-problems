package generateParentheses

func generateParentheses(n int) []string {
	return genBrancher([]rune{'('}, n, 1, 0)
}

func genBrancher(list []rune, n, open, close int) []string {
	if n == open {
		for len(list) < n*2 {
			list = append(list, ')')
		}
		return []string{string(list)}
	}
	if open > close {
		return append(genBrancher(append(list, '('), n, open+1, close), genBrancher(append(list, ')'), n, open, close+1)...)
	}
	if close == open {
		return genBrancher(append(list, '('), n, open+1, close)
	}
	return []string{}
}
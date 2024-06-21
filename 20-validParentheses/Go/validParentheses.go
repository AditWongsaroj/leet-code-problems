package validParentheses

func isValid(s string) bool {
	seen := []rune{}

	for _, bracket := range s {

		switch bracket {
		case '[':
			seen = append(seen, '[')
		case '{':
			seen = append(seen, '{')
		case '(':
			seen = append(seen, '(')
		case ']':
			if len(seen) == 0 || seen[len(seen)-1] != '[' {
				return false
			}
			seen = seen[:len(seen)-1]
		case '}':
			if len(seen) == 0 || seen[len(seen)-1] != '{' {
				return false
			}
			seen = seen[:len(seen)-1]
		case ')':

			if len(seen) == 0 || seen[len(seen)-1] != '(' {
				return false
			}
			seen = seen[:len(seen)-1]
		}
	}
	return len(seen) == 0
}
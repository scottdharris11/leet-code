package validparentheses

func isValid(s string) bool {
	// odd length string is always invalid
	if len(s)%2 != 0 {
		return false
	}

	// for each starting character, add matching end to stack, when
	// ending character encountered, ensure it matches the latest
	// starting character
	var stack []rune
	stackLen := 0
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
			stackLen++
		case '[':
			stack = append(stack, ']')
			stackLen++
		case '{':
			stack = append(stack, '}')
			stackLen++
		default:
			n := stackLen - 1
			if n < 0 {
				return false
			}
			validClose := stack[n]
			if c != validClose {
				return false
			}
			stack = stack[:n]
			stackLen--
		}
	}
	return stackLen == 0
}

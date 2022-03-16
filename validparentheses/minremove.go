package validparentheses

func minRemoveToMakeValid(s string) string {
	o := make([]rune, len(s))
	oIdx := 0
	var stack []int
	stackLen := 0
	for _, c := range s {
		switch c {
		case ')':
			// skip character if not matching begin paren on stack
			n := stackLen - 1
			if n < 0 {
				continue
			}
			stack = stack[:n]
			stackLen--
		case '(':
			// capture index of begin paren on stack
			stack = append(stack, oIdx)
			stackLen++
		}
		o[oIdx] = c
		oIdx++
	}

	// splice out any beginning parens that were not closed
	// start at back so index values remain valid
	for stackLen > 0 {
		i := stack[stackLen-1]
		o = append(o[:i], o[i+1:]...)
		stackLen--
		oIdx--
	}
	return string(o[:oIdx])
}

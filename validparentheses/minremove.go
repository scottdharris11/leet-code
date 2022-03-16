package validparentheses

func minRemoveToMakeValid(s string) string {
	o := make([]rune, len(s))
	oIdx := 0
	var stack []int
	stackLen := 0
	for _, c := range s {
		switch c {
		case ')':
			n := stackLen - 1
			if n < 0 {
				continue
			}
			stack = stack[:n]
			stackLen--
			fallthrough
		default:
			if c == '(' {
				stack = append(stack, oIdx)
				stackLen++
			}
			o[oIdx] = c
			oIdx++
		}
	}

	for stackLen > 0 {
		i := stack[stackLen-1]
		o = append(o[:i], o[i+1:]...)
		stackLen--
		oIdx--
	}
	return string(o[:oIdx])
}

package validparentheses

func scoreOfParentheses(s string) int {
	ps, _ := score(s, 1, len(s))
	return ps
}

func score(s string, i int, l int) (int, int) {
	scoreVal := 0
	idx := i

	switch s[idx] {
	case ')':
		scoreVal = 1
	case '(':
		rs, ri := score(s, idx+1, l)
		scoreVal = 2 * rs
		idx = ri + 1
	}

	if idx+1 < l && s[idx+1] == '(' {
		rs, ri := score(s, idx+2, l)
		scoreVal += rs
		idx = ri
	}

	return scoreVal, idx
}

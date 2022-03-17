package validatestack

func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	popIdx := 0
	var stack []int
	stackLen := 0
	for pushIdx := 0; pushIdx < l; pushIdx++ {
		stack = append(stack, pushed[pushIdx])
		stackLen++
		for n := stackLen - 1; n >= 0; n-- {
			if stack[n] != popped[popIdx] {
				break
			}
			popIdx++
			stack = stack[:n]
			stackLen--
		}
	}
	return popIdx == l
}

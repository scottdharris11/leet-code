package simplifypath

func simplifyPath(path string) string {
	var stack []string
	bIdx := -1
	eIdx := -1
	stackLen := 0

	evalSection := func() {
		if bIdx > -1 {
			section := path[bIdx : eIdx+1]
			switch section {
			case ".":
				break
			case "..":
				if stackLen > 0 {
					stackLen--
					stack = stack[:stackLen]
				}
			default:
				stack = append(stack, section)
				stackLen++
			}
			bIdx = -1
		}
	}

	for i, c := range path {
		switch c {
		case '/':
			evalSection()
		default:
			if bIdx == -1 {
				bIdx = i
			}
			eIdx = i
		}
	}

	evalSection()

	if stackLen == 0 {
		return "/"
	}

	output := ""
	for _, s := range stack {
		output += "/" + s
	}
	return output
}

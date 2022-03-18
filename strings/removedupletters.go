package strings

func removeDuplicateLetters(s string) string {
	// Track the indexes of each character in the string
	var cIndexes [26][]int
	cCnt := 0
	mIdx := int32(0)
	for i, c := range s {
		cIdx := c - 'a'
		if cIndexes[cIdx] == nil {
			cIndexes[cIdx] = make([]int, 0, 10)
			cCnt++
			if cIdx > mIdx {
				mIdx = cIdx
			}
		}
		cIndexes[cIdx] = append(cIndexes[cIdx], i)
	}

	// Cycle through set of characters, applying the lowest
	// character with each pass that has remaining characters
	// after its index
	buf := make([]rune, cCnt)
	loc := -1
	for i := 0; i < cCnt; i++ {
		cIdx := int32(-1)
		for j := int32(0); j <= mIdx; j++ {
			// Find the lowest index for the character past the last applied location
			if cIndexes[j] == nil {
				continue
			}

			checkIdx := -1
			for k := 0; k < len(cIndexes[j]); k++ {
				if cIndexes[j][k] > loc {
					checkIdx = cIndexes[j][k]
					break
				}
			}

			if checkIdx < 0 {
				continue
			}

			// Determine if all remaining characters have entries after this character
			valid := true
			for k := int32(0); k <= mIdx; k++ {
				if j == k || cIndexes[k] == nil {
					continue
				}
				afterIdx := false
				for l := 0; l < len(cIndexes[k]); l++ {
					if cIndexes[k][l] > checkIdx {
						afterIdx = true
						break
					}
				}

				if !afterIdx {
					valid = false
					break
				}
			}

			// Found our character
			if valid {
				cIdx = j
				loc = checkIdx
				break
			}
		}

		buf[i] = 'a' + cIdx
		cIndexes[cIdx] = nil
	}

	return string(buf)
}

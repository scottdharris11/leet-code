package partitionlabels

func partitionLabels(s string) []int {
	l := len(s)
	pBegin := 0
	pEnd := 0
	var partitions []int
	for i, c := range s {
		last := lastIndexOf(s, uint8(c), l, pEnd)
		if last > pEnd {
			pEnd = last
		}

		if i == pEnd {
			partitions = append(partitions, i-pBegin+1)
			pBegin = i + 1
		}
	}
	return partitions
}

func lastIndexOf(s string, c uint8, l int, e int) int {
	for i := l - 1; i >= e; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

package medianofsortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	count := 0
	merged := make([]int, l1+l2)
	var v1, v2 int
	for i1, i2 := 0, 0; i1 < l1 || i2 < l2; {
		if i1 < l1 {
			v1 = nums1[i1]
		}
		if i2 < l2 {
			v2 = nums2[i2]
		}

		if i2 >= l2 || (i1 < l1 && v1 < v2) {
			merged[count] = v1
			count++
			i1++
			continue
		}

		merged[count] = v2
		count++
		i2++
	}

	mIdx := count / 2
	if count%2 == 0 {
		return float64(merged[mIdx-1]+merged[mIdx]) / 2.0
	}
	return float64(merged[mIdx])
}

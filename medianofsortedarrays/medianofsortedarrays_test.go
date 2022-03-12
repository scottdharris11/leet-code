package medianofsortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected float64
	}{
		{"1", []int{2, 3, 4}, []int{5, 6, 7}, 4.5},
		{"2", []int{2, 3, 4}, []int{5, 6, 7, 8}, 5.0},
		{"3", []int{1, 3}, []int{2}, 2.0},
		{"4", []int{1, 2}, []int{3, 4}, 2.5},
		{"5", []int{}, []int{3, 4}, 3.5},
		{"6", []int{1, 2}, []int{}, 1.5},
		{"7", []int{1, 2}, nil, 1.5},
		{"8", nil, []int{3, 4}, 3.5},
		{"9", []int{-2, 3, 4}, []int{-5, 6, 7}, 3.5},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := findMedianSortedArrays(tt.l1, tt.l2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

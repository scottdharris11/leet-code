package validatestack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		name     string
		pushed   []int
		popped   []int
		expected bool
	}{
		{"1", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{"2", []int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
		{"3", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 5, 3, 6, 7, 4}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, validateStackSequences(tt.pushed, tt.popped))
		})
	}
}

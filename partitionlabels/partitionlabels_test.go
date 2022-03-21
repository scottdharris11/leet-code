package partitionlabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		partitions []int
	}{
		{"1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"2", "eccbbbbdec", []int{10}},
		{"3", "caedbdedda", []int{1, 9}},
		{"4", "caedbdeddaf", []int{1, 9, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := partitionLabels(tt.input)
			assert.ElementsMatch(t, tt.partitions, result)
		})
	}
}

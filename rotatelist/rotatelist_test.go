package rotatelist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		list     *ListNode
		places   int
		expected []int
	}{
		{"1", buildListNode([]int{1, 2, 3, 4, 5}), 1, []int{5, 1, 2, 3, 4}},
		{"2", buildListNode([]int{1, 2, 3, 4, 5}), 2, []int{4, 5, 1, 2, 3}},
		{"3", buildListNode([]int{1, 2, 3, 4, 5}), 3, []int{3, 4, 5, 1, 2}},
		{"4", buildListNode([]int{1, 2, 3, 4, 5}), 4, []int{2, 3, 4, 5, 1}},
		{"5", buildListNode([]int{1, 2, 3, 4, 5}), 5, []int{1, 2, 3, 4, 5}},
		{"6", buildListNode([]int{1, 2, 3, 4, 5}), 6, []int{5, 1, 2, 3, 4}},
		{"7", buildListNode([]int{1, 2, 3, 4, 5}), 7, []int{4, 5, 1, 2, 3}},
		{"8", buildListNode([]int{1, 2, 3, 4, 5}), 8, []int{3, 4, 5, 1, 2}},
		{"9", buildListNode([]int{1, 2, 3, 4, 5}), 9, []int{2, 3, 4, 5, 1}},
		{"10", buildListNode([]int{1, 2, 3, 4, 5}), 10, []int{1, 2, 3, 4, 5}},
		{"11", buildListNode([]int{1}), 0, []int{1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := rotateRight(tt.list, tt.places)
			for i, v := range tt.expected {
				if result == nil {
					assert.Fail(t, fmt.Sprintf("Expected value missing %d at index %d", v, i))
					continue
				}
				assert.Equal(t, v, result.Val, fmt.Sprintf("Expected value %d at index %d", v, i))
				result = result.Next
			}
			assert.Nil(t, result)
		})
	}
}

func buildListNode(values []int) *ListNode {
	b := &ListNode{Val: values[0]}
	p := b
	for i := 1; i < len(values); i++ {
		n := &ListNode{Val: values[i]}
		p.Next = n
		p = n
	}
	return b
}

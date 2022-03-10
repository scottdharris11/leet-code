package addtwonumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected []int
	}{
		{"1", buildListNode([]int{2, 4, 3}), buildListNode([]int{5, 6, 4}), []int{7, 0, 8}},
		{"2", buildListNode([]int{0}), buildListNode([]int{0}), []int{0}},
		{"3", buildListNode([]int{9, 9, 9, 9, 9, 9, 9}), buildListNode([]int{9, 9, 9, 9}), []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := addTwoNumbers(tt.l1, tt.l2)
			for i, v := range tt.expected {
				if result == nil {
					assert.Fail(t, fmt.Sprintf("Expected value missing %d at index %d", v, i))
					continue
				}
				assert.Equal(t, v, result.Val)
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

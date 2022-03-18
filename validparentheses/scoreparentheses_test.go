package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreOfParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"1", "()", 1},
		{"2", "()()()", 3},
		{"3", "(())", 2},
		{"4", "((()))", 4},
		{"5", "((()()))", 8},
		{"6", "((()()))((()()()))", 20},
		{"7", "(((()())()))((()()()))", 32},
		{"8", "(())()", 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, scoreOfParentheses(tt.input))
		})
	}
}

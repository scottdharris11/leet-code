package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", "lee(t(c)o)de)", "lee(t(c)o)de"},
		{"2", "a)b(c)d", "ab(c)d"},
		{"3", "))((", ""},
		{"4", "((abc)", "(abc)"},
		{"5", "(((abc))", "((abc))"},
		{"6", "())()(((", "()()"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, minRemoveToMakeValid(tt.input))
		})
	}
}

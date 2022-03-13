package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"1", "()", true},
		{"2", "()[]{}", true},
		{"3", "())", false},
		{"4", "([{}])", true},
		{"5", "{", false},
		{"6", "(]", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isValid(tt.input))
		})
	}
}

package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStackSequences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", "bcabc", "abc"},
		{"2", "cbacdcbc", "acdb"},
		{"3", "mitnlruhznjfyzmtmfnstsxwktxlboxutbic", "ilrhjfyzmnstwkboxuc"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, removeDuplicateLetters(tt.input))
		})
	}
}

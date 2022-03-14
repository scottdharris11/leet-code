package simplifypath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1", "/home/", "/home"},
		{"2", "/../", "/"},
		{"3", "/home//foo/", "/home/foo"},
		{"4", "/home/../var/./log/.../../.././temp/", "/var/temp"},
		{"5", "/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, simplifyPath(tt.input))
		})
	}
}

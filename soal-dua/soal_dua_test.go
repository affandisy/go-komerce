package main

import (
	"testing"

	"github.com/go-openapi/testify/assert"
)

func TestSortCharacter(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		anggota     []int
		expectedBus int
	}{
		{"First Case", 5, []int{1, 2, 4, 3, 3}, 4},
		{"Second Case", 8, []int{2, 3, 4, 4, 2, 1, 3, 1}, 5},
		{"Third Case", 5, []int{1, 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := HitungMinibus(tt.n, tt.anggota)
			assert.Equal(t, tt.expectedBus, v)
		})
	}
}

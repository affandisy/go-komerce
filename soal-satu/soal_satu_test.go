package main

import (
	"testing"

	"github.com/go-openapi/testify/assert"
)

func TestSortCharacter(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedVowel     string
		expectedConsonant string
	}{
		{"First Case", "Sample Case", "aaee", "ssmplc"},
		{"Second Case", "Next Case", "eea", "nxtcs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, k := SortCharacters(tt.input)
			assert.Equal(t, tt.expectedVowel, v)
			assert.Equal(t, tt.expectedConsonant, k)
		})
	}
}

func TestSortCharacterBetter(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedVowel     string
		expectedConsonant string
	}{
		{"First Case", "Sample Case", "aaee", "ssmplc"},
		{"Second Case", "Next Case", "eea", "nxtcs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, k := SortCharactersBetter(tt.input)
			assert.Equal(t, tt.expectedVowel, v)
			assert.Equal(t, tt.expectedConsonant, k)
		})
	}
}

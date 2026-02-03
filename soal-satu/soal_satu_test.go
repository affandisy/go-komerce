package main

import (
	"fmt"
	"testing"
)

func TestSortCharacter(t *testing.T) {
	input1 := "Sample Case"
	vowels, consonants := SortCharacters(input1)

	fmt.Println("Vowel Characters: ", vowels)
	fmt.Println("Consonant Characters: ", consonants)

	input2 := "Next Case"
	vowels2, consonants2 := SortCharacters(input2)
	fmt.Println("Vowel Characters: ", vowels2)
	fmt.Println("Consonant Characters: ", consonants2)
}

package main

import (
	"fmt"
	"strings"
)

func SortCharacters(input string) (string, string) {
	vowels := ""
	consonants := ""
	// panjangInput := len(input)

	inputLowCase := strings.ToLower(input)

	for _, char := range inputLowCase {
		sChar := string(char)

		if sChar == " " || sChar == "\n" {
			continue
		}

		if sChar == "a" || sChar == "e" || sChar == "i" || sChar == "o" || sChar == "u" {
			vowels = vowels + sChar
		} else {
			consonants = consonants + sChar
		}
	}

	return vowels, consonants
}

func main() {
	input1 := "Sample Case"
	vowels, consonants := SortCharacters(input1)

	fmt.Println("Vowel Characters: ", vowels)
	fmt.Println("Consonant Characters: ", consonants)

}

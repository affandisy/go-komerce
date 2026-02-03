package main

import (
	"fmt"
	"strings"
)

func SortCharacters(input string) (string, string) {
	vowels := ""
	consonants := ""
	panjangInput := len(input)

	inputLowCase := strings.ToLower(input)

	for i := 0; i < panjangInput; i++ {
		char := inputLowCase[i]

		if char == ' ' || char == '\n' {
			continue
		}

		if char >= 'a' && char <= 'z' {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				vowels = vowels + string(char)
			} else {
				consonants = consonants + string(char)
			}
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

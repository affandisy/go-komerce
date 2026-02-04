package main

import "strings"

func SortCharactersBetter(input string) (string, string) {
	counts := make(map[rune]int)

	var vowelOrder []rune
	var consonantOrder []rune

	inputLowCase := strings.ToLower(input)

	for _, char := range inputLowCase {
		if char == ' ' || char == '\n' {
			continue
		}

		isVowel := false
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			isVowel = true
		}

		if counts[char] == 0 {
			if isVowel {
				vowelOrder = append(vowelOrder, char)
			} else {
				consonantOrder = append(consonantOrder, char)
			}
		}

		counts[char]++
	}

	vowels := ""
	for _, char := range vowelOrder {
		count := counts[char]
		for i := 0; i < count; i++ {
			vowels = vowels + string(char)
		}
	}

	consonants := ""
	for _, char := range consonantOrder {
		count := counts[char]
		for i := 0; i < count; i++ {
			consonants = consonants + string(char)
		}
	}

	return vowels, consonants
}

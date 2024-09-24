package main

import (
	"fmt"
)

func Ft_max_substring(s string) int {
	charIndex := make(map[byte]int)

	maxLength := 0
	start := 0

	for i := 0; i <= len(s); i++ {
		char := s[i]

		if lastIndex, found := charIndex[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		charIndex[char] = i
	}

	return maxLength
}

func main() {
	fmt.Println(Ft_max_substring("abcabcbb")) // Résultat attendu : 3
	fmt.Println(Ft_max_substring("bbbbb"))    // Résultat attendu : 1
}

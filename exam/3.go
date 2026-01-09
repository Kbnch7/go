package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}

	first := word[0]
	rest := word[1:]
	reversed := ""

	for i := len(rest) - 1; i >= 0; i-- {
		reversed = reversed + string(rest[i])
	}

	return string(first) + reversed
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)

	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += encryptWord(word)
	}
	return result
}

func main() {
	phrases := []string{
		"Pepe Schnele is a legend",
		"Plaki plaki",
		"Tun tun tun sahur",
	}

	for _, phrase := range phrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("Было:  %s\nСтало: %s\n\n", phrase, encrypted)
	}
}
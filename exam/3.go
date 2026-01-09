package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
    if len(word) <= 1 {
        return word
    }

    runes := []rune(word)
    first := runes[0]

    reversed := ""
    for i := len(runes) - 1; i >= 1; i-- {
        reversed += string(runes[i])
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
		"Pepe Schnele is a legend але",
		"Plaki plaki",
		"Tun tun tun sahur",
	}

	for _, phrase := range phrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("Было:  %s\nСтало: %s\n\n", phrase, encrypted)
	}
}
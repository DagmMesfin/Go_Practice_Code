package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(wordCount("Hello there Hello there, how are you, and you"))
	fmt.Println(palindromeChecker("akabak"))
}

func wordCount(sent string) map[string]int {
	sent = strings.ToLower(sent)

	var clean strings.Builder

	for _, ch := range sent {
		if unicode.IsLetter(ch) || unicode.IsSpace(ch) || unicode.IsDigit(ch) {
			clean.WriteRune(ch)
		}
	}

	word_arr := strings.Fields(clean.String())

	dicto := make(map[string]int)

	for _, word := range word_arr {
		dicto[word] += 1
	}

	return dicto
}

func palindromeChecker(word string) bool {
	n := len(word)
	for i := range n / 2 {
		if word[i] != word[n-i-1] {
			return false
		}
	}
	return true
}

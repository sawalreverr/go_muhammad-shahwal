package task

import (
	"slices"
	"strings"
)

func GroupPalindrome(words []string) []any {
	var palindrome []string
	var nonPalindrome []string
	var returnGroup []any

	for _, word := range words {
		revSliceWord := strings.Split(word, "")
		slices.Reverse(revSliceWord)

		if slices.Equal(strings.Split(word, ""), revSliceWord) {
			palindrome = append(palindrome, strings.Join(revSliceWord, ""))
		} else {
			nonPalindrome = append(nonPalindrome, word)
		}
	}

	if len(palindrome) != 0 {
		returnGroup = append(returnGroup, palindrome)
	}

	for _, word := range nonPalindrome {
		returnGroup = append(returnGroup, word)
	}

	return returnGroup
}

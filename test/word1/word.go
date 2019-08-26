// Package word provide some word functions.
package word

import (
	"unicode"
)

// IsPalindrome will check if a string is a palindrome string.
// (The first attempt)
func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}

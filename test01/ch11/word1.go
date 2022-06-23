package word

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	fmt.Println(letters)
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

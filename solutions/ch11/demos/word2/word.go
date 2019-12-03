package word2

import "unicode"

type Values struct {
	a int
}

func IsPalindrome(s string) bool {
	var letters []rune
	for _, c := range s {
		if unicode.IsLetter(c) {
			letters = append(letters, unicode.ToLower(c))
		}
	}

	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}

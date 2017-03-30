package anagram

import (
	"strings"
)

func isAnagram(first, second string) bool {
	first = strings.ToLower(first)
	second = strings.ToLower(second)
	return first != "" && second != "" && doesContainSameChars(first, second) && doesContainSameChars(second, first)
}

func doesContainSameChars(first, second string) bool {
	for _, r := range first {
		c := string(r)
		if c != " " && strings.Count(first, c) != strings.Count(second, c) {
			return false
		}
	}

	return true
}

package anagram

import "testing"

func TestAnagram(t *testing.T) {
	cases := []struct {
		first, second string
		isAnagram     bool
	}{
		{"Dave Barry", "Ray Adverb", true},
		{"Tom Marvolo Riddle", "I am Lord Voldemort", true},
		{"просветитель", "терпеливость", true},
		{"равновесие", "своенравие", true},

		{"test", "fail", false},
		{"ave Barry", "Ray Adverb", false},
		{"", "", false},
	}

	for _, c := range cases {
		got := isAnagram(c.first, c.second)
		if got != c.isAnagram {
			t.Errorf("isAnagram(%v, %v) == %v, want %v", c.first, c.second, got, c.isAnagram)
		}
	}
}

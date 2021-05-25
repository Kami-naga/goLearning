package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbbb", 1},
		{"abcabcabcd", 4},
		// Chinese support
		{"今天天气不错，我们去散步吧", 11},
		{"锄禾日当午", 5},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

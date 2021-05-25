package main

import "testing"

//use go test -coverprofile=c.out
//to get a coverage profile file &
//use go tool cover -html=c.out to see the result(use cmd,then you can get a html)
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

//we can also do below tests by cmd
// go test -bench .
//we can also use
//go test -bench . -cpuprofile cpu.out
//to get a performance file(binary), then use
//go tool pprof cpu.out            to see the result
//in pprof, type web to get the graph(graphviz needed)
func BenchmarkSubstr(b *testing.B) {
	s, ans := "今天天气不错，我们去散步吧", 12
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))

	b.ResetTimer() //just count time of below codes

	// how many times we need to run below codes?
	// b.N times! benchmark decides that
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d", actual, s, ans)
		}
	}
}

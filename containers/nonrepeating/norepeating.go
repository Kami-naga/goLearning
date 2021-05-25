package main

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("PPP今日は晴れだ!"))
}

//put slice making outside the function to avoid it making slice many times,
//then we can get better performance in benchmark
//but I just found the result is the same, so it maybe optimized
var lastOccurred = make([]int, 0xffff) //65535 length, 'a' = 97d = 0x61, so lastOccurred['a'] = x is just like a map

//form the result of pprof, we can get the most time-consuming part of the code is
//1. decoding string to rune
//2. map operation(map access & map assign)
//so how to optimize? 1 is inevitable because we need rune to decode and understand the contents,
//but we can optimize map(just use another data structure)
//each time we add an element into a map,
//map needs to calculate hash, check if duplicate, assign space if not enough...
//so we can use more space to get better time, just make a large slice instead of a map!
//lastOccurred := make(map[rune]int)
func lengthOfNonRepeatingSubStr(s string) int {
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0
	//rune(s) -> decoding& make a new rune array, not re-understanding the same memory,
	for i, ch := range []rune(s) { //only range s will not run well because each jp/ch charac will add 3(length)
		//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

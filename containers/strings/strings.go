package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "PPP今日は晴れだ!"             // UTF=8
	fmt.Println(len(s))           // 22 why? 3 + 3x6 + 1 获得字节长度
	fmt.Printf("%s\n", []byte(s)) // 用[]byte获得字节
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is rune
		fmt.Printf("(%d %X)", i, ch) // utf-8解码，解完每个字符转成unicode放进rune
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) //utf8.RuneCountInString(s)获得字符数量
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) { //此处转rune并非是对同一块内存内容的重新理解，而是对其解码后新开一个rune数组，不是同一块内存
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	//string operations, find it in strings.~~~
	fmt.Println(strings.Fields("asdia asidhoa ashdo")) //split on the basis of white spaces and store in a string array
	fmt.Println(strings.Split("sadsa ddd ss, ssss", ","))
	fmt.Println(strings.Join(strings.Split("sadsa ddd ss, ssss", ","), "$$$"))
	fmt.Println(strings.Contains("sadsa ddd ss, ssss", "s"))
	fmt.Println(strings.Index("sadsa ddd ss, ssss", "s"))
	fmt.Println(strings.ToLower("SURGEnnn"))
	fmt.Println(strings.ToUpper("SURGEnnn"))
	fmt.Println(strings.Trim("nnnSURnnnGEnnn", "n"))
	fmt.Println(strings.TrimLeft("nnnSURnnnGEnnn", "n"))
	fmt.Println(strings.TrimRight("nnnSURnnnGEnnn", "n"))

}

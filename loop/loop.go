package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "b.txt"

func main() {
	//in println, if last brace is in a new line, don't forget to add a comma
	//fmt.Println(
	//	convertToBin(5),
	//	convertToBin(13),
	//	convertToBin(0))
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	printFile(filename)
	easyCheck()
	//forever()
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}
	return result
}

// read file 1 line per time
func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		//below for is what called while in other languages
		//No "while" in GO
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

//in GO, it's easy to write Infinite loop
func forever() {
	for {
		fmt.Println("forever looping")
	}
}

func easyCheck() {
	s := bufio.NewScanner(os.Stdin)
	for times := 1; times <= 3; times++ {
		fmt.Println("Please input your information of account:")
		s.Scan()
		ss := strings.Split(s.Text(), " ")
		if ss[0] == "a" && ss[1] == "666" {
			fmt.Println("OK")
			break
		} else {
			fmt.Printf("You have %d more chances!\n", 3-times)
		}
	}
}

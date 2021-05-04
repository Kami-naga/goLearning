package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "a.txt"

func main() {
	readFile1()
	readFileSimple()
	fmt.Println(trySwitch(3, 4, "+"))
	fmt.Println(
		grade(0),
		grade(44),
		grade(66),
		grade(88),
		grade(99),
		grade(100),
	//grade(-1)
	)
}

func readFile1() {
	contents, err := ioutil.ReadFile(filename)
	// no braces after if in GO
	if err != nil {
		fmt.Println(err)
	} else {
		//println will get [97 98 99 13 10 49 50 51 13 10 121 121 121 13 10 116 116 116 116] ascii code
		//fmt.Println(contents)
		fmt.Printf("%s\n", contents)
	}
}

func readFileSimple() {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		//println will get [97 98 99 13 10 49 50 51 13 10 121 121 121 13 10 116 116 116 116] ascii code
		//fmt.Println(contents)
		fmt.Printf("%s\n", contents)
	}
	//the scope of contents is only if, outside if, we can't access it
	//fmt.Println(contents)
}

func trySwitch(a, b int, op string) int {
	// break will be automatically added in each case
	// so only when you don't need break ,you should add fallthrough
	// panic is something like cerr
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	// no expr after switch is also OK
	//just add it in case
	switch {
	case score < 0 || score > 100:
		// we can get a string by sprintf function
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

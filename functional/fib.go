package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//GO特点之一，能给函数实现接口
type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	//TODO： incorrect if p(buffer) is too small
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//below for is what called while in other languages
	//No "while" in GO
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//Go语言闭包运用
//例一： 斐波那契数列
//例二：为函数实现接口
//例三： 使用函数来遍历二叉树
//GO闭包的特点
//1.更为自然，不需要修饰如何访问自由变量，对函数内自由变量可变性管得较松
//2.没有Lambda表达式，但是有匿名函数，加强匿名函数用法后，Lambda纯粹只是语法上的添加，没必要再加Lambda表达式了，所以GO没Lambda
//3.
func main() {
	f := fibonacci()

	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 2
	//fmt.Println(f()) // 3
	//fmt.Println(f()) // 5
	//fmt.Println(f()) // 8
	//fmt.Println(f()) // 13
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	printFileContents(f)
}

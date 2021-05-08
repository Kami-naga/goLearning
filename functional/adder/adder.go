package main

import "fmt"

//函数式编程
//函数体中含有局部变量（下边的v）以及自由变量，函数环境的变量，即下边的sum
//函数体（局部变量&自由变量） + （自由变量若为一个结构，则还会包括该结构关联的所有变量，最终为一棵树，此处为int的sum，所以就一个值）= 闭包
//函数返回时返回的是一个闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//func orthodox approach
type iAdder func(int) (int, iAdder) //recursion

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

//python支持原生闭包，自由变量需加nonlocal， 五匿名函数，所以需要起名字
//c++也支持，以前用stl或boost库，c++11后支持闭包，对于自由变量sum的修饰有两点，[=]表示传值进来，mutable表示可变
//java也支持， 闭包引用的sum不可以改变其值，所以建一个sum对象，里头一个value，可以改sum.value，然后Lambda表达式返回，或匿名类
func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}

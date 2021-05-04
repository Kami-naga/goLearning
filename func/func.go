package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {

	fmt.Println(div(13, 3))

	//fmt.Println(eval1(3, 4, "*"))

	if result, err := eval2(3, 4, "n"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	//try functional programming
	// function can also be its params
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}

//没有默认函数，可选参数，函数重载，操作符重载，只有一个可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func eval1(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		//result = a / b
		//how to use only one of the return value?
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

//in eval1, if you go to panic, it will stop the program which is not so good
//so we use an error to deal with panic
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//result = a / b
		//how to use only one of the return value?
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsupported operator:" + op)
		return 0, fmt.Errorf("unsupported operator:" + op)
	}
}

// try functional programming !
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//easy division
// way 1
//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

// it is good to give a name for the return value
// then you can use ctrl + alt + v to div(1,2)
// then you will get q, r := div(1,2) quickly

//way 2
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//way 3 ,but it makes return value not clear, so it's not good
// use way 2 will better
//func div(a, b int) (q, r int) {
//	q = a / b
//	r = a % b
//	return
//}

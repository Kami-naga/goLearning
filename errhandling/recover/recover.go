package main

import (
	"fmt"
)

//use error ! use less panic as you can
//only those unexpected error you can use panic
//(you don't know what will happen &
//you don't know how to deal with them)
func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover() //recover returns an interface! (any type) &recover can only be used in defer
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do with: %v", r))
		}
	}() // we need the () here,
	// which means we called a function we created(definition+ function call)
	// just like defer func()

	//panic(errors.New("this is an error"))

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123)
}

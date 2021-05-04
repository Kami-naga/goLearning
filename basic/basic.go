package main

import "fmt"

// var can't be omitted here by :=
// the scope of aa & ss & ff here is package, not global!
var aa = 666
var ss = "rrr"
var ff = true

//above code can be simplified as below
var (
	bb = 777
	cc = "yyy"
	dd = false
)

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTpeDeduction()
	variableShorter()
	fmt.Println(aa, ss, ff)
	fmt.Println(bb, cc, dd)
}

func variableZeroValue() {
	// 1. variable names are before their types
	// 2. all variables must be used somewhere
	// 3. in GO, variables have its own initial value
	var a int
	var s string
	fmt.Println("in zero value")
	fmt.Println(a, s)
	fmt.Printf("%d %s\n", a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//bool, string
	//(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
	//byte <- same as (u)int8
	//rune(char type in GO) <- same as (u)int32
	//float32, float64, complex64, complex128 (实数)
	fmt.Println("in init value")
	var a int = 777
	var b, c int = 88, 99
	var d, e, f bool
	var s string = "qqq"
	fmt.Println(a, b, c, d, e, f, s)
}

func variableTpeDeduction() {
	// GO can know the type of the variable from its initial value
	//so no need to add a type behind if you give it an initial value
	fmt.Println("in type deduction")
	var a, g int = 777, 999
	var b, c, d, e, f = 88, 99, false, 9.8, "ppp"
	fmt.Println(a, b, c, d, e, f, g)
}

func variableShorter() {
	// "var" can also be omitted by := when you first define the variable
	// but for those "var" which outside the functions, you can't omit it
	// for second time or later, just use =, not :=
	fmt.Println("in shorter")
	b, c, d, e, f := 88, 99, false, 9.8, "ppp"
	b = 777
	fmt.Println(b, c, d, e, f)
}

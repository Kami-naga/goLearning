package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, ff)
	fmt.Println(bb, cc, dd)
	fmt.Println("-------------------")
	euler()
	triangle()
	consts()
	enums()
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
	//float32, float64
	//complex64, complex128 (实数)
	//-> complex64 = float32 + float32
	//-> complex128 = float64 + float64
	fmt.Println("in init value")
	var a int = 777
	var b, c int = 88, 99
	var d, e, f bool
	var s string = "qqq"
	fmt.Println(a, b, c, d, e, f, s)
}

func variableTypeDeduction() {
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

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	//verify e^(i*pi) + 1 = 0
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) //use 1i to show i
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         //same as above
	//but answer is not 0
	// change the format to show 0, just take 3 digits
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println()
}

func triangle() {
	// no implicit type change in GO!!
	// so you should do type changes explicitly
	var a, b int = 3, 4

	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	//can't be compiled if you write following code
	//c = math.Sqrt(a * a + b * b)
	// do these type changes explicitly!
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// const can also defined outside functions, scope is package
//const filename = "abc.txt"

//in other languages, we often let those consts written in upper form(like FILENAME)
//but in GO, capitalized letter have other meanings(when you capitalize the first word, it means the variable is public )
//so just use lower case
func consts() {
	//const filename string= "abc.txt" is also OK
	const filename = "abc.txt"
	//const a, b int = 3, 4 is also OK, const is something like text replacement, so it can be used as any type
	// but if you add int, c = int(math.Sqrt(a*a + b*b)) can't be compiled because of type change issue
	const a, b = 3, 4
	// below is also OK
	const (
		filename2 = "abc.txt"
		a2, b2    = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//const (
	//	dog   = 0
	//	cat   = 1
	//	mouse = 2
	//	bird  = 3
	//)
	//above codes can simplified as below
	// iota means autoincrement
	const (
		dog = iota
		_
		mouse
		bird
		bear
	)

	//iota can do more things~
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(dog, bear, mouse, bird)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

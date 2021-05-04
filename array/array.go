package main

import "fmt"

func main() {
	var arr1 [5]int
	// when use := , you must give it an initial value
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int //it means it has 4 arrays which has 5 elements

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//数组为值类型!! func(arr [10]int)会走值传递，即它会拷贝数组
	//所以我们可以使用指针,但即便如此也不方便,所以一般不直接使用数组，而是使用切片!
	//Array is not convenient! So use slice instead!
	printArray(arr1)
	// can't print arr2! because printArray need [5]int, but arr2 is [3]int
	//printArray(arr2)
	printArray(arr3)

	// Unlike C++,array name is not the head of the array, so & is necessary
	printArrayPointer(&arr1)
	printArrayPointer(&arr3)

	// use slice!
	printSlice(arr1[:])
	printSlice(arr2[:])
	printSlice(arr3[:])
}

func printArray(arr [5]int) {
	//in other languages, we might write codes like following one
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr3[i])
	//}
	//but in GO, we can use range(like for each in Java & python, but range in GO can get both i & v)
	// only v ? for _, v := range arr
	// only i ? for i := range arr
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayPointer(arr *[5]int) {
	arr[0] = 100 // same as (*arr)[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSlice(arr []int) {
	arr[0] = 100 // same as (*arr)[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

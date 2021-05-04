package main

import "fmt"

func main() {
	//slice不是值类型
	// slice is a view of array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6] // [2,6)
	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:](s1) = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:](s2) = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//reslice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//important point!!!!
	//in slice, there are 3 variable:
	//ptr: point to the first ele of the slice
	//len: tell the length of the slice(e.g. length of arr[2:6] is 4)
	//cap(capacity): tell the length from the ptr to the end of ARRAY
	//(e.g. cap of arr[2:6] is 6 because array is {0, 1, 2, 3, 4, 5, 6, 7} & ptr is 2, so its cap = 6(2,3,4,5,6,7))
	//so slice can extend! 只可向后扩展，不能向前扩展
	//s[i]不可超越len(s)，向后扩展不可超越底层数组cap(s)
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr = ", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // s2 = ?
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2)) // s2 = [5, 6] !!! why it can get 6 ??
	fmt.Println(s1[3:6])
	fmt.Println(s1[3:7])
}

func updateSlice(s []int) {
	s[0] = 100
}

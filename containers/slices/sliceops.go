package main

import "fmt"

func main() {
	fmt.Println("create slice")
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6}
	printSlice(s1) // len=cap=3

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("copy slice")
	copy(s2, s1)
	printSlice(s1)
	printSlice(s2)

	fmt.Println("delete elements from slice")
	s2 = append(s2[:3], s2[4:]...) //s2[:3] + s2[4:]
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	//for the first time , the init value is nil, but len&cap shows to be 0, not to be crashed
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

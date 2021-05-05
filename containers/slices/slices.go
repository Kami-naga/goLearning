package main

import "fmt"

func main() {
	//在GO中一切都是值传递
	//slice也不例外，但slice是一个array的视图，它是一个指针，指向存放真正元素的数据块
	//因此即便值传递，也只是指针复制了一下，原指针和新复制出的指针都指向同一数据
	//因此对slice的元素进行更改时，别的指向同一数据的slice以及原数组array都会发生变化
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
	//fmt.Println(s1[3:7]) it can't be compiled

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	//s4 and s5 no longer view arr.
	fmt.Println("arr = ", arr) //[0 1 2 3 4 5 6 10] where is 11 & 12?
	//添加元素时若>cap,系统会重新分配一个更大的底层数组(每次cap*2)，上边11，12就在新的大数组里头了
	//原小数组仍在，当没人用它了，则会被垃圾回收掉
	//由于值传递关系，我们一定要接受返回值，即写成 s= append(s, val)的形式
	//why？ 因为新append后，可能原数组cap不够了，此时会去指向新的大数组，即slice指针会变掉，如果还用原来的数组，新的值就莫得了
}

func updateSlice(s []int) {
	s[0] = 100
}

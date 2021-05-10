package main

import (
	"bufio"
	"fmt"
	"goLearning/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1) //postpone the execution of the code,it will be executed when return or panic, or other errors occurs
	defer fmt.Println(2) //postpone the execution of the code
	fmt.Println(3)
	panic("error occurred") //return
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		//it will print 30,29,...0, not 30,30...30,
		// which means i will be calculated first and stocked,
		//then when return or error occurred, it will print that
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // 直接用file写文件比较慢，所以用这个，这个先写进内存，大到一定程度一起倒进文件
	defer writer.Flush()            //从内存倒东西进文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

//GO中使用defer来进行资源管理
//defer确保调用在函数结束时发生
//参数在defer语句时计算
//defer列表为后进先出
//建了什么东西需要些指令与其成对时，用defer， 如Open/Close, Lock/Unlock, PrintHeader/PrintFooter
func main() {
	tryDefer2() //what's the result? 321 panic, so in defer, there is a stack
	writeFile("fib.txt")
}

package main

import (
	"fmt"
	"goLearning/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())
	//问：q := queue.Queue{1}的q和最后fmt.Println(q.IsEmpty())的q是同一个q，同一个slice吗？
	//答：不是哦~GO语言比较特殊，
	//push和pop可以改变这个q的值，因为push和pop定义中的接收者是指针接收者。
	//所以可以改变里面的值，即他把q改了
}

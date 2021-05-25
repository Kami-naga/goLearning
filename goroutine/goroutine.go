package main

import (
	"fmt"
	"time"
)

func printHello() {}

//coroutine features:
//1. coroutines are kind of light threads
//2. They are compiler/interpreter/VM level multitasking, not OS level
//3. Multiple coroutines may run on one or more threads
//4. non-preemptive multitasking,
//which means that coroutines hand over the control actively instead of having control taken away
// so when context switch happens, only a few params need to be saved
//(won't be stopped at some unexpected time,when you have to save more contexts),
//so it's fast!
//但有个问题，routine可能交不出控制权，导致就他一个一直loop，别的动不了，导致死循环
//（main函数也是一个goroutine,别人交不出控制权给main，所以main只能一直sleep不出来）
//从 go1.14，goroutine现在可以被异步抢占。因此没有函数调用的循环不再对调度器造成死锁或造成垃圾回收的大幅变慢。

//虽然之前底层实现是非抢占式的，但我们在写代码的时候其实和抢占式的切换没什么区别，
//我们也不应该去依赖go语言啥时切换来实现自己的逻辑

func main() {
	//runtime.GOMAXPROCS(1) //make it run on only 1 processor

	var a [10]int
	for i := 0; i < 10; i++ {
		//go printHello()
		//we can do it like above,but in most of time, we just write them directly

		//use `go` to do them parallel or the program will just print hello 0 endlessly
		//but only `go` here is not enough, because father and its sons are running simultaneously
		//so if father run faster, it will get to return and sons will be all killed (they are Coroutine协程)
		//so we let father wait for a millisecond at last

		//what happens if you write like this:
		//go func() -> use the i outside ?
		//it will cause race condition!!! => main will add i, coroutines will also use the same i,
		//so if main add i to 10, then it will get out of the loop,
		//and if at this time a coroutine used the i, it will get a[10] => array out of range!
		//so we need to let every coroutine has its private i
		//we can use `go run -race goroutine.go` to check if we have a race condition in the code
		go func(i int) {
			for {
				//fmt.Printf("Hello from goroutine %d\n", i) //进行print等io操作时，会有一个等待的过程，这时就会自动交出控制权给别人
				a[i]++
				//runtime.Gosched() //手动交出控制权 go schedule

			}
		}(i)

	}
	time.Sleep(time.Millisecond)
	fmt.Println(a) //race condition here! main may read  the a when coroutine are writing a! we will solve it by channel
}

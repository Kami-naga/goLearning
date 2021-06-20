package main

import (
	"fmt"
	"time"
)

func printHello() {}

// Subroutines(like function call) are special cases of more general program components, called coroutines.
// In contrast to the unsymmetric
// - Donnald Knuth "The Art of Computer Programming. Vol1"
//子程序是协程的特例

// in C++, we can use Boost.Coroutine to use coroutines
// in Java, it is not supported, but in some third party JVM, coroutine is supported
// in python, coroutine is a bit well supported compared to cpp & java,
// in python, you can use yield to implement coroutine, after python3.5, async def is added to support coroutine, which means coroutine should be decided when defining
// then in GO, coroutine is very well supported, for any function, just add a `go`, then you can send this func to the scheduler to treat it as a coroutine(goroutine) (more convenient than python)
// the scheduler will decide the position of this goroutine, maybe 1 goroutine in 1 thread, or more goroutines in 1 thread
// the scheduler will switch coroutines at a proper time, which is a bit different from traditional coroutine
// in traditional coroutine, we should decide when to switch, but in GO, scheduler decides it, so just write functions as normal
// some switching time(it may, not must):
// 1. I/O, select
// 2. channel
// 3. waiting for a lock
// 4. function call
// 5. runtime.Gosched()

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
	//如果这里的10改成1000，scheduler会开多少线程呢？
	//可以发现线程数一般不会超过我们cpu的核数，现在由scheduler来控制了，就没必要多开thread让系统来调度这些了（因为开销大）

	//对别的语言而言，一下子开十个，一百个线程差不多极限了，上千就不能单靠这样简单开线程了，得使用异步IO
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

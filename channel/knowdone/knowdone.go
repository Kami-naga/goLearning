package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("doWork %d received %c\n", id, n)
		//go func() { done <- true }()
		done <- true
	}
}

func chanDemo1() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	//we can replace the ugly time sleep by using done channel
	//time.Sleep(time.Millisecond)
	//but the above way(line38&42) will make it wait one by one which is meaningless,
	//we should let it done in free order which means done should be got at last,
	//so we then have following codes
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
	// but it will cause a dead lock because after receiving the first value,
	// it will wait for a done(line 25), but done will be sent after it receiving second value

	// how to solve it?
	// an easy way is to change the code in doWork
	// done <- true   ===>   go func(){done <- true}
	// then it won't be stuck
}

func chanDemo2() {
	// in most situation, we only wait for 1 value, so below code is enough in most situation
	// it will get all value1 before getting any value2
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func main() {

	//chanDemo1()
	//chanDemo2()
	chanDemo3()
}

func chanDemo3() {
	// go also gives us a library to deal with "waiting"
	//the difference between demo2 is that
	//the receiving of value1 & value2 will done parallel
	var wg sync.WaitGroup

	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}

	wg.Add(20) // or we can wg.Add(1) in below code
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

}

type worker2 struct {
	in chan int
	//it's not proper to add a waiting group directly in a worker,
	//so we use function to wrap the done operation
	done func()
}

func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in:   make(chan int),
		done: func() { wg.Done() },
	}
	go doWork2(id, w)
	return w
}

func doWork2(id int, worker worker2) {
	for n := range worker.in {
		fmt.Printf("doWork %d received %c\n", id, n)
		worker.done()
	}
}

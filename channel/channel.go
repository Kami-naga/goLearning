package main

import (
	"fmt"
	"time"
)

//CSP
//don't communicate by sharing memory; share memory by communicating
func createWorker(id int) chan<- int {
	//var c chan int   // c == nil
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	//way 1
	//for {
	//	n, ok := <-c
	//	//receiver will still get initial zero value(int:0, string:"") after sender closes,
	//	// so it should be checked
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %c\n", id, n)
	//}

	//way2
	//use range can do the same thing as above!
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
	}
	//how it will be if sender doesn't have a close? it will receive forever,
	//so do not forget to close if receive with range
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	// msg sent must be received or the program will get stuck
	//but we can give it a buffer and until the buffer gets full,
	//the channel can continue receiving msgs
	c := make(chan int, 3) // create a channel with buffer=3
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c' //program will not get stuck
	// c <- 4  //program will get stuck here
	time.Sleep(time.Millisecond)
}

func channelClose() { // if the sending data has an explicit end
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//channel as first-class citizen
	//chanDemo()

	//buffered channel
	//bufferedChannel()

	//channel close and range
	channelClose()
}

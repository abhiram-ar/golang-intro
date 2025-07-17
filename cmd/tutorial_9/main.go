package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // channel which can hold one element of type input
	go process(c)
	for i := range c { // we will again get deadlock - because after the processing 5 element, range c will still listing for channel
		fmt.Println(i)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("staring buffered channel ")

	buffC := make(chan int, 5) // buffered  channel of size 5
	go process(buffC)
	for i := range buffC {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c) // do this just before the function exists
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("exiting processs")
	// close(c)
}

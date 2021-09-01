package main

import "fmt"

// c chan<- int write vars enter the channel channel
// c <-chan int read vars leave the channel  channel

func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for v := range in {
		out <- 2 * v
	}
	close(out)
}

func Print(c <-chan int) {
	for v := range c {
		fmt.Printf("Id %d\n", v)
	}
}

// func main() {
// 	generator := make(chan int)
// 	double := make(chan int)

// 	go Generator(generator)
// 	go Double(generator, double)
// 	Print(double)
// }

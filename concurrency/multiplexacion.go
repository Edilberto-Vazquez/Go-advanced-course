package main

import (
	"fmt"
	"time"
)

func doSomething(i time.Duration, c chan int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	d1 := time.Second * 4
	d2 := time.Second * 2

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channel1 := <-c1:
			fmt.Println(channel1)
		case channel2 := <-c2:
			fmt.Println(channel2)
		}
	}

}

package main

import (
	"log"
	"time"
)

/*
	When the program starts, both goroutines are executed concurrently. Because there is nothing in place to prevent
	race conditions, the result might be affected based on the execution.

	In order to have the expected result goroutine1 needs to be executed first (x = 1), then goroutine2 runs resulting
	in x = 2. However, sometimes goroutine2 may run first and result in x = 1 as at that point x would have its original
	value.

	Expected:
	x:0
	x:2

	Expected:
	x:1
	x:2

	Race Condition:
	x:0
	x:1
*/

func goroutine1(x *int) {
	log.Println("first")
	*x++
}

func goroutine2(x *int) {
	log.Println("second")
	log.Printf("x: %d", *x)
	*x++
	log.Printf("x: %d", *x)
}

func main() {
	var x int

	go goroutine1(&x)
	go goroutine2(&x)

	time.Sleep(30 * time.Millisecond)
}

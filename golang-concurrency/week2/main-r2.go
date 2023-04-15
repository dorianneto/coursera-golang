package main

import (
	"fmt"
	"sync"
)

var X int = 0

func routine1(wg *sync.WaitGroup){
	originalX := X
	X = X + 3
	fmt.Printf("Routine 1 Original X: %d incremented X: %d\n", originalX, X)
	wg.Done()
}

func routine2(wg *sync.WaitGroup){
	originalX := X
	X = X + 5
	fmt.Printf("Routine 2 Original X: %d incremented X: %d\n", originalX, X)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go routine1(&wg)
	go routine2(&wg)
	wg.Wait()
}

//© - 213e7b0e-d378-4300-abf7-432bc079dad2
package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	leftCS, rightCS *ChopStick
	name int
}

var wg sync.WaitGroup
var counter int = 0
var mutex sync.Mutex

//Give permission by coming out of this method if the counter is 2 or less
//Each counter represents a philosopher eating food
func getPermission(wg1 *sync.WaitGroup) {
	defer wg1.Done()
	for {
		flag := false
		mutex.Lock()
		if counter < 2 {
			counter = counter+1
			flag = true
		}
		mutex.Unlock()

		if flag {
			break
		}
	}
}

//Decrement the counter as one philosopher is done eating
func decrementCounter() {
	mutex.Lock()
	counter = counter - 1
	mutex.Unlock()
}

func (p Philosopher) eat() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		var wg1 sync.WaitGroup
		wg1.Add(1)
		go getPermission(&wg1)
		//Block unless you got permission
		wg1.Wait()
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.name)
		fmt.Printf("finishing eating %d\n", p.name)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		decrementCounter()
	}
}

func main() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher { chopSticks[i], chopSticks[(i+1) % 5], i+1 }
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
	wg.Wait()
}
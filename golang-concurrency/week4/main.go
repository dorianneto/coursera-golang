package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	PHILOSOPHERS_ON_TABLE = 5
	MAX_ALLOWED_BY_HOST   = 2
	MAX_EACH_CAN_EAT      = 3
)

type chopstick struct{ sync.Mutex }

type philosopher struct {
	leftChopstick, rightChopstick *chopstick
	number                        int
}

func (p philosopher) eat(host *sync.WaitGroup) {
	defer host.Done()

	p.number++

	for i := 0; i < MAX_EACH_CAN_EAT; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat #%d - round %d \n", p.number, i+1)
		fmt.Printf("finishing eating #%d - round %d \n", p.number, i+1)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
}

func main() {
	chopsticks := make([]*chopstick, PHILOSOPHERS_ON_TABLE)
	for i := 0; i < PHILOSOPHERS_ON_TABLE; i++ {
		chopsticks[i] = &chopstick{}
	}

	philosophers := make([]*philosopher, PHILOSOPHERS_ON_TABLE)
	for i := 0; i < PHILOSOPHERS_ON_TABLE; i++ {
		philosophers[i] = &philosopher{
			number:         i,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	rand.Shuffle(len(philosophers), func(i, j int) {
		philosophers[i], philosophers[j] = philosophers[j], philosophers[i]
	})

	var host sync.WaitGroup

	host.Add(MAX_ALLOWED_BY_HOST)
	counter := 0

	for _, p := range philosophers {
		counter++

		go p.eat(&host)

		if counter > 1 {
			host.Wait()
			host.Add(MAX_ALLOWED_BY_HOST)
			counter = 0

			fmt.Print("\nThe previous philosophers got permission from host to eat at the same time.\n\n")
		}
	}
}

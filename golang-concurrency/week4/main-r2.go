package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type philosopher struct {
	id         int
	leftstick  *chopstick
	rightstick *chopstick
}

type chopstick struct {
	sync.Mutex
}

// The "host" mentioned in the assignment is the main() function, it has its own goroutine.
func main() {
	const PhilosophersCount int = 5
	const MaxConcurrentEaters int = 5

	var wg sync.WaitGroup

	chopsticks := make([]*chopstick, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		chopsticks[i] = new(chopstick)
	}

	philosophers := make([]*philosopher, PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		philosophers[i] = &philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%PhilosophersCount]}
	}

	EatersCountChannel := make(chan byte, MaxConcurrentEaters)

	wg.Add(PhilosophersCount)
	for i := 0; i < PhilosophersCount; i++ {
		go philosophers[i].eat(&wg, EatersCountChannel)
	}
	wg.Wait()
}

func (phil philosopher) eat(wg *sync.WaitGroup, EatersCountChannel chan byte) {
	for i := 1; i <= 3; i++ {
		EatersCountChannel <- 1
		randPickOrder := rand.Intn(2)
		if randPickOrder == 0 {
			phil.leftstick.Lock()
			phil.rightstick.Lock()
		} else {
			phil.rightstick.Lock()
			phil.leftstick.Lock()
		}

		fmt.Println("Philosopher", phil.id, ": Starting to eat - round", i)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Philosopher", phil.id, ": Finishing eating - round", i)

		if randPickOrder == 0 {
			phil.rightstick.Unlock()
			phil.leftstick.Unlock()
		} else {
			phil.leftstick.Unlock()
			phil.rightstick.Unlock()
		}

		<-EatersCountChannel
	}
	wg.Done()
}

package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	mu sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

type Host struct {
	numberOfEating int
	request        sync.Mutex
}

func (h *Host) AskForPermission() bool {
	h.request.Lock()
	if h.numberOfEating < 2 { //This is not a fair test, but does prevent deadlock.
		return true
	} else {
		return false
	}
}

func (p Philosopher) PickedUpChopsticks(god *Host) {
	god.numberOfEating -= 1
	god.request.Unlock()
}

func (p Philosopher) EatRice(god *Host, wg *sync.WaitGroup) {
	if god.AskForPermission() {
		god.numberOfEating += 1
		p.leftChopstick.mu.Lock()
		p.rightChopstick.mu.Lock()
		p.PickedUpChopsticks(god)
		fmt.Printf("Starting to eat #%d\n", p.number)
		fmt.Printf("Finished eating #%d\n", p.number)
		p.leftChopstick.mu.Unlock()
		p.rightChopstick.mu.Unlock()
		wg.Done()
	} else {
		p.EatRice(god, wg)
	}
}

func main() {
	var wg sync.WaitGroup
	var groupOfPhilosophers []Philosopher
	chopsticks := make([]Chopstick, 5)
	bob := Host{numberOfEating: 0}

	for i := 0; i < 5; i++ {
		phil := Philosopher{
			number:         i + 1,
			leftChopstick:  &chopsticks[(i+5)%5],
			rightChopstick: &chopsticks[(i+1)%5],
		}
		groupOfPhilosophers = append(groupOfPhilosophers, phil)
	}

	for i := 0; i < len(groupOfPhilosophers)*3; i++ {
		wg.Add(1)
		for i := 0; i < 5; i++ {
			go groupOfPhilosophers[i].EatRice(&bob, &wg)
			//note: It may seem to be running sequentially - its not, can increase parent loop to see.
		}
	}
	wg.Wait()

}

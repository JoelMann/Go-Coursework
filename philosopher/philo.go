// Implement the dining philosopher’s problem with the following constraints/modifications.

// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

// The host allows no more than 2 philosophers to eat concurrently.

// Each philosopher is numbered, 1 through 5.

// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

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
	if h.numberOfEating < 2 {
		return true
	} else {
		return false
	}
}

func (p Philosopher) ImDoneEatingMom(god *Host) {
	god.numberOfEating -= 1
	god.request.Unlock()
}

func (p Philosopher) EatRice(god *Host) {
	if god.AskForPermission() {
		god.numberOfEating += 1
		p.leftChopstick.mu.Lock()
		p.rightChopstick.mu.Lock()
		fmt.Printf("Philosopher %d is eating\n", p.number)
		p.leftChopstick.mu.Unlock()
		p.rightChopstick.mu.Unlock()
		p.ImDoneEatingMom(god)
	}
}

func main() {
	fmt.Println("sanity test")
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

	for {
		for i := 0; i < 5; i++ {
			go groupOfPhilosophers[i].EatRice(&bob)
		}
	}

}

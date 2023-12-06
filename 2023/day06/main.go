package main

import "fmt"

type race struct {
	time int
	dist int
}

func (r race) wins() int {
	total := 0
	for chargeTime := 1; chargeTime < r.time-1; chargeTime++ {
		travelTime := r.time - chargeTime
		if chargeTime*travelTime > r.dist {
			total += 1
		}
	}
	return total
}

func totalWins(races []race) int {
	total := 1
	for _, r := range races {
		total *= r.wins()
	}
	return total
}

func main() {
	races := []race{
		{time: 44, dist: 277},
		{time: 89, dist: 1136},
		{time: 96, dist: 1890},
		{time: 91, dist: 1768},
	}
	fmt.Println("Solution 1:", totalWins(races))

	races = []race{
		{time: 44899691, dist: 277113618901768},
	}
	fmt.Println("Solution 2:", totalWins(races))
}

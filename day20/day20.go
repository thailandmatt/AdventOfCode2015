package main

import (
	"fmt"
)

func main() {
	part2()
}

func part2() {
	target := 34000000

	//831600 ended up being the number
	for i := 820000; i < target; i++ {
		score := 0
		for j := i / 50; j <= i; j++ {
			if i%j == 0 && j*50 >= i {
				score += 11 * j
			}
		}

		if i%1000 == 0 {
			fmt.Println(i, ":", score)
		}
		if score > target {
			fmt.Println("FOUND", i, ":", score)
			break
		}
	}
}

func part1() {
	target := 34000000

	//ended up being 786240 but took a hot minute to find it
	for i := 780000; i < target; i++ {
		score := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				score += 10 * j
			}
		}

		if i%10000 == 0 {
			fmt.Println(i, ":", score)
		}
		if score > target {
			fmt.Println("FOUND", i, ":", score)
			break
		}
	}
}

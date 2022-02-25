package main

import "fmt"

func main() {
	//Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
	//Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
	//Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
	//Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1

	ingredients := [][]int{{4, -2, 0, 0, 5}, {0, 5, -1, 0, 8}, {-1, 0, 5, 0, 6}, {0, 0, -2, 2, 1}}

	//part 1 - brute force
	//trials count for curiosity
	trials := 0
	max := 0
	var maxForm [4]int
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-a-b; c++ {
				d := 100 - a - b - c

				//part 2
				cal := (a*ingredients[0][4] + b*ingredients[1][4] + c*ingredients[2][4] + d*ingredients[3][4])
				if cal != 500 {
					continue
				}

				score := 1
				for i := 0; i < 4; i++ {
					x1 := (a*ingredients[0][i] + b*ingredients[1][i] + c*ingredients[2][i] + d*ingredients[3][i])
					if x1 < 0 {
						x1 = 0
					}
					score *= x1
				}

				if score > max {
					max = score
					maxForm = [4]int{a, b, c, d}
				}

				trials++
			}
		}
	}

	fmt.Printf("%v trials, %v max, %v formula", trials, max, maxForm)

}

package main

import (
	"fmt"
	"io/ioutil"
)

type location struct {
	x, y int
}

func main() {
	input, err := ioutil.ReadFile("day3.txt")
	visits := make(map[location]int)
	part2Visits := make(map[location]int)

	if err != nil {
		panic("no day 3 input file")
	}

	var curr, santaCurr, roboCurr location

	visits[curr]++
	part2Visits[santaCurr]++
	part2Visits[roboCurr]++

	for i, c := range input {
		if c == '^' {
			curr.y++
			if i%2 == 0 {
				santaCurr.y++
			} else {
				roboCurr.y++
			}
		} else if c == 'v' {
			curr.y--
			if i%2 == 0 {
				santaCurr.y--
			} else {
				roboCurr.y--
			}
		} else if c == '>' {
			curr.x++
			if i%2 == 0 {
				santaCurr.x++
			} else {
				roboCurr.x++
			}
		} else if c == '<' {
			curr.x--
			if i%2 == 0 {
				santaCurr.x--
			} else {
				roboCurr.x--
			}
		}

		visits[curr]++

		if i%2 == 0 {
			part2Visits[santaCurr]++
		} else {
			part2Visits[roboCurr]++
		}
	}

	fmt.Printf("Houses with at least one (part 1):%v", len(visits))
	fmt.Printf("Houses with at least one (part 2):%v", len(part2Visits))
}

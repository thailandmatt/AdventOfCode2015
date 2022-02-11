package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("day1.txt")

	if err != nil {
		panic("no day 1 input file")
	}

	floor := 0
	basement := false

	for i, character := range input {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}

		//part 2
		if floor < 0 && !basement {
			fmt.Printf("First basement:%v\n", i+1)
			basement = true
		}
	}

	fmt.Printf("Ending Floor:%v", floor)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "1113222113"

	for iter := 0; iter < 50; iter++ {
		newInput := ""
		var curLetter byte
		curCount := 0
		for i := 0; i < len(input); i++ {
			if input[i] != curLetter {
				if curCount != 0 {
					newInput += strconv.Itoa(curCount) + string(curLetter)
				}
				curLetter = input[i]
				curCount = 1
			} else {
				curCount++
			}
		}
		newInput += strconv.Itoa(curCount) + string(curLetter)
		input = newInput
		fmt.Printf("%v - %v\n", iter, len(input))
	}
}

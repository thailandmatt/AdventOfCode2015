package main

import "fmt"

func main() {
	input := []byte("vzbxkghb")
	input = []byte("vzbxxyzz") //part 2

	for {
		input = increment(input)
		if isValidPassword(input) {
			break
		}
	}

	fmt.Printf("%v", string(input))
}

func isValidPassword(input []byte) bool {
	doubleCount := 0
	runCount := 0
	for i := 0; i < len(input); i++ {
		if i < len(input)-2 {
			if input[i] == input[i+1]-1 && input[i] == input[i+2]-2 {
				runCount++
			}
		}
	}

	for i := 0; i < len(input); i++ {
		if i < len(input)-1 && input[i] == input[i+1] {
			doubleCount++
			i++
		}
	}

	return doubleCount > 1 && runCount > 0
}

func increment(input []byte) []byte {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 122 {
			//z
			continue
		}

		input[i]++

		if input[i] == 105 || input[i] == 108 || input[i] == 111 {
			//i o or l
			input[i]++
		}

		for j := i + 1; j < len(input); j++ {
			input[j] = 97 //back to a
		}

		break
	}

	return input
}

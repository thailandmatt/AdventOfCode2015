package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	for scanner.Scan() {
		word := scanner.Text()

		hasRepeat := false
		hasDuplicate := false
		for i, c := range word {
			if i < len(word)-2 && rune(word[i+2]) == c {
				hasRepeat = true
			}

			if i < len(word)-4 {
				if strings.Contains(word[i+2:], word[i:i+2]) {
					hasDuplicate = true
				}
			}
		}

		if hasRepeat && hasDuplicate {
			niceCount++
		}
		//fmt.Printf("Word:%v, repeat:%v, duplicate:%v\n", word, hasRepeat, hasDuplicate)
	}

	fmt.Printf("Nice Words:%v", niceCount)
}

func part1() {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	vowels := "aeiou"
	badCombos := "ab cd pq xy "
	niceCount := 0
	for scanner.Scan() {
		vowelCount := 0
		doubleLetter := false
		badCombo := false

		word := scanner.Text()

		for i, c := range word {
			if strings.Contains(vowels, string(c)) {
				vowelCount++
			}

			if i < len(word)-1 {
				if rune(word[i+1]) == c {
					doubleLetter = true
				}

				if strings.Contains(badCombos, word[i:i+2]+" ") {
					badCombo = true
				}
			}
		}

		fmt.Printf("Word:%v, Vowels:%v, Double:%v, badCombo:%v\n", word, vowelCount, doubleLetter, badCombo)

		if vowelCount >= 3 && doubleLetter && !badCombo {
			niceCount++
		}
	}

	fmt.Printf("Nice Words:%v", niceCount)
}

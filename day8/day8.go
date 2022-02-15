package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLen := 0
	actualLen := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		newStr := "\""

		for i := 0; i < len(text); i++ {
			if text[i] == '"' {
				newStr += "\\\""
			} else if text[i] == '\\' {
				newStr += "\\\\"
			} else {
				newStr += string(text[i])
			}
		}

		newStr += "\""
		totalLen += len(text)
		actualLen += len(newStr)
		fmt.Printf("%v -> %v\n", text, newStr)
	}

	fmt.Printf("%v - %v = %v", actualLen, totalLen, actualLen-totalLen)
}

func partone() {
	file, err := os.Open("day8.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLen := 0
	actualLen := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		totalLen += len(text)
		actualStr := ""
		for i := 1; i < len(text)-1; i++ {
			if text[i] == '\\' && i < len(text)-2 {
				actualStr += string(text[i+1])
				if text[i+1] == 'x' {
					i = i + 3
				} else {
					i = i + 1
				}
			} else {
				actualStr += string(text[i])
			}
		}
		actualLen += len(actualStr)
		fmt.Printf("%v (%v) -> %v (%v)\n", text, len(text), actualStr, len(actualStr))
	}

	fmt.Printf("%v - %v = %v", actualLen, totalLen, actualLen-totalLen)
}

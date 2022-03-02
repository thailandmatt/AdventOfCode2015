package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day19.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := ""
	myMap := make(map[string][]string)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			scanner.Scan()
			start = scanner.Text()
		} else {
			split := strings.Split(text, " ")
			myMap[split[0]] = append(myMap[split[0]], split[2])
		}
	}

	//part 1
	set := make(map[string]bool)
	curStart := 0
	for i, x := range start {
		if unicode.IsUpper(x) && i != 0 {
			//replace last token
			for _, replace := range myMap[start[curStart:i]] {
				newStr := start[:curStart] + replace + start[i:]
				set[newStr] = true
			}
			curStart = i
		}
	}

	for _, replace := range myMap[start[curStart:]] {
		newStr := start[:curStart] + replace
		set[newStr] = true
	}

	fmt.Printf("Total combos:%v\n", len(set))

	//part 2
	fmt.Printf("Len:%v, (:%v, ):%v, ',':%v\n", len(start), strings.Count(start, "("), strings.Count(start, ")"), strings.Count(start, ","))

	//this doesn't work - you have to just use the math above
	//i'm sure i could do some backtrack thing but I don't have the
	//time right now
	reverseMap := make(map[string]string)
	var keys []string
	for k, v := range myMap {
		for _, x := range v {
			reverseMap[x] = k
			keys = append(keys, x)
		}
	}

	// sort.Slice(keys, func(i, j int) bool {
	// 	return len(keys[i]) <= len(keys[j])
	// })

	molecule := start
	steps := 0
	for {
		test := len(molecule)
		for _, key := range keys {
			if strings.Contains(molecule, key) {
				molecule = strings.Replace(molecule, key, reverseMap[key], 1)
				steps++
				break
			}
		}
		fmt.Println("Step", steps, molecule)
		if molecule == "e" {
			break
		} else if test == len(molecule) {
			panic("bad")
		}
	}
	fmt.Printf("Steps:%v", steps)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	partone()
	parttwo()
}

func parttwo() {
	file, err := os.Open("day16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	needed := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	have := make(map[string]map[string]int)

	//parse graph
	for scanner.Scan() {
		text := scanner.Text()
		name := text[:strings.Index(text, ":")]
		text = text[strings.Index(text, ": ")+2:]
		split := strings.Split(text, ",")

		have[name] = make(map[string]int)
		for _, s := range split {
			split2 := strings.Split(strings.TrimSpace(s), ":")
			have[name][strings.TrimSpace(split2[0])], _ = strconv.Atoi(strings.TrimSpace(split2[1]))
		}
	}

	//cats, trees >
	//pomeranians, goldfish <
	for aunt, auntMap := range have {
		found := true
		for k, v := range auntMap {
			if k == "cats" || k == "trees" {
				if v <= needed[k] {
					found = false
					break
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if v >= needed[k] {
					found = false
					break
				}
			} else if v != needed[k] {
				found = false
				break
			}
		}

		if found {
			fmt.Println(aunt)
		}
	}
}

func partone() {
	file, err := os.Open("day16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	needed := []string{
		"children: 3",
		"cats: 7",
		"samoyeds: 2",
		"pomeranians: 3",
		"akitas: 0",
		"vizslas: 0",
		"goldfish: 5",
		"trees: 3",
		"cars: 2",
		"perfumes: 1",
	}

	for scanner.Scan() {
		text := scanner.Text()
		name := text[:strings.Index(text, ":")]
		text = text[strings.Index(text, ": ")+2:]
		split := strings.Split(text, ",")

		found := true
		for _, s := range split {
			if !contains(needed, strings.TrimSpace(s)) {
				found = false
				break
			}
		}

		if found {
			fmt.Println(name)
		}
	}
}

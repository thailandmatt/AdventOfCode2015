package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	partone()
	parttwo()
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func parttwo() {
	input, err := ioutil.ReadFile("day12.txt")

	if err != nil {
		panic("no day 12 input file")
	}

	var f interface{}
	json.Unmarshal(input, &f)

	total := processUnknown(f)

	fmt.Printf("Part 2:%v\n", total)
}

func processUnknown(x interface{}) float64 {
	total := float64(0)

bigLoop:
	switch unk := x.(type) {
	case []interface{}:
		//array
		for _, val := range unk {
			total += processUnknown(val)
		}
	case float64:
		//int in an array
		total += unk
	case map[string]interface{}:
		//json object
		for _, val := range unk {
			if val == "red" {
				break bigLoop
			}
		}
		for _, val := range unk {
			total += processUnknown(val)
		}
	}
	return total
}

func partone() {
	input, err := ioutil.ReadFile("day12.txt")

	if err != nil {
		panic("no day 12 input file")
	}

	//I don't really want to mess with unmarshalling weird formatted json
	jsonStr := string(input)

	tokens := SplitAny(jsonStr, "[]{},: \n")

	//part 1 = add up all integers
	total := 0
	for _, token := range tokens {
		x, err := strconv.Atoi(token)
		if err == nil {
			total += x
		}
	}

	fmt.Printf("Part 1:%v\n", total)

}

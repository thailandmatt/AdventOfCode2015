package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	vals := make(map[string]uint16)
	processFile(vals)
	b := vals["a"]

	vals = make(map[string]uint16)
	vals["b"] = b
	processFile(vals)
}

func processFile(vals map[string]uint16) {
	var leftovers []string

	file, err := os.Open("day7.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		leftovers = processLine(text, vals, leftovers)
	}

	for {
		text := leftovers[0]
		leftovers = leftovers[1:]
		leftovers = processLine(text, vals, leftovers)

		if len(leftovers) == 0 {
			break
		}
	}

	fmt.Printf("%v+\n", vals)
}

func processLine(text string, vals map[string]uint16, leftovers []string) []string {
	split := strings.Split(text, " ")
	if split[1] == "->" {
		//can't override an existing signal (needed for part 2)
		_, has := vals[split[2]]
		if !has {
			//assignment
			v, err := strconv.Atoi(split[0])
			if err == nil {
				vals[split[2]] = uint16(v)
			} else {
				//see if we have a value for the operand
				_, ok := vals[split[0]]
				if ok {
					vals[split[2]] = vals[split[0]]
				} else {
					leftovers = append(leftovers, text)
				}
			}
		}
	} else if split[0] == "NOT" {
		//see if we have a value for the operand
		_, ok := vals[split[1]]
		if ok {
			vals[split[3]] = ^vals[split[1]]
		} else {
			leftovers = append(leftovers, text)
		}
	} else {
		//see if we have values for the operands
		x, ok1 := vals[split[0]]
		y, ok2 := vals[split[2]]
		a, err1 := strconv.Atoi(split[0])
		b, err2 := strconv.Atoi(split[2])

		if (ok1 || err1 == nil) && (ok2 || err2 == nil) {
			if err1 == nil {
				x = uint16(a)
			}
			if err2 == nil {
				y = uint16(b)
			}

			if split[1] == "AND" {
				vals[split[4]] = x & y
			} else if split[1] == "OR" {
				vals[split[4]] = x | y
			} else if split[1] == "LSHIFT" {
				vals[split[4]] = x << y
			} else if split[1] == "RSHIFT" {
				vals[split[4]] = x >> y
			}
		} else {
			leftovers = append(leftovers, text)
		}
	}

	return leftovers
}

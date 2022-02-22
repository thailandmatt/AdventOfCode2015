package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	prmt "github.com/gitchander/permutation"
)

func main() {
	file, err := os.Open("day13.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(map[string]map[string]int)
	//part 2
	graph["me"] = make(map[string]int)

	//parse graph
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")

		personMap, ok := graph[split[0]]
		if !ok {
			personMap = make(map[string]int)
			graph[split[0]] = personMap
			//part 2
			personMap["me"] = 0
			graph["me"][split[0]] = 0
		}

		other := strings.Replace(split[10], ".", "", -1)
		personMap[other], _ = strconv.Atoi(split[3])
		if split[2] == "lose" {
			personMap[other] = 0 - personMap[other]
		}
	}

	var people []string
	for k, _ := range graph {
		people = append(people, k)
	}

	p := prmt.New(prmt.StringSlice(people))
	max := 0
	order := ""
	for p.Next() {
		total := 0
		for i, person := range people {
			if i == 0 {
				total += graph[people[len(people)-1]][person]
				total += graph[person][people[len(people)-1]]
			} else {
				total += graph[people[i-1]][person]
				total += graph[person][people[i-1]]
			}
		}

		if total > max {
			max = total
			order = fmt.Sprintf("%v", people)
		}
	}

	fmt.Printf("%v - %v", max, order)
}

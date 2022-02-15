package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day9.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(map[string]map[string]int)

	//parse graph
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")

		m, ok := graph[split[0]]
		if !ok {
			m = make(map[string]int)
			graph[split[0]] = m
		}

		m[split[2]], _ = strconv.Atoi(split[4])

		//also do the reverse
		m, ok = graph[split[2]]
		if !ok {
			m = make(map[string]int)
			graph[split[2]] = m
		}

		m[split[0]], _ = strconv.Atoi(split[4])
	}

	var cities []string
	for k, _ := range graph {
		cities = append(cities, k)
	}

	//now search all possible paths
	min := -1
	path := ""
	for _, city := range cities {

		var newcities []string
		for _, other := range cities {
			if city != other {
				newcities = append(newcities, other)
			}
		}

		distance, newpath := searchPaths(newcities, graph, city, 0, city)
		if distance < min || min == -1 { //for part 2 change to >
			min = distance
			path = newpath
		}
	}

	fmt.Printf("Min: %v, %v", min, path)
}

func searchPaths(cities []string, graph map[string]map[string]int, curNode string, distanceSoFar int, pathSoFar string) (int, string) {
	if len(cities) == 0 {
		return distanceSoFar, pathSoFar
	}

	min := -1
	path := pathSoFar
	for _, city := range cities {
		var newcities []string
		for _, other := range cities {
			if city != other {
				newcities = append(newcities, other)
			}
		}

		distance, newpath := searchPaths(newcities, graph, city, distanceSoFar+graph[curNode][city], pathSoFar+","+city)

		if distance < min || min == -1 { //for part 2 change to >
			min = distance
			path = newpath
		}
	}

	return min, path
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	paperNeeded := 0
	ribbonNeeded := 0

	scanner := bufio.NewScanner(file)
	d := make([]int, 3)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "x")
		d[0], _ = strconv.Atoi(split[0])
		d[1], _ = strconv.Atoi(split[1])
		d[2], _ = strconv.Atoi(split[2])
		sort.Ints(d)

		paperNeeded += (2 * d[0] * d[1]) + (2 * d[0] * d[2]) + (2 * d[1] * d[2]) + (d[0] * d[1])

		ribbonNeeded += (d[0] * d[1] * d[2]) + (2 * (d[0] + d[1]))
	}

	fmt.Printf("Paper Needed:%v\n", paperNeeded)
	fmt.Printf("Ribbon Needed:%v", ribbonNeeded)
}

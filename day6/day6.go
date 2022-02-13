package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [1000][1000]bool
	var grid2 [1000][1000]int

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Replace(text, "turn ", "turn", -1)
		split := strings.Split(text, " ")
		from := strings.Split(split[1], ",")
		to := strings.Split(split[3], ",")

		x1, _ := strconv.Atoi(from[0])
		y1, _ := strconv.Atoi(from[1])
		x2, _ := strconv.Atoi(to[0])
		y2, _ := strconv.Atoi(to[1])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if split[0] == "turnon" {
					grid[x][y] = true
					grid2[x][y]++
				} else if split[0] == "turnoff" {
					grid[x][y] = false
					grid2[x][y]--
					if grid2[x][y] < 0 {
						grid2[x][y] = 0
					}
				} else {
					grid[x][y] = !grid[x][y]
					grid2[x][y] = grid2[x][y] + 2
				}
			}
		}
	}

	total := 0
	total2 := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				total++
			}
			total2 += grid2[x][y]
		}
	}

	fmt.Printf("On:%v\n", total)
	fmt.Printf("Total:%v", total2)
}

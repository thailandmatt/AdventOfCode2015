package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [100][100]int
	//parse graph
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x, c := range text {
			if c == '#' {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
		y++
	}

	grid[0][0] = 1
	grid[0][99] = 1
	grid[99][0] = 1
	grid[99][99] = 1

	//animate 100 steps
	for i := 0; i < 100; i++ {
		var newgrid [100][100]int

		for y := 0; y < 100; y++ {
			for x := 0; x < 100; x++ {
				c := 0
				//check all 8 directions
				if y > 0 && x > 0 && grid[y-1][x-1] == 1 {
					c++
				}

				if y > 0 && grid[y-1][x] == 1 {
					c++
				}

				if y > 0 && x < 99 && grid[y-1][x+1] == 1 {
					c++
				}

				if x > 0 && grid[y][x-1] == 1 {
					c++
				}

				if x < 99 && grid[y][x+1] == 1 {
					c++
				}

				if y < 99 && x > 0 && grid[y+1][x-1] == 1 {
					c++
				}

				if y < 99 && grid[y+1][x] == 1 {
					c++
				}

				if y < 99 && x < 99 && grid[y+1][x+1] == 1 {
					c++
				}

				//if on, and not 2 or 3 neighbors on, turn off
				if grid[y][x] == 1 && (c == 2 || c == 3) {
					//leave on
					newgrid[y][x] = 1
				}

				//if off, and 3 neighbors on, turn on
				if grid[y][x] == 0 && c == 3 {
					newgrid[y][x] = 1
				}

			}
		}

		grid = newgrid
		grid[0][0] = 1
		grid[0][99] = 1
		grid[99][0] = 1
		grid[99][99] = 1
	}

	total := 0
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			total += grid[y][x]
		}
	}

	fmt.Printf("%v", total)
}

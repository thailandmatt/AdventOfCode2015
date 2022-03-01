package main

import "fmt"

func main() {
	sizes := []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}
	targetSize := 150
	//sizes := []int{20, 15, 10, 5, 5}
	//targetSize := 25

	//dynamic programming - let the grid be targetSize + 1 "tall" and len(sizes) + 1 "long"
	//initialize [0,0] to 1
	//then for each size, work from targetSize back (i) and len(sizes) (j) back and set
	//grid[i + size, j] = grid[i, j - 1]

	//this will make each cell in the grid [x, y] represent the number of ways that you can achieve
	//X liters with Y containers

	//the possibilities that add up to targetSize will be in the grid[targetSize, ...] column of the grid
	//the first non-zero possibility on the last row is the number of ways to achieve the smallest possible container

	//initialize grid
	grid := make([][]int, targetSize+1)
	for i := 0; i <= targetSize; i++ {
		grid[i] = make([]int, len(sizes)+1)
	}

	//initialize first cell
	grid[0][0] = 1

	//populate dynamic table
	for _, size := range sizes {
		for i := targetSize - size; i >= 0; i-- {
			for j := len(sizes); j > 0; j-- {
				grid[i+size][j] += grid[i][j-1]
			}
		}
	}

	partOne := 0
	partTwo := -1
	for i := 0; i < len(sizes); i++ {
		partOne += grid[targetSize][i]
		if grid[targetSize][i] != 0 && partTwo == -1 {
			partTwo = grid[targetSize][i]
		}
	}
	fmt.Printf("Part one - %v\n", partOne)
	fmt.Printf("Part two - %v", partTwo)
}

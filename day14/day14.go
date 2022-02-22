package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reindeer struct {
	name     string
	speed    int
	flying   int
	resting  int
	distance int
	points   int
}

func main() {
	file, err := os.Open("day14.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var Reindeers []*Reindeer

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")

		r := &Reindeer{
			name: split[0],
		}

		r.speed, _ = strconv.Atoi(split[3])
		r.flying, _ = strconv.Atoi(split[6])
		r.resting, _ = strconv.Atoi(split[13])
		r.distance = 0

		Reindeers = append(Reindeers, r)
	}

	//2503 is puzzle input
	for i := 0; i < 2503; i++ {
		max := 0
		for _, r := range Reindeers {
			x := i % (r.flying + r.resting)
			if x < r.flying {
				r.distance += r.speed
			}
			if r.distance > max {
				max = r.distance
			}
		}

		for _, r := range Reindeers {
			if r.distance == max {
				r.points++
			}
		}
	}

	for _, r := range Reindeers {
		fmt.Printf("%+v\n", r)
	}

}

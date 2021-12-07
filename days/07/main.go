package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func crabWalk(distance int) int {
	return (distance * (distance + 1)) / 2
}

func identity(distance int) int {
	return distance
}

func part1(positions []int, min, max int) int {
	return helper(positions, min, max, identity)
}

func part2(positions []int, min, max int) int {
	return helper(positions, min, max, crabWalk)
}

func helper(positions []int, min, max int, moveCost func(distance int) int) int {
	cheapest := math.MaxInt32

	for k := min; k <= max; k++ {
		cost := 0

		for _, position := range positions {
			if position > k {
				cost += moveCost(position - k)
			} else if position < k {
				cost += moveCost(k - position)
			}
		}

		if cost < cheapest {
			cheapest = cost
		}
	}

	return cheapest
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var positions []int = make([]int, 0)
	min := math.MaxInt32
	max := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, value := range strings.Split(line, ",") {
			position, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}

			positions = append(positions, position)

			if position < min {
				min = position
			}

			if position > max {
				max = position
			}
		}
	}

	fmt.Println(part1(positions, min, max))
	fmt.Println(part2(positions, min, max))
}

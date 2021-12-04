package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(measurements []int) int {
	sum := 0

	for _, measurement := range measurements {
		sum += measurement
	}

	return sum
}

func part1(measurements []int) int {
	increases := 0

	for k := 1; k < len(measurements); k++ {
		if measurements[k] > measurements[k-1] {
			increases += 1
		}
	}

	return increases
}

func part2(measurements []int) int {
	increases := 0

	for k := 3; k < len(measurements); k++ {
		prev := sum(measurements[k-3 : k])
		cur := sum(measurements[k-2 : k+1])

		if cur > prev {
			increases += 1
		}
	}

	return increases
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var measurements []int = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		measurements = append(measurements, depth)
	}

	fmt.Println(part1(measurements))
	fmt.Println(part2(measurements))
}

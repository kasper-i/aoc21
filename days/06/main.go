package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(fishes []int) int {
	return helper(fishes, 80)
}

func part2(fishes []int) int {
	return helper(fishes, 256)
}

func helper(fishes []int, iterations int) int {
	var timers map[int]int = make(map[int]int)

	for k := 0; k <= 8; k++ {
		timers[k] = 0
	}

	for _, fish := range fishes {
		timers[fish] += 1
	}

	for k := 0; k < iterations; k++ {
		spawns := timers[0]

		for i := 0; i < 8; i++ {
			timers[i] = timers[i+1]
		}

		timers[6] += spawns
		timers[8] = spawns
	}

	total := 0

	for _, timers := range timers {
		total += timers
	}

	return total
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fishes []int = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		for _, fish := range strings.Split(line, ",") {
			timer, err := strconv.Atoi(fish)

			if err != nil {
				log.Fatal(err)
			}

			fishes = append(fishes, timer)
		}
	}

	fmt.Println(part1(fishes))
	fmt.Println(part2(fishes))
}

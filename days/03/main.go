package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(readings []int64) uint64 {
	var gamma uint64 = 0
	var epsilon uint64 = 0

	for k := 11; k >= 0; k-- {
		var ones int = 0

		for _, reading := range readings {
			bit := (reading >> k) & 0x1
			ones += int(bit)
		}

		if ones > 500 {
			gamma += (1 << k)
		} else {
			epsilon += (1 << k)
		}
	}

	return gamma * epsilon
}

func part2(readings []int64) int64 {
	oxygen := helper(readings, "oxygen")
	co2 := helper(readings, "co2")

	return oxygen * co2
}

func helper(readings []int64, rating string) int64 {
	for k := 11; k >= 0; k-- {
		if len(readings) == 1 {
			return readings[0]
		}

		ones := 0

		for _, reading := range readings {
			bit := (reading >> k) & 0x1
			ones += int(bit)
		}

		var criteria int

		if len(readings)%2 == 0 && ones == len(readings)/2 {
			criteria = 1
		} else if ones > len(readings)/2 {
			criteria = 1
		} else {
			criteria = 0
		}

		if rating == "co2" {
			criteria = ^criteria & 1
		}

		var selection []int64 = make([]int64, 0)

		for _, reading := range readings {
			bit := (reading >> k) & 0x1
			if int(bit) == criteria {
				selection = append(selection, reading)
			}
		}

		readings = selection
	}

	return readings[0]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var readings []int64 = make([]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()

		reading, err := strconv.ParseInt(line, 2, 16)
		if err != nil {
			log.Fatal(err)
		}

		readings = append(readings, reading)
	}

	fmt.Println(part1(readings))
	fmt.Println(part2(readings))
}

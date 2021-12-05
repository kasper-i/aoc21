package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(lines []Line) int {
	var points map[Coordinate]int = make(map[Coordinate]int, 0)

	for _, line := range(lines) {
		x1, x2, y1, y2 := line.x1, line.x2, line.y1, line.y2

		if x1 == x2 || y1 == y2 {
			processLine(line, points)
		}		
	}

	return countOverlaps(points)
}

func part2(lines []Line) int {
	var points map[Coordinate]int = make(map[Coordinate]int, 0)

	for _, line := range(lines) {
		processLine(line, points)
	}

	return countOverlaps(points)
}

func processLine(line Line, vents map[Coordinate]int) {
	x1, x2, y1, y2 := line.x1, line.x2, line.y1, line.y2
		
	xSlope := 1
	if x2 < x1 {
		xSlope = -1
	}

	ySlope := 1
	if y2 < y1 {
		ySlope = -1
	}

	for x, y := x1, y1;; {
		coord := Coordinate{x: x, y: y}
		
		if _, ok := vents[coord]; !ok {
			vents[coord] = 1
		} else {
			vents[coord] = vents[coord] + 1
		}

		if x == x2 && y == y2 {
			break
		}

		if x != x2 {
			x += xSlope
		}

		if y != y2 {
			y += ySlope
		}
	}		
}

func countOverlaps(points map[Coordinate]int) int {
	overlaps := 0

	for _, point := range(points) {
		if point > 1 {
			overlaps += 1
		}
	}

	return overlaps
}

type Coordinate struct {
	x int;
	y int;
}

type Line struct {
	x1 int;
	x2 int;
	y1 int;
	y2 int;
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []Line = make([]Line, 0)

	for scanner.Scan() {
		input := scanner.Text()

		line := Line{}
		fmt.Sscanf(input, "%d,%d -> %d,%d", &line.x1, &line.y1, &line.x2, &line.y2)

		lines = append(lines, line)
	}


	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

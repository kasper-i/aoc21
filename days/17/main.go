package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}

func simulateTrajectory(ixv, iyv, x1, x2, y1, y2 int) bool {
	xpos, ypos := 0, 0
	xvel, yvel := ixv, iyv

	for xpos < x2 && ypos > y1 {
		xpos += xvel
		ypos += yvel

		if xpos >= x1 && xpos <= x2 && ypos >= y1 && ypos <= y2 {
			return true
		}

		if xvel > 0 {
			xvel -= 1
		}

		yvel -= 1
	}

	return false
}

func part1(y1, y2 int) int {
	maxHeight := 0

	for yiv := y1; yiv < -y1; yiv++ {
		if simulateTrajectory(1, yiv, 0, 1, y1, y2) {
			heightReached := (yiv * (yiv + 1)) / 2

			if heightReached > maxHeight {
				maxHeight = heightReached
			}
		}
	}

	return maxHeight
}

func part2(x1, x2, y1, y2 int) int {
	var hits map[Coord]bool = make(map[Coord]bool)

	for xiv := 1; xiv <= x2; xiv++ {
		for yiv := y1; yiv < -y1; yiv++ {
			if simulateTrajectory(xiv, yiv, x1, x2, y1, y2) {
				hits[Coord{x: xiv, y: yiv}] = true
			}
		}
	}

	return len(hits)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	var x1, x2, y1, y2 int

	fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	fmt.Println(part1(y1, y2))
	fmt.Println(part2(x1, x2, y1, y2))
}

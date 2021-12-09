package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type Location struct {
	height   int
	lowPoint bool
	basin    bool
}

type Coord struct {
	x int
	y int
}

type HeightMap [][]*Location

func (heightMap HeightMap) exploreBasin(x, y int) *Coord {
	if y < 0 || y > len(heightMap)-1 || x < 0 || x > len(heightMap[0])-1 {
		return nil
	}

	location := heightMap[y][x]

	if !location.basin && location.height != 9 {
		location.basin = true
		return &Coord{x: x, y: y}
	}

	return nil
}

func part1(heightMap HeightMap) int {
	sum := 0

	for y, row := range heightMap {
		for x, location := range row {
			height := location.height

			up := math.MaxInt32
			down := math.MaxInt32
			left := math.MaxInt32
			right := math.MaxInt32

			if y != len(heightMap)-1 {
				up = heightMap[y+1][x].height
			}

			if y != 0 {
				down = heightMap[y-1][x].height
			}

			if x != 0 {
				left = heightMap[y][x-1].height
			}

			if x != len(heightMap[0])-1 {
				right = heightMap[y][x+1].height
			}

			if height < up && height < down && height < left && height < right {
				sum += height + 1
				location.lowPoint = true
				location.basin = true
			}
		}
	}

	return sum
}

func part2(heightMap HeightMap) int {
	var lowPoints []Coord = make([]Coord, 0)

	for y, row := range heightMap {
		for x, location := range row {
			if location.lowPoint {
				lowPoints = append(lowPoints, Coord{x: x, y: y})
			}
		}
	}

	var basinSizes []int = make([]int, 0)

	for _, lowPointCoord := range lowPoints {
		var origins []Coord = []Coord{lowPointCoord}
		basinSize := 1

		for len(origins) != 0 {
			var next []Coord = make([]Coord, 0)

			for _, origin := range origins {
				x, y := origin.x, origin.y

				if continuation := heightMap.exploreBasin(x, y+1); continuation != nil {
					next = append(next, *continuation)
				}

				if continuation := heightMap.exploreBasin(x, y-1); continuation != nil {
					next = append(next, *continuation)
				}

				if continuation := heightMap.exploreBasin(x-1, y); continuation != nil {
					next = append(next, *continuation)
				}

				if continuation := heightMap.exploreBasin(x+1, y); continuation != nil {
					next = append(next, *continuation)
				}
			}

			basinSize += len(next)
			origins = next
		}

		basinSizes = append(basinSizes, basinSize)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var heightMap HeightMap = make(HeightMap, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var row []*Location = make([]*Location, len(line))

		for idx, height := range line {
			number, err := strconv.Atoi(string(height))
			if err != nil {
				log.Fatal(err)
			}

			row[idx] = &Location{
				height: number,
			}
		}

		heightMap = append(heightMap, row)
	}

	fmt.Println(part1(heightMap))
	fmt.Println(part2(heightMap))
}

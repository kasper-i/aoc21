package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type EnergyMap [][]int

type Coord struct {
	x int
	y int
}

func part1(energyMap EnergyMap) int {
	totalFlashCount := 0

	for k := 0; k < 100; k++ {
		totalFlashCount += iterateOnce(energyMap)
	}

	return totalFlashCount
}

func part2(energyMap EnergyMap, offset int) int {
	for k := 0; k < 1000; k++ {
		simulFlashes := iterateOnce(energyMap)

		if simulFlashes == 100 {
			return k + 1 + offset
		}
	}

	return 0
}

func iterateOnce(energyMap EnergyMap) int {
	flashCount := 0
	var flashes []Coord = make([]Coord, 0)

	for y, row := range energyMap {
		for x, energy := range row {
			energyMap[y][x] = energy + 1

			if energyMap[y][x] > 9 {
				flashes = append(flashes, Coord{x: x, y: y})
			}
		}
	}

	for ; len(flashes) > 0; {
		var next []Coord = make([]Coord, 0)

		for _, flash := range flashes {
			flashCount += 1
			x, y := flash.x, flash.y
			
			if coord := energyMap.energize(x - 1, y - 1); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x, y - 1); coord != nil {
				next = append(next, *coord)
			}		

			if coord := energyMap.energize(x + 1, y - 1); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x + 1, y); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x + 1, y + 1); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x, y + 1); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x - 1, y + 1); coord != nil {
				next = append(next, *coord)
			}

			if coord := energyMap.energize(x - 1, y); coord != nil {
				next = append(next, *coord)
			}
		}

		flashes = next;
	}

	for y, row := range energyMap {
		for x, energy := range row {
			if energy > 9 {
				energyMap[y][x] = 0
			}
		}
	}

	return flashCount
}

func (energyMap EnergyMap) energize(x, y int) *Coord {
	if x < 0 || y < 0 || x > 9 || y > 9 {
		return nil
	}

	energy := energyMap[y][x]

	if energy > 9 {
		return nil
	}

	energyMap[y][x] = energy + 1

	if energyMap[y][x] > 9 {
		return &Coord{x: x, y: y}
	}

	return nil	
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var energyMap EnergyMap = make(EnergyMap, 0)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int = make([]int, 10)
		
		for idx, num := range line {
			row[idx] = int(num - '0')
		}

		energyMap = append(energyMap, row)
	}

	fmt.Println(part1(energyMap))
	fmt.Println(part2(energyMap, 100))
}

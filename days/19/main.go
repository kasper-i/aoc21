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
	z int
}

type Scanner struct {
	number int
	beacons []Coord
}

func (coord Coord) format() string {
	return fmt.Sprintf("%d,%d,%d", coord.x, coord.y, coord.z)
}

func getOffset(src, dst Coord) Coord {
	return Coord{
		x: dst.x - src.x,
		y: dst.y - src.y,
		z: dst.z - src.z,
	}
}

func withOffset(beacon, offset Coord) Coord {
	return Coord{
		x: beacon.x + offset.x,
		y: beacon.y + offset.y,
		z: beacon.z + offset.z,
	}
}

func checkOverlap(s1, s2 Scanner, requiredOverlaps int) (bool, Coord) {
	for _, origin := range s1.beacons {
		for k := 0; k < len(s2.beacons); k++ {
			offset := getOffset(s2.beacons[k], origin)
			overlapCounter := 1

			for j := 0; j < len(s2.beacons); j++ {
				if j == k {
					continue
				}

				needle := withOffset(s2.beacons[j], offset)

				for _, s1Beacon := range s1.beacons {
					if needle == s1Beacon {
						overlapCounter += 1
						break
					}
				}
			}

			if overlapCounter >= requiredOverlaps {
				return true, offset
			}
		}
	}

	return false, Coord{}
}

func main() {
	file, err := os.Open("tiny")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	var scanners []Scanner = make([]Scanner, 0)
	scannerNumber := 0
	var scanner Scanner	

	for lineScanner.Scan() {
		line := lineScanner.Text()

		if line == "" {
			scanners = append(scanners, scanner)
		} else if line[0:3] == "---" {
			scanner = Scanner{
				number: scannerNumber,
				beacons: make([]Coord, 0),
			}
			scannerNumber += 1
		} else {
			beacon := Coord{}
			fmt.Sscanf(line, "%d,%d,%d", &beacon.x, &beacon.y, &beacon.z)
			scanner.beacons = append(scanner.beacons, beacon)
		}		
	}

	scanners = append(scanners, scanner)

	for _, scanner := range scanners {
		fmt.Printf("--- scanner %d ---\n", scanner.number)
		for _, beacon := range scanner.beacons {
			fmt.Printf("%d,%d,%d\n", beacon.x, beacon.y, beacon.z)
		}
		fmt.Println()
	}

	if overlap, offset := checkOverlap(scanners[0], scanners[1], 3); overlap {
		fmt.Printf("found overlap with offset: %s\n", offset.format())
	} else {
		fmt.Println("no overlap")
	}

	//fmt.Println(part1(roots))
}

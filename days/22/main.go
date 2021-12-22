package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(cuboids []Cuboid) int {
	var enabled map[Vector]bool = make(map[Vector]bool)

	for _, cuboid := range cuboids {
		if cuboid.x1 < -50 || cuboid.x2 > 50 || cuboid.y1 < -50 || cuboid.y2 > 50 || cuboid.z1 < -50 || cuboid.z2 > 50 {
			continue
		}

		for x := cuboid.x1; x <= cuboid.x2; x++ {
			for y := cuboid.y1; y <= cuboid.y2; y++ {
				for z := cuboid.z1; z <= cuboid.z2; z++ {
					if cuboid.on {
						enabled[Vector{x: x, y: y, z: z}] = true
					} else {
						delete(enabled, Vector{x: x, y: y, z: z})
					}
				}
			}
		}
	}

	return len(enabled)
}

func part2(cuboids []Cuboid) int {
	var enabled map[Cuboid]bool = make(map[Cuboid]bool)

	for _, cuboid := range cuboids {
		if !cuboid.on {
			continue
		}

		foundOverlap := false

		for enabledCuboid := range enabled {
			if cuboid.hasOverlap(enabledCuboid) {
				foundOverlap = true
				fmt.Printf("%s overlaps with %s\n", cuboid.format(), enabledCuboid.format())
				break
			}
		}

		if !foundOverlap {
			fmt.Printf("adding %s\n", cuboid.format())
			enabled[cuboid] = true
		}
	}

	return 0
}

func doRangesOverlap(x1, x2, y1, y2 int) bool {
	return x1 <= y2 && y1 <= x2
}

type Vector struct {
	x int
	y int
	z int
}

type Cuboid struct {
	on     bool
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

func (cuboid Cuboid) format() string {
	return fmt.Sprintf("x=%d..%d,y=%d..%d,z=%d..%d", cuboid.x1, cuboid.x2, cuboid.y1, cuboid.y2, cuboid.z1, cuboid.z2)
}

func (cuboid Cuboid) size() int {
	return (cuboid.x2 - cuboid.x1) * (cuboid.y2 - cuboid.y1) * (cuboid.z2 - cuboid.z1)
}

func (c1 Cuboid) hasOverlap(c2 Cuboid) bool {
	return doRangesOverlap(c1.x1, c1.x2, c2.x1, c2.x2) || doRangesOverlap(c1.y1, c1.y2, c2.y1, c2.y2) || doRangesOverlap(c1.z1, c1.z2, c2.z1, c2.z2)
}

func (c1 Cuboid) sub(c2 Cuboid) []Cuboid {
	if !c1.hasOverlap(c2) {
		return []Cuboid{c1}
	}

	return []Cuboid{}
}

func main() {
	file, err := os.Open("small")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cuboids []Cuboid = make([]Cuboid, 0)

	for scanner.Scan() {
		line := scanner.Text()
		cuboid := Cuboid{}
		var on string
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &cuboid.x1, &cuboid.x2, &cuboid.y1, &cuboid.y2, &cuboid.z1, &cuboid.z2)

		if on == "on" {
			cuboid.on = true
		}

		cuboids = append(cuboids, cuboid)
	}

	//fmt.Println(part1(cuboids))
	fmt.Println(part2(cuboids))
}

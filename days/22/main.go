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

func main() {
	file, err := os.Open("input")
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

	fmt.Println(part1(cuboids))
}

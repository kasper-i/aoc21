package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Reactor struct {
	cuboids map[Cuboid]bool
	count int
}

func (reactor *Reactor) enable(cuboid Cuboid) {
	subCuboids := []Cuboid{ cuboid }

	for len(subCuboids) > 0 {
		nextSubCuboids := make([]Cuboid, 0)

		for _, subCuboid := range subCuboids {
			foundOverlap := false
			isFullyContained := false

			for enabledCuboid := range reactor.cuboids {
				if enabledCuboid.fullyContains(subCuboid) {
					isFullyContained = true
					break
				}

				if subCuboid.hasOverlap(enabledCuboid) {
					foundOverlap = true

					for _, subCuboid := range subCuboid.split(enabledCuboid) {
						if !subCuboid.hasOverlap(enabledCuboid) {
							nextSubCuboids = append(nextSubCuboids, subCuboid)
						}
					}

					break
				}
			}

			if isFullyContained {
				continue
			}

			if !foundOverlap {
				reactor.cuboids[subCuboid] = true
				reactor.count += subCuboid.size()
			}
		}

		subCuboids = nextSubCuboids
	}
}

func (reactor *Reactor) disable(cuboid Cuboid) {
	toDisable := make([]Cuboid, 0)
	toPutBack := make([]Cuboid, 0)

	for otherCuboid := range reactor.cuboids {
		if cuboid.fullyContains(otherCuboid) {
			toDisable = append(toDisable, otherCuboid)
		} else if cuboid.hasOverlap(otherCuboid) {
			toDisable = append(toDisable, otherCuboid)

			for _, subCuboid := range otherCuboid.split(cuboid) {
				if !subCuboid.hasOverlap(cuboid) {
					toPutBack = append(toPutBack, subCuboid)
				}
			}
		}
	}

	for _, c := range toDisable {
		delete(reactor.cuboids, c)
		reactor.count -= c.size()
	}


	for _, c := range toPutBack {
		reactor.cuboids[c] = true
		reactor.count += c.size()
	}
}

func rebootReactor(cuboids []Cuboid) int {
	reactor := &Reactor{}
	reactor.cuboids = make(map[Cuboid]bool)

	for _, cuboid := range cuboids {
		if cuboid.on {
			reactor.enable(cuboid)
		} else {
			reactor.disable(cuboid)
		}
	}

	return reactor.count
}

type Pair struct {
	a int
	b int
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
	return (cuboid.x2 - cuboid.x1 + 1) * (cuboid.y2 - cuboid.y1 + 1) * (cuboid.z2 - cuboid.z1 + 1)
}

func (c1 Cuboid) hasOverlap(c2 Cuboid) bool {
	return doRangesOverlap(c1.x1, c1.x2, c2.x1, c2.x2) && doRangesOverlap(c1.y1, c1.y2, c2.y1, c2.y2) && doRangesOverlap(c1.z1, c1.z2, c2.z1, c2.z2)
}

func (c1 Cuboid) fullyContains(c2 Cuboid) bool {
	return c1.x1 <= c2.x1 && c1.x2 >= c2.x2 && c1.y1 <= c2.y1 && c1.y2 >= c2.y2 && c1.z1 <= c2.z1 && c1.z2 >= c2.z2
}

func (c1 Cuboid) split(c2 Cuboid) []Cuboid {
	smallCuboids := make([]Cuboid, 0)

	mx1, mx2 := max(c2.x1, c1.x1), min(c2.x2, c1.x2)
	my1, my2 := max(c2.y1, c1.y1), min(c2.y2, c1.y2)
	mz1, mz2 := max(c2.z1, c1.z1), min(c2.z2, c1.z2)

	xCuts := []Pair{ {c1.x1, mx1-1}, {mx1, mx2}, {mx2+1, c1.x2} }
	yCuts := []Pair{ {c1.y1, my1-1}, {my1, my2}, {my2+1, c1.y2} }
	zCuts := []Pair{ {c1.z1, mz1-1}, {mz1, mz2}, {mz2+1, c1.z2} }

	for _, xp := range xCuts {
		for _, yp := range yCuts {
			for _, zp := range zCuts {
				if xp.b < xp.a || yp.b < yp.a || zp.b < zp.a {
					continue
				}

				c := Cuboid{
					x1: xp.a, x2: xp.b,
					y1: yp.a, y2: yp.b,
					z1: zp.a, z2: zp.b,
				}

				smallCuboids = append(smallCuboids, c)
			}
		}
	}

	return smallCuboids
}

func doRangesOverlap(x1, x2, y1, y2 int) bool {
	return x1 <= y2 && y1 <= x2
}

func min(v1, v2 int) int {
	if (v1 <= v2) {
		return v1
	}

	return v2
}

func max(v1, v2 int) int {
	if (v1 >= v2) {
		return v1
	}

	return v2
}

func part1(cuboids []Cuboid) int {
	small := make([]Cuboid, 0)

	for _, cuboid := range cuboids {
		if cuboid.x1 < -50 || cuboid.x2 > 50 || cuboid.y1 < -50 || cuboid.y2 > 50 || cuboid.z1 < -50 || cuboid.z2 > 50 {
			continue
		}

		small = append(small, cuboid)
	}

	return rebootReactor(small)
}

func part2(cuboids []Cuboid) int {
	return rebootReactor(cuboids)
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
	fmt.Println(part2(cuboids))
}

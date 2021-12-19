package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Scanner struct {
	number  int
	beacons []Vector
	offset  Vector
	aligned bool
}

type RotationMatrix struct {
	m11, m12, m13 int
	m21, m22, m23 int
	m31, m32, m33 int
}

func (a RotationMatrix) multiply(b RotationMatrix) RotationMatrix {
	return RotationMatrix{
		m11: a.m11*b.m11 + a.m12*b.m21 + a.m13*b.m31, m12: a.m11*b.m12 + a.m12*b.m22 + a.m13*b.m32, m13: a.m11*b.m13 + a.m12*b.m23 + a.m13*b.m33,
		m21: a.m21*b.m11 + a.m22*b.m21 + a.m23*b.m31, m22: a.m21*b.m12 + a.m22*b.m22 + a.m23*b.m32, m23: a.m21*b.m13 + a.m22*b.m23 + a.m23*b.m33,
		m31: a.m31*b.m11 + a.m32*b.m21 + a.m33*b.m31, m32: a.m31*b.m12 + a.m32*b.m22 + a.m33*b.m32, m33: a.m31*b.m13 + a.m32*b.m23 + a.m33*b.m33,
	}
}

func (rm RotationMatrix) format() string {
	out := ""
	out += fmt.Sprintf("%d %d %d\n", rm.m11, rm.m12, rm.m13)
	out += fmt.Sprintf("%d %d %d\n", rm.m21, rm.m22, rm.m23)
	out += fmt.Sprintf("%d %d %d", rm.m31, rm.m32, rm.m33)

	return out
}

func sin(degree int) int {
	switch degree {
	case 0:
		return 0
	case 90:
		return 1
	case 180:
		return 0
	case 270:
		return -1
	}

	return 0
}

func cos(degree int) int {
	switch degree {
	case 0:
		return 1
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	}

	return 0
}

func makeRx(theta int) RotationMatrix {
	return RotationMatrix{
		m11: 1, m12: 0, m13: 0,
		m21: 0, m22: cos(theta), m23: -sin(theta),
		m31: 0, m32: sin(theta), m33: cos(theta),
	}
}

func makeRy(theta int) RotationMatrix {
	return RotationMatrix{
		m11: cos(theta), m12: 0, m13: sin(theta),
		m21: 0, m22: 1, m23: 0,
		m31: -sin(theta), m32: 0, m33: cos(theta),
	}
}

func makeRz(theta int) RotationMatrix {
	return RotationMatrix{
		m11: cos(theta), m12: -sin(theta), m13: 0,
		m21: sin(theta), m22: cos(theta), m23: 0,
		m31: 0, m32: 0, m33: 1,
	}
}

type Vector struct {
	x int
	y int
	z int
}

func (v1 Vector) add(v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}

func (v1 Vector) sub(v2 Vector) Vector {
	return Vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
		z: v1.z - v2.z,
	}
}

func (coord Vector) format() string {
	return fmt.Sprintf("%d,%d,%d", coord.x, coord.y, coord.z)
}

func (v Vector) rotate(rm RotationMatrix) Vector {
	return Vector{
		x: rm.m11*v.x + rm.m12*v.y + rm.m13*v.z,
		y: rm.m21*v.x + rm.m22*v.y + rm.m23*v.z,
		z: rm.m31*v.x + rm.m32*v.y + rm.m33*v.z,
	}
}

func checkOverlap(s1, s2 Scanner, rotations []RotationMatrix, requiredOverlaps int) (bool, Vector, RotationMatrix) {
	for _, origin := range s1.beacons {
		for k := 0; k < len(s2.beacons); k++ {
			for _, rm := range rotations {
				offset := origin.sub(s2.beacons[k].rotate(rm))
				overlaps := 1

				for j := 0; j < len(s2.beacons); j++ {
					if j == k {
						continue
					}

					needle := s2.beacons[j].rotate(rm).add(offset)

					for _, s1Beacon := range s1.beacons {
						if needle == s1Beacon {
							overlaps += 1
							break
						}
					}
				}

				if overlaps >= requiredOverlaps {
					return true, offset, rm
				}
			}
		}
	}

	return false, Vector{}, RotationMatrix{}
}

func locateBeacons(scanners []Scanner, rotations []RotationMatrix) []Vector {
	var beacons map[Vector]bool = make(map[Vector]bool)
	var origins []int = make([]int, 1)

	for _, beacon := range scanners[0].beacons {
		beacons[beacon] = true
	}

	for len(origins) > 0 {
		var nextOrigins []int = make([]int, 0)

		for _, originIndex := range origins {
			for k := 0; k < len(scanners); k++ {
				if scanners[k].aligned {
					continue
				}

				if overlap, offset, rm := checkOverlap(scanners[originIndex], scanners[k], rotations, 12); overlap {
					totalOffset := scanners[originIndex].offset.add(offset)

					for i, beacon := range scanners[k].beacons {
						rotatedBeacon := beacon.rotate(rm)
						scanners[k].beacons[i].x = rotatedBeacon.x
						scanners[k].beacons[i].y = rotatedBeacon.y
						scanners[k].beacons[i].z = rotatedBeacon.z

						beacons[rotatedBeacon.add(totalOffset)] = true
					}

					nextOrigins = append(nextOrigins, k)
					scanners[k].offset = totalOffset
					scanners[k].aligned = true
				}
			}
		}

		origins = nextOrigins
	}

	var beaconPositions []Vector = make([]Vector, 0)

	for beacon := range beacons {
		beaconPositions = append(beaconPositions, beacon)
	}

	return beaconPositions
}

func part1(beacons []Vector) int {
	return len(beacons)
}

func part2(scanners []Scanner) int {
	maxManhattanDistance := 0

	for _, s1 := range scanners {
		for _, s2 := range scanners {
			p1, p2 := s1.offset, s2.offset
			distance := abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)

			if distance > maxManhattanDistance {
				maxManhattanDistance = distance
			}
		}
	}

	return maxManhattanDistance
}

func abs(value int) int {
	if value < 0 {
		return (-1) * value
	}

	return value
}

func main() {
	file, err := os.Open("input")
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
				number:  scannerNumber,
				beacons: make([]Vector, 0),
			}
			scannerNumber += 1
		} else {
			beacon := Vector{}
			fmt.Sscanf(line, "%d,%d,%d", &beacon.x, &beacon.y, &beacon.z)
			scanner.beacons = append(scanner.beacons, beacon)
		}
	}

	scanners = append(scanners, scanner)

	degrees := []int{0, 90, 180, 270}
	var rotationSet map[RotationMatrix]bool = make(map[RotationMatrix]bool)
	var rotations []RotationMatrix = make([]RotationMatrix, 0)

	for _, xRot := range degrees {
		for _, yRot := range degrees {
			for _, zRot := range degrees {
				rm := makeRx(xRot).multiply(makeRy(yRot)).multiply(makeRz(zRot))
				rotationSet[rm] = true
			}
		}
	}

	for rm := range rotationSet {
		rotations = append(rotations, rm)
	}

	beacons := locateBeacons(scanners, rotations)

	fmt.Println(part1(beacons))
	fmt.Println(part2(scanners))
}

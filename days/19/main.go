package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RotationMatrix struct {
	m11, m12, m13 int
	m21, m22, m23 int
	m31, m32, m33 int
}

func (a RotationMatrix) multiply(b RotationMatrix) RotationMatrix {
	return RotationMatrix{
		m11: a.m11*b.m11+a.m12*b.m21+a.m13*b.m31, m12: a.m11*b.m12+a.m12*b.m22+a.m13*b.m32, m13: a.m11*b.m13+a.m12*b.m23+a.m13*b.m33,
		m21: a.m21*b.m11+a.m22*b.m21+a.m23*b.m31, m22: a.m21*b.m12+a.m22*b.m22+a.m23*b.m32, m23: a.m21*b.m13+a.m22*b.m23+a.m23*b.m33,
		m31: a.m31*b.m11+a.m32*b.m21+a.m33*b.m31, m32: a.m31*b.m12+a.m32*b.m22+a.m33*b.m32, m33: a.m31*b.m13+a.m32*b.m23+a.m33*b.m33,
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

func makeVector(x, y, z int) Vector {
	return Vector{
		x: x,
		y: y,
		z: z,
	}
}

type Scanner struct {
	number int
	beacons []Vector
}

func (coord Vector) format() string {
	return fmt.Sprintf("%d,%d,%d", coord.x, coord.y, coord.z)
}

func getOffset(src, dst Vector) Vector {
	return Vector{
		x: dst.x - src.x,
		y: dst.y - src.y,
		z: dst.z - src.z,
	}
}

func withOffset(beacon, offset Vector) Vector {
	return Vector{
		x: beacon.x + offset.x,
		y: beacon.y + offset.y,
		z: beacon.z + offset.z,
	}
}


func rotate(v Vector, rm RotationMatrix) Vector {
	return Vector{
		x: rm.m11 * v.x + rm.m12 * v.y + rm.m13 * v.z,
		y: rm.m21 * v.x + rm.m22 * v.y + rm.m23 * v.z,
		z: rm.m31 * v.x + rm.m32 * v.y + rm.m33 * v.z,
	}
}

func checkOverlap(s1, s2 Scanner, rm RotationMatrix, requiredOverlaps int) (bool, Vector) {
	for _, origin := range s1.beacons {
		for k := 0; k < len(s2.beacons); k++ {
			offset := getOffset(rotate(s2.beacons[k], rm), origin)
			overlapCounter := 1

			for j := 0; j < len(s2.beacons); j++ {
				if j == k {
					continue
				}

				needle := withOffset(rotate(s2.beacons[j], rm), offset)

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

	return false, Vector{}
}

func part1(scanners []Scanner, rotations []RotationMatrix) {
	for _, rm := range rotations {
		if overlap, offset := checkOverlap(scanners[0], scanners[1], rm, 12); overlap {
			fmt.Printf("overlap with offset %s\n", offset.format())
		}
	}
}

func main() {
	file, err := os.Open("small")
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

//	for _, scanner := range scanners {
//		fmt.Printf("--- scanner %d ---\n", scanner.number)
//		for _, beacon := range scanner.beacons {
//			fmt.Printf("%d,%d,%d\n", beacon.x, beacon.y, beacon.z)
//		}
//		fmt.Println()
//	}
//
//	if overlap, offset := checkOverlap(scanners[0], scanners[1], 3); overlap {
//		fmt.Printf("found overlap with offset: %s\n", offset.format())
//	} else {
//		fmt.Println("no overlap")
//	}

	//fmt.Println(rotate(makeVector(404,-588,-901), makeRx(90).multiply(makeRy(90)).multiply(makeRz(90))).format())

	degrees := []int{ 0, 90, 180, 270 }
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

	//fmt.Println(len(rotations))
	for rm := range rotationSet {
	//	fmt.Println(rm.format())
	//	fmt.Println()
		rotations = append(rotations, rm)
	}



	part1(scanners, rotations)
}

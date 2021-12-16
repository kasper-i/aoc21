package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var versionSum uint64 = 0

func bsToInt(bitstring []int8) uint64 {
	var out uint64 = 0

	for pos, b := range bitstring {
		if b == 1 {
			out += 1 << (len(bitstring) - pos - 1)
		}
	}

	return out
}

func runOperation(packetType uint64, values []uint64) []uint64 {
	var result uint64 = 0

	switch packetType {
	case 0:
		for _, value := range values {
			result += value
		}
	case 1:
		result = 1

		for _, value := range values {
			result *= value
		}
	case 2:
		result = math.MaxUint64

		for _, value := range values {
			if value < result {
				result = value
			}
		}
	case 3:
		result = 0

		for _, value := range values {
			if value > result {
				result = value
			}
		}
	case 5:
		if values[0] > values[1] {
			result = 1
		}
	case 6:
		if values[0] < values[1] {
			result = 1
		}
	case 7:
		if values[0] == values[1] {
			result = 1
		}
	}

	return []uint64{uint64(result)}
}

func evaluateExpression(bitstring []int8) (int, []uint64) {
	var values []uint64 = make([]uint64, 0)

	version := bsToInt(bitstring[0:3])
	packetType := bsToInt(bitstring[3:6])

	versionSum += version

	if packetType != 4 {
		lengthType := bsToInt(bitstring[6:7])

		switch lengthType {
		case 0:
			totalLength := int(bsToInt(bitstring[7:22]))
			containedPackets := bitstring[22 : 22+totalLength]

			offset := 0
			for offset < len(containedPackets) {
				read, containedValues := evaluateExpression(containedPackets[offset:])
				offset += read

				values = append(values, containedValues...)
			}

			return 22 + totalLength, runOperation(packetType, values)
		case 1:
			numSubPackets := int(bsToInt(bitstring[7:18]))

			offset := 18
			for n := 0; n < numSubPackets; n++ {
				read, containedValues := evaluateExpression(bitstring[offset:])
				offset += read

				values = append(values, containedValues...)
			}

			return offset, runOperation(packetType, values)
		}
	} else {
		var concat []int8 = make([]int8, 0)
		read := 0

		for g := 0;; g += 5 {
			group := bitstring[6+g : 6+g+5]
			last := group[0] == 0
			concat = append(concat, group[1:5]...)
			read += 5

			if last {
				break
			}
		}

		value := bsToInt(concat)
		return 6 + read, []uint64{value}
	}

	return 0, make([]uint64, 0)
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
	var bitstring []int8 = make([]int8, 4*len(line))

	bitIdx := 0
	for k := 0; k < len(line); k++ {
		hex := line[k : k+1]
		bin, _ := strconv.ParseUint(hex, 16, 4)

		for k := 3; k >= 0; k-- {
			mask := 1 << k
			if int(bin) & mask == mask {
				bitstring[bitIdx] = 1
			} else {
				bitstring[bitIdx] = 0
			}

			bitIdx += 1
		}
	}

	_, results := evaluateExpression(bitstring)
	fmt.Println(versionSum)
	fmt.Println(results[0])
}

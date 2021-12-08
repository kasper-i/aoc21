package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type SevenSeg = uint8

type Entry struct {
	signals []SevenSeg
	codes   []SevenSeg
}

const A = 1 << 6
const B = 1 << 5
const C = 1 << 4
const D = 1 << 3
const E = 1 << 2
const F = 1 << 1
const G = 1

func part1(entries []Entry) int {
	count := 0

	for _, entry := range entries {
		for _, code := range entry.codes {
			switch countSegments(code) {
			case 2, 3, 4, 7:
				count += 1
			}
		}
	}

	return count
}

func part2(entries []Entry) int {
	sum := 0

	for _, entry := range entries {
		number := 0
		decodedSignals := decodeSignals(entry.signals)
		codes := entry.codes

		number += 1000 * decode(codes[0], decodedSignals)
		number += 100 * decode(codes[1], decodedSignals)
		number += 10 * decode(codes[2], decodedSignals)
		number += 1 * decode(codes[3], decodedSignals)

		sum += number
	}

	return sum
}

func decode(code SevenSeg, decodedSignals map[SevenSeg]int) int {
	if code == 0b1111111 {
		return 8
	}

	switch countSegments(code) {
		case 2:
			return 1
		case 3:
			return 7
		case 4:
			return 4
	}

	if hit, ok := decodedSignals[code]; ok {
		return hit
	}

	return 0
}

func convert(str string) SevenSeg {
	var output SevenSeg = 0

	for _, position := range str {
		switch position {
		case 'a':
			output += A
		case 'b':
			output += B
		case 'c':
			output += C
		case 'd':
			output += D
		case 'e':
			output += E
		case 'f':
			output += F
		case 'g':
			output += G
		}
	}

	return output
}

func countSegments(value SevenSeg) int {
	count := 0

	count += int(1 & (value >> 6))
	count += int(1 & (value >> 5))
	count += int(1 & (value >> 4))
	count += int(1 & (value >> 3))
	count += int(1 & (value >> 2))
	count += int(1 & (value >> 1))
	count += int(1 & (value))

	return count
}

func decodeSignals(signals []SevenSeg) map[SevenSeg]int {
	//  0000
	// 1    2
	// 1    2
	//  3333
	// 4    5
	// 4    5
	//  6666

	var s0 SevenSeg = 0
	var s1 SevenSeg = 0
	var s2 SevenSeg = 0
	var s3 SevenSeg = 0
	var s4 SevenSeg = 0
	var s5 SevenSeg = 0
	var s6 SevenSeg = 0

	for _, signal := range signals {
		if countSegments(signal) == 2 {
			s2 = signal
			s5 = signal
			break
		}
	}

	for _, signal := range signals {
		if countSegments(signal) == 3 {
			s0 = signal & ^s2
			break
		}
	}

	for _, signal := range signals {
		if countSegments(signal) == 4 {
			s1 = signal & ^s2
			s3 = signal & ^s2
			break
		}
	}

	// find 9
	for _, signal := range signals {
		mask := (s0 | s1 | s5)
		if countSegments(signal) == 6 && signal&mask == mask {
			s6 = signal & ^mask
			break
		}
	}

	// find 5
	for _, signal := range signals {
		mask := s0 | s1 | s6
		if countSegments(signal) == 5 && signal&mask == mask {
			s5 = signal & ^mask
			s2 = s2 & ^s5
			break
		}
	}

	// find 3
	for _, signal := range signals {
		mask := s0 | s2 | s5 | s6
		if countSegments(signal) == 5 && signal&mask == mask {
			s3 = signal & ^(s0 | s2 | s5 | s6)
			s1 = s1 & ^s3
			break
		}
	}

	s4 = ^(s0 | s1 | s2 | s3 | s5 | s6)

	var out map[SevenSeg]int = make(map[SevenSeg]int)

	out[0x7f&(s0|s2|s3|s4|s6)] = 2
	out[0x7f&(s0|s2|s3|s5|s6)] = 3
	out[0x7f&(s0|s1|s3|s5|s6)] = 5
	out[0x7f&(s0|s1|s3|s4|s5|s6)] = 6
	out[0x7f&(s0|s1|s2|s3|s5|s6)] = 9

	return out
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []Entry = make([]Entry, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		
		var signals []SevenSeg = make([]SevenSeg, 10)
		var codes []SevenSeg = make([]SevenSeg, 4)

		for index, formatted := range parts[0:10] {
			signals[index] = convert(formatted)
		}

		for index, formatted := range parts[11:15] {
			codes[index] = convert(formatted)
		}

		entry := Entry{
			signals: signals,
			codes:   codes,
		}

		entries = append(entries, entry)
	}

	fmt.Println(part1(entries))
	fmt.Println(part2(entries))
}

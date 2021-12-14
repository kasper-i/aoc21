package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type CharCounter map[byte]int

type PolymerPairs map[string]int

type InsertionRules map[string]byte

func (counter CharCounter) add(char byte, quantity int) {
	if _, ok := counter[char]; !ok {
		counter[char] = quantity
	} else {
		counter[char] += quantity
	}
}

func (pairs PolymerPairs) add(pair string, quantity int) {
	if _, ok := pairs[pair]; !ok {
		pairs[pair] = quantity
	} else {
		pairs[pair] += quantity
	}
}

func part1(template string, rules InsertionRules) int {
	return helper(template, rules, 10)
}

func part2(template string, rules InsertionRules) int {
	return helper(template, rules, 40)
}

func helper(template string, rules InsertionRules, iterations int) int {
	var counter CharCounter = make(CharCounter)
	var polymer PolymerPairs = make(PolymerPairs)

	for i := 0; i < len(template); i++ {
		counter.add(template[i], 1)
	}

	for k := 0; k < len(template)-1; k++ {
		pair := template[k : k+2]
		polymer.add(pair, 1)
	}

	for k := 0; k < iterations; k++ {
		var nextPolymer PolymerPairs = make(PolymerPairs)

		for pair, count := range polymer {
			insertion, ok := rules[pair]

			if !ok {
				continue
			}

			lPair := string([]byte{pair[0], insertion})
			rPair := string([]byte{insertion, pair[1]})

			nextPolymer.add(lPair, count)
			nextPolymer.add(rPair, count)

			counter.add(insertion, count)
		}

		polymer = nextPolymer
	}

	var maxRune byte
	var minRune byte
	max := 0
	min := math.MaxInt64

	for k, v := range counter {
		if v > max {
			max = v
			maxRune = k
		}

		if v < min {
			min = v
			minRune = k
		}
	}

	return counter[maxRune] - counter[minRune]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineno := 0
	var template string
	var pairInsertionRules InsertionRules = make(InsertionRules)

	for scanner.Scan() {
		lineno += 1
		line := scanner.Text()

		if lineno == 1 {
			template = line
		} else if lineno >= 3 {
			var pair string
			var r byte

			fmt.Sscanf(line, "%2s -> %c", &pair, &r)
			pairInsertionRules[pair] = r
		}
	}

	fmt.Println(part1(template, pairInsertionRules))
	fmt.Println(part2(template, pairInsertionRules))
}

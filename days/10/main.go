package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		score, _ := checkLine(line)
		sum += score
	}

	return sum
}

func part2(lines []string) int {
	var scores []int = make([]int, 0)

	for _, line := range lines {
		lineScore := 0
		score, stack := checkLine(line)

		if score > 0 {
			continue
		}

		for i := len(stack) - 1; i >= 0; i-- {
			char := stack[i]

			switch char {
			case '(':
				lineScore = lineScore*5 + 1
			case '[':
				lineScore = lineScore*5 + 2
			case '{':
				lineScore = lineScore*5 + 3
			case '<':
				lineScore = lineScore*5 + 4
			}
		}

		scores = append(scores, lineScore)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	return scores[len(scores)/2]
}

func checkLine(line string) (int, []rune) {
	var stack []rune = make([]rune, 0)

	for _, char := range line {
		switch char {
		case '(', '[', '{', '<':
			stack = append(stack, char)
		default:
			pop := stack[len(stack)-1]

			switch char {
			case ')':
				if pop != '(' {
					return 3, stack
				}
			case ']':
				if pop != '[' {
					return 57, stack
				}
			case '}':
				if pop != '{' {
					return 1197, stack
				}
			case '>':
				if pop != '<' {
					return 25137, stack
				}
			}

			stack = stack[0 : len(stack)-1]
		}
	}

	return 0, stack
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string = make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

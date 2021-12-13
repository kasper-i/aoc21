package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Matrix [][]int

type FoldInstruction struct {
	along rune
	at int
}

func (paper Matrix) width() int {
	return len(paper[0])
}

func (paper Matrix) height() int {
	return len(paper)
}

func cutPaper(width, height int) Matrix {
	var paper Matrix = make(Matrix, height)

	for y := 0; y < height; y++ {
		paper[y] = make([]int, width)
	}

	return paper
}

func countDots(paper Matrix) int {
	count := 0

	for _, row := range paper {
		for _, value := range row {
			if value > 0 {
				count += 1
			}
		}
	}

	return count
}

func dump(paper Matrix) {
	for _, row := range paper {
		for _, value := range row {
			if value > 0 {
				fmt.Print("▩")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func part1(paper Matrix, instructions []FoldInstruction) int {
	folded := performFoldInstructions(paper, instructions[0:1])
	return countDots(folded)
}

func part2(paper Matrix, instructions []FoldInstruction) {
	paper = performFoldInstructions(paper, instructions)
	dump(paper)
}

func performFoldInstructions(paper Matrix, instructions []FoldInstruction) Matrix {
	for _, fold := range instructions {
		switch fold.along {
		case 'x':
			paper = foldX(paper, fold.at)
		case 'y':
			paper = foldY(paper, fold.at)
		}
	}

	return paper
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func foldX(dotPaper Matrix, foldAt int) Matrix {
	foldedPaper := cutPaper(foldAt, dotPaper.height())

	for y, row := range dotPaper {
		for x, value := range row {
			if x == foldAt {
				continue
			}

			if x < foldAt {
				foldedPaper[y][x] += value
			} else {
				foldedPaper[y][dotPaper.width()-1-x] += value
			}
		}
	}

	return foldedPaper
}

func foldY(dotPaper Matrix, foldAt int) Matrix {
	foldedPaper := cutPaper(dotPaper.width(), foldAt)

	for y, row := range dotPaper {
		if y == foldAt {
			continue
		}

		for x, value := range row {
			if y < foldAt {
				foldedPaper[y][x] += value
			} else {
				foldedPaper[dotPaper.height()-1-y][x] += value
			}
		}
	}

	return foldedPaper
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxX, maxY := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	paper := cutPaper(maxX+1, maxY+1)
	dotsRead := false
	var instructions []FoldInstruction = make([]FoldInstruction, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if dotsRead {
			var instruction FoldInstruction
			fmt.Sscanf(line, "fold along %c=%d", &instruction.along, &instruction.at)
			instructions = append(instructions, instruction)
		} else {
			var x, y int
			_, err := fmt.Sscanf(line, "%d,%d", &x, &y)

			if err != nil {
				dotsRead = true
				continue
			}

			paper[y][x] = 1
		}
	}

	fmt.Println(part1(paper, instructions))
	part2(paper, instructions)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BingoBoard = [][]BingoNumber

type BingoNumber struct {
	value int;
	encircled bool;
}

func encircleNumber(number int, board BingoBoard) (bool, int, int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if board[r][c].value == number {
				board[r][c].encircled = true
				return true, r, c
			}
		}
	}

	return false, 0, 0
}

func checkBingo(board BingoBoard, row, col int) bool {
	rowBingo := true
	for c := 0; c < 5; c++ {
		if !board[row][c].encircled {
			rowBingo = false
		}
	}

	if rowBingo {
		return true
	}

	colBingo := true
	for r := 0; r < 5; r++ {
		if !board[r][col].encircled {
			colBingo = false
		}
	}

	if colBingo {
		return true
	}

	return false
}

func calculateScore(board BingoBoard) int {
	score := 0

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !board[r][c].encircled {
				score += board[r][c].value
			}
		}
	}

	return score
}

func part1(sequence []int, boards []BingoBoard) int {
	for _, calledNumber := range(sequence) {
		for _, board := range(boards) {
			hit, row, col := encircleNumber(calledNumber, board)

			if hit {
				if checkBingo(board, row, col) {
					return calledNumber * calculateScore(board)
				}
			}
		}
	}

	return 0
}
func part2(sequence []int, boards []BingoBoard) int {
	lastScore := 0
	var boardsWithBingo []bool = make([]bool, len(boards))

	for _, calledNumber := range(sequence) {
		for boardIndex, board := range(boards) {
			if boardsWithBingo[boardIndex] {
				continue
			}

			hit, row, col := encircleNumber(calledNumber, board)

			if hit {
				if checkBingo(board, row, col) {
					lastScore = calledNumber * calculateScore(board)
					boardsWithBingo[boardIndex] = true
				}
			}
		}
	}

	return lastScore
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequence []int = make([]int, 0)
	var data []BingoNumber = make([]BingoNumber, 0)

	re := regexp.MustCompile("\\s+")

	for scanner.Scan() {
		line := scanner.Text()

		if len(sequence) == 0 {
			for _, number := range(strings.Split(line, ",")) {
				parsed, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)		
				}

				sequence = append(sequence, parsed)
			}

			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		for _, number := range(re.Split(line, -1)) {
			if number == "" {
				continue
			}

			parsed, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)		
			}

			data = append(data, BingoNumber{value: parsed, encircled: false})
		}
	}

	var boards []BingoBoard = make([]BingoBoard, 0)

	for k := 0; k < len(data) - 24; k += 25 {
		board := make(BingoBoard, 5)

		board[0] = data[k:k+5] 
		board[1] = data[k+5:k+10]
		board[2] = data[k+10:k+15]
		board[3] = data[k+15:k+20]
		board[4] = data[k+20:k+25]

		boards = append(boards, board)
	}

	fmt.Println(part1(sequence, boards))
	fmt.Println(part2(sequence, boards))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(commands []command) int {
	horizontal := 0
	depth := 0

	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			horizontal += cmd.distance
		case "up":
			depth -= cmd.distance
		case "down":
			depth += cmd.distance
		}
	}

	return horizontal * depth
}

func part2(commands []command) int {
	horizontal := 0
	aim := 0
	depth := 0

	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			horizontal += cmd.distance
			depth += aim * cmd.distance
		case "up":
			aim -= cmd.distance
		case "down":
			aim += cmd.distance
		}
	}

	return horizontal * depth
}

type command struct {
	direction string
	distance  int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var commands []command = make([]command, 0)

	for scanner.Scan() {
		line := scanner.Text()

		cmd := command{}
		fmt.Sscanf(line, "%s %d", &cmd.direction, &cmd.distance)

		commands = append(commands, cmd)
	}

	fmt.Println(part1(commands))
	fmt.Println(part2(commands))
}

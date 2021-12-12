package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type CaveMap map[string][]string

func isSmallCave(name string) bool {
    for _, r := range name {
        if !unicode.IsLower(r) {
            return false
        }
    }

    return true
}

func isPreviouslyVisited(path []string, cave string) bool {
	for _, visitedCave := range path {
		if visitedCave == cave {
			return true
		}
	}

	return false
}

func explore(cave string, path []string, visitTwiceOption bool, caveMap CaveMap) int {
	endCount := 0

	if cave == "end" {
		return 1
	}

	for _, nextCave := range caveMap[cave] {
		if nextCave == "start" {
			continue
		}

		if isSmallCave(nextCave) && isPreviouslyVisited(path, nextCave) {
			if visitTwiceOption {
				endCount += explore(nextCave, append(path, cave), false, caveMap)
			}

			continue
		}

		endCount += explore(nextCave, append(path, cave), visitTwiceOption, caveMap)
	}

	return endCount
}

func part1(caveMap CaveMap) int {
	return explore("start", make([]string, 0), false, caveMap)
}

func part2(caveMap CaveMap) int {
	return explore("start", make([]string, 0), true, caveMap)
}

func addMapEntry(caveMap CaveMap, start, end string) {
	if list, ok := caveMap[start]; !ok {
		caveMap[start] = []string{ end }
	} else {
		caveMap[start] = append(list, end)
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var caveMap CaveMap = make(CaveMap)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")
		start := parts[0]
		end := parts[1]
		
		addMapEntry(caveMap, start, end)
		addMapEntry(caveMap, end, start)		
	}

	fmt.Println(part1(caveMap))
	fmt.Println(part2(caveMap))
}

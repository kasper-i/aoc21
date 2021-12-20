package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Image [][]int

func (image Image) render() string {
	out := ""

	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[0]); x++ {
			if image[y][x] == 1 {
				out += "#"
			} else {
				out += "."
			}
		}

		out += fmt.Sprintln()
	}

	return out
}

func (image Image) get(y, x, fallback int) int {
	if x < 0 || y < 0 || x >= len(image[0]) || y >= len(image) {
		return fallback
	}

	return image[y][x]
}

func (image Image) enhance(enhancementAlgorithm string, fallbackPixel int) Image {
	padding := 1

	var enhancedImage Image = make(Image, len(image)+(padding*2))
	for y := range enhancedImage {
		enhancedImage[y] = make([]int, len(image[0])+(padding*2))
	}

	for y := -padding; y < len(image)+padding; y++ {
		for x := -padding; x < len(image[0])+padding; x++ {
			square := make([]int, 9)
			square[0] = image.get(y-1, x-1, fallbackPixel)
			square[1] = image.get(y-1, x, fallbackPixel)
			square[2] = image.get(y-1, x+1, fallbackPixel)
			square[3] = image.get(y, x-1, fallbackPixel)
			square[4] = image.get(y, x, fallbackPixel)
			square[5] = image.get(y, x+1, fallbackPixel)
			square[6] = image.get(y+1, x-1, fallbackPixel)
			square[7] = image.get(y+1, x, fallbackPixel)
			square[8] = image.get(y+1, x+1, fallbackPixel)

			var enhancementIndex int

			for idx, value := range square {
				if value == 1 {
					enhancementIndex += (1 << (8 - idx))
				}
			}

			var enhancedPixel int
			if enhancementAlgorithm[enhancementIndex] == '#' {
				enhancedPixel = 1
			}

			enhancedImage[y+padding][x+padding] = enhancedPixel
		}
	}

	return enhancedImage
}

func (image Image) countLitPixels() int {
	count := 0

	for _, scanLine := range image {
		for _, pixel := range scanLine {
			if pixel == 1 {
				count += 1
			}
		}
	}

	return count
}

func part1(image Image, enhancementAlgorithm string) int {
	return process(image, enhancementAlgorithm, 2)
}

func part2(image Image, enhancementAlgorithm string) int {
	return process(image, enhancementAlgorithm, 50)
}

func process(image Image, enhancementAlgorithm string, iterations int) int {
	for k := 0; k < iterations; k++ {
		fallback := 0
		if k%2 == 1 && enhancementAlgorithm[0] == '#' {
			fallback = 1
		}
		image = image.enhance(enhancementAlgorithm, fallback)
	}

	return image.countLitPixels()
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	enhancementAlgorithm := ""
	var image Image = make(Image, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if enhancementAlgorithm == "" {
			enhancementAlgorithm = line
		} else if line != "" {
			scanLine := make([]int, len(line))

			for i, r := range line {
				if r == '#' {
					scanLine[i] = 1
				}
			}

			image = append(image, scanLine)
		}
	}

	fmt.Println(part1(image, enhancementAlgorithm))
	fmt.Println(part2(image, enhancementAlgorithm))
}

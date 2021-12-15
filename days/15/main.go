package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

type Item struct {
	node  Coord
	risk  int
	from  Coord
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].risk < pq[j].risk
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, from Coord, risk int) {
	item.from = from
	item.risk = risk
	heap.Fix(pq, item.index)
}

type Coord struct {
	x int
	y int
}

type Matrix [][]int

func createMatrix(squareSize int) Matrix {
	var matrix Matrix = make(Matrix, squareSize)

	for k := 0; k < squareSize; k++ {
		matrix[k] = make([]int, squareSize)
	}

	return matrix
}

func (matrix Matrix) size() int {
	return len(matrix)
}

func findLowestCost(risks Matrix) int {
	dim := risks.size()

	visited := createMatrix(risks.size())
	var itemLookup map[Coord]*Item = make(map[Coord]*Item)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startNode := &Item{risk: 0}
	visited[0][0] = 1
	itemLookup[Coord{x: 0, y: 0}] = startNode

	current := startNode.node

	for {
		currentRisk := itemLookup[Coord{x: current.x, y: current.y}].risk

		var neighbours []Coord = []Coord{current, current, current, current}
		neighbours[0].y -= 1
		neighbours[1].y += 1
		neighbours[2].x -= 1
		neighbours[3].x += 1

		for _, neighbour := range neighbours {
			x, y := neighbour.x, neighbour.y

			if x < 0 || y < 0 || x > dim-1 || y > dim-1 {
				continue
			}

			if visited[y][x] != 0 {
				continue
			}

			newRisk := risks[y][x] + currentRisk

			if existing, ok := itemLookup[neighbour]; !ok {
				new := &Item{
					node: neighbour,
					risk: newRisk,
					from: current,
				}
				itemLookup[neighbour] = new
				heap.Push(&pq, new)
			} else if newRisk < existing.risk {
				pq.update(existing, current, newRisk)
			}
		}

		visited[current.y][current.x] = 1

		var next Coord
		if pq.Len() > 0 {
			next = heap.Pop(&pq).(*Item).node
		}

		if next.x == 0 && next.y == 0 {
			break
		} else {
			current = next
		}
	}

	renderImage(dim, itemLookup)

	return itemLookup[Coord{x: dim - 1, y: dim - 1}].risk
}

func renderImage(dim int, itemLookup map[Coord]*Item) {
	cursor := itemLookup[Coord{x: dim - 1, y: dim - 1}]

	width := dim
	height := dim

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			img.Set(x, y, color.Black)
		}
	}

	for {
		img.Set(cursor.node.x, cursor.node.y, color.White)

		if cursor.node.x == 0 && cursor.node.y == 0 {
			break
		}

		cursor = itemLookup[cursor.from]
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func timesFive(orig Matrix) Matrix {
	dim := len(orig)
	var x5 Matrix = make(Matrix, 5*dim)

	for y := 0; y < 5*dim; y++ {
		x5[y] = make([]int, 5*dim)
	}

	for y, row := range orig {
		for x := range row {
			for h := 0; h < 5; h++ {
				for v := 0; v < 5; v++ {
					prev := 0

					if h == 0 && v == 0 {
						x5[y][x] = orig[y][x]
						continue
					}

					if h == 0 {
						prev = x5[y+(v-1)*dim][x+h*dim]
					} else {
						prev = x5[y+v*dim][x+(h-1)*dim]
					}

					next := 0
					if prev == 9 {
						next = 1
					} else {
						next = prev + 1
					}

					x5[y+v*dim][x+h*dim] = next
				}
			}

		}
	}

	return x5
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix Matrix = make(Matrix, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var row []int = make([]int, 0)
		for _, r := range line {
			value, err := strconv.Atoi(string(r))

			if err != nil {
				log.Fatal(err)
			}

			row = append(row, value)
		}

		matrix = append(matrix, row)
	}

	fmt.Println(findLowestCost(matrix))
	fmt.Println(findLowestCost(timesFive(matrix)))
}

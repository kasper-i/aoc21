package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  int
}

func format(node *Node) string {
	if node.left == nil && node.right == nil {
		return fmt.Sprint(node.value)
	} else {
		return fmt.Sprintf("[%s,%s]", format(node.left), format(node.right))
	}
}

func deepCopy(node *Node) *Node {
	copy := &Node{}

	copy.value = node.value

	var leftCopy *Node
	var rightCopy *Node

	if node.left != nil {
		leftCopy = deepCopy(node.left)
		leftCopy.parent = copy
	}

	if node.right != nil {
		rightCopy = deepCopy(node.right)
		rightCopy.parent = copy
	}

	copy.left = leftCopy
	copy.right = rightCopy

	return copy
}

func add(n1, n2 *Node) *Node {
	sum := &Node{}

	sum.left = n1
	sum.right = n2

	n1.parent = sum
	n2.parent = sum

	return sum
}

func findLeft(node *Node) *Node {
	for {
		if node.parent == nil {
			return nil
		}

		prev := node
		node = node.parent

		if node.left != prev {
			node = node.left
			break
		}
	}

	for node.right != nil {
		node = node.right
	}

	return node
}

func findRight(node *Node) *Node {
	for {
		if node.parent == nil {
			return nil
		}

		prev := node
		node = node.parent

		if node.right != prev {
			node = node.right
			break
		}
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

func explode(node *Node) {
	leftValue := node.left.value
	rightValue := node.right.value

	var toLeft *Node = findLeft(node)
	var toRight *Node = findRight(node)

	if toLeft != nil {
		toLeft.value += leftValue
	}

	if toRight != nil {
		toRight.value += rightValue
	}

	node.value = 0
	node.left = nil
	node.right = nil
}

func split(node *Node) {
	value := node.value

	left := &Node{}
	right := &Node{}

	left.value = value / 2
	right.value = value/2 + value%2

	left.parent = node
	right.parent = node
	node.left = left
	node.right = right
}

func isRegularNumber(node *Node) bool {
	return node.left == nil && node.right == nil
}

func isPairOfRegularNumbers(node *Node) bool {
	return node.left != nil && node.right != nil && isRegularNumber(node.left) && isRegularNumber(node.right)
}

func reduce(node *Node) bool {
	if reduceWithExplode(node, 0) {
		return true
	}

	if reduceWithSplit(node, 0) {
		return true
	}

	return false
}

func reduceWithExplode(node *Node, depth int) bool {
	if isPairOfRegularNumbers(node) && depth >= 4 {
		explode(node)
		return true
	}

	if node.left != nil && reduceWithExplode(node.left, depth+1) {
		return true
	}

	if node.right != nil && reduceWithExplode(node.right, depth+1) {
		return true
	}

	return false
}

func reduceWithSplit(node *Node, depth int) bool {
	if isRegularNumber(node) {
		if node.value >= 10 {
			split(node)
			return true
		}

		return false
	}

	if reduceWithSplit(node.left, depth+1) {
		return true
	}

	if reduceWithSplit(node.right, depth+1) {
		return true
	}

	return false
}

func magnitude(node *Node) int {
	if isRegularNumber(node) {
		return node.value
	}

	return 3*magnitude(node.left) + 2*magnitude(node.right)
}

func part1(numbers []*Node) int {
	var sum *Node

	sum = add(deepCopy(numbers[0]), deepCopy(numbers[1]))
	for reduce(sum) {
	}

	for k := 2; k < len(numbers); k++ {
		sum = add(sum, deepCopy(numbers[k]))
		for reduce(sum) {
		}
	}

	return int(magnitude(sum))
}

func part2(numbers []*Node) int {
	greatestMagnitude := 0

	for _, n1 := range numbers {
		for _, n2 := range numbers {
			sum := add(deepCopy(n1), deepCopy(n2))

			for reduce(sum) {
			}

			if mag := magnitude(sum); mag > greatestMagnitude {
				greatestMagnitude = mag
			}
		}
	}

	return greatestMagnitude
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var roots []*Node = make([]*Node, 0)

	for scanner.Scan() {
		line := scanner.Text()

		root := &Node{}
		var node *Node = root

		for _, char := range line {
			switch char {
			case '[':
				left := &Node{}
				left.parent = node
				node.left = left
				node = left
			case ']':
				node = node.parent
			case ',':
				node = node.parent
				right := &Node{}
				right.parent = node
				node.right = right
				node = right
			default:
				value, _ := strconv.Atoi(string(char))
				node.value = value
			}
		}

		roots = append(roots, root)
	}

	fmt.Println(part1(roots))
	fmt.Println(part2(roots))
}

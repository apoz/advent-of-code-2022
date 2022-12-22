package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getMatrix(lines []string) [][]int {
	var matrix [][]int = make([][]int, 0)
	for _, line := range lines {
		var runesRow []string = strings.Split(line, "")
		var intRow []int = make([]int, 0)
		for _, r := range runesRow {
			rInt, _ := strconv.Atoi(r)
			intRow = append(intRow, rInt)
		}
		matrix = append(matrix, intRow)
	}
	return matrix
}

func visibleTop(x int, y int, matrix [][]int) bool {
	var treeHeight int = matrix[x][y]
	for i := x - 1; i >= 0; i-- {
		println("visibleTop x:", x, "y:", y, "i:", i)
		if matrix[i][y] >= treeHeight {
			return false
		}
	}
	println("Tree x:", x, " y:", y, "height:", treeHeight, "is visible from the top")
	return true
}

func visibleBottom(x int, y int, matrix [][]int) bool {
	var treeHeight int = matrix[x][y]
	var xLimit int = len(matrix)
	for i := x + 1; i < xLimit; i++ {
		if matrix[i][y] >= treeHeight {
			return false
		}
	}
	println("Tree x:", x, " y:", y, "height:", treeHeight, "is visible from the bottom")
	return true
}

func visibleLeft(x int, y int, matrix [][]int) bool {
	var treeHeight int = matrix[x][y]
	for i := y - 1; i >= 0; i-- {
		if matrix[x][i] >= treeHeight {
			return false
		}
	}
	println("Tree x:", x, " y:", y, "height:", treeHeight, "is visible from the left")
	return true
}

func visibleRight(x int, y int, matrix [][]int) bool {
	var treeHeight int = matrix[x][y]
	var yLimit int = len(matrix[0])
	for i := y + 1; i < yLimit; i++ {
		if matrix[x][i] >= treeHeight {
			return false
		}
	}
	println("Tree x:", x, " y:", y, "height:", treeHeight, "is visible from the right")
	return true
}

func sumVisible(matrix [][]int) int {
	var visibleTrees int = sumVisibleBorder(matrix)
	fmt.Println("Border visibleTrees ", visibleTrees)
	xLimit := len(matrix)
	yLimit := len(matrix[0])
	for x := 1; x < xLimit-1; x++ {
		for y := 1; y < yLimit-1; y++ {
			if visibleTop(x, y, matrix) ||
				visibleBottom(x, y, matrix) ||
				visibleLeft(x, y, matrix) ||
				visibleRight(x, y, matrix) {
				visibleTrees += 1
			}
		}
	}
	return visibleTrees
}

func sumVisibleBorder(matrix [][]int) int {
	var total int = len(matrix) * 2
	total += (len(matrix[0]) - 2) * 2
	return total
}

func main() {
	inputLines := readStdin()
	matrix := getMatrix(inputLines)
	total := sumVisible(matrix)
	fmt.Printf("Total %d\n", total)
}

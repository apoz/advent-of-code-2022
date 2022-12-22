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

func visibleTop(x int, y int, matrix [][]int) int {
	var treeHeight int = matrix[x][y]
	var count int = 0
	for i := x - 1; i >= 0; i-- {
		count++
		if matrix[i][y] >= treeHeight {
			return count
		}
	}
	return count
}

func visibleBottom(x int, y int, matrix [][]int) int {
	var treeHeight int = matrix[x][y]
	var xLimit int = len(matrix)
	var count int = 0
	for i := x + 1; i < xLimit; i++ {
		count++
		if matrix[i][y] >= treeHeight {
			return count
		}
	}
	return count
}

func visibleLeft(x int, y int, matrix [][]int) int {
	var treeHeight int = matrix[x][y]
	var count int = 0
	for i := y - 1; i >= 0; i-- {
		count++
		if matrix[x][i] >= treeHeight {
			return count
		}
	}
	return count
}

func visibleRight(x int, y int, matrix [][]int) int {
	var treeHeight int = matrix[x][y]
	var yLimit int = len(matrix[0])
	var count = 0
	for i := y + 1; i < yLimit; i++ {
		count++
		if matrix[x][i] >= treeHeight {
			return count
		}
	}
	return count
}

func maxMult(matrix [][]int) int {
	var maxMult int = 0

	xLimit := len(matrix)
	yLimit := len(matrix[0])
	for x := 1; x < xLimit-1; x++ {
		for y := 1; y < yLimit-1; y++ {
			currentMult := visibleTop(x, y, matrix) *
				visibleBottom(x, y, matrix) *
				visibleLeft(x, y, matrix) *
				visibleRight(x, y, matrix)
			if currentMult > maxMult {
				maxMult = currentMult
			}
		}
	}
	return maxMult
}

func main() {
	inputLines := readStdin()
	matrix := getMatrix(inputLines)
	maxTrees := maxMult(matrix)
	fmt.Printf("maxTrees %d\n", maxTrees)
}

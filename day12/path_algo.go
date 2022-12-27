package main

import (
	"bufio"
	"fmt"
	"os"
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

func getMatrixes(lines []string) ([][]string, [][]int) {
	var matrixStr [][]string = make([][]string, 0)
	var matrixInt [][]int = make([][]int, 0)
	for _, line := range lines {
		var stringsRow []string = strings.Split(line, "")
		var intRow []int = make([]int, 0)
		for _, c := range stringsRow {
			intRow = append(intRow, -1)
			_ = c
		}
		matrixStr = append(matrixStr, stringsRow)
		matrixInt = append(matrixInt, intRow)
	}
	return matrixStr, matrixInt
}

func getStartPos(matrix [][]string) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func getEndPos(matrix [][]string) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "E" {
				return i, j
			}
		}
	}
	return -1, -1
}

func printMatrix(strMatrix [][]string, intMatrix [][]int) {
	print("=============================\n")
	for i := 0; i < len(strMatrix); i++ {
		fmt.Printf("%+v\n", strMatrix[i])
	}
	for i := 0; i < len(strMatrix); i++ {
		fmt.Printf("%+v\n", intMatrix[i])
	}
	print("=============================\n")
}

func getMaxValue(matrix [][]int) int {
	var max int = -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}
	return max
}

func updateSurroundings(matrix *[][]int, x int, y int, stringMatrix [][]string) bool {
	couldMove := false
	// x + 1, y
	if canMove(stringMatrix, x, y, x+1, y) && (*matrix)[x+1][y] < 0 {
		(*matrix)[x+1][y] = (*matrix)[x][y] + 1
		couldMove = true
	}
	// x, y + 1
	if canMove(stringMatrix, x, y, x, y+1) && (*matrix)[x][y+1] < 0 {
		(*matrix)[x][y+1] = (*matrix)[x][y] + 1
		couldMove = true
	}
	// x - 1, y
	if canMove(stringMatrix, x, y, x-1, y) && (*matrix)[x-1][y] < 0 {
		(*matrix)[x-1][y] = (*matrix)[x][y] + 1
		couldMove = true
	}

	// x, y - 1
	if canMove(stringMatrix, x, y, x, y-1) && (*matrix)[x][y-1] < 0 {
		(*matrix)[x][y-1] = (*matrix)[x][y] + 1
		couldMove = true
	}
	return couldMove
}

func canMove(stringMatrix [][]string, origX int, origY int, destX int, destY int) bool {
	if destX < 0 || destY < 0 || destX >= len(stringMatrix) || destY >= len(stringMatrix[0]) {
		return false
	}
	origRuneSlice := []rune(stringMatrix[origX][origY])
	destRuneSlice := []rune(stringMatrix[destX][destY])
	origInt := int(origRuneSlice[0])
	destInt := int(destRuneSlice[0])
	fmt.Printf("OrigInt %d  (%s) DestInt %d (%s)\n", origInt, stringMatrix[origX][origY], destInt, stringMatrix[destX][destY])
	if destInt <= origInt+1 {
		fmt.Printf("Return TRUE\n")
		return true
	}
	fmt.Printf("Return FALSE\n")
	return false
}

func stepForward(matrix *[][]int, stringMatrix [][]string) bool {
	max := getMaxValue(*matrix)
	couldMove := false
	pointMove := false

	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len((*matrix)[i]); j++ {
			if (*matrix)[i][j] == max {
				pointMove = updateSurroundings(matrix, i, j, stringMatrix)
				if pointMove {
					couldMove = true
				}
			}
		}
	}
	return couldMove
}

func main() {
	inputLines := readStdin()
	stringMatrix, intMatrix := getMatrixes(inputLines)
	startX, startY := getStartPos(stringMatrix)
	endX, endY := getEndPos(stringMatrix)
	fmt.Printf("Destination %d, %d\n", endX, endY)

	intMatrix[startX][startY] = 0 // setting start to 0
	intMatrix[endX][endY] = -2
	stringMatrix[startX][startY] = "a"
	stringMatrix[endX][endY] = "z"
	counter := 0
	couldMove := false
	for {
		couldMove = stepForward(&intMatrix, stringMatrix)
		printMatrix(stringMatrix, intMatrix)
		if counter%100 == 0 {
			fmt.Println("Steps -> ", counter)
		}

		if intMatrix[endX][endY] != -2 {
			fmt.Printf("The result is %d\n", intMatrix[endX][endY])
			return
		}
		if couldMove == false {
			fmt.Println("COULD NOT MOVE WHEN COUNTER WAS ", counter)
			return
		}
		counter++
		couldMove = false
	}

}

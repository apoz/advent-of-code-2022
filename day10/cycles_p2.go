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

func runProgram(lines []string) int {
	var cycles []int = make([]int, 0)
	cycles = append(cycles, 1)
	for _, line := range lines {
		op, val := getOpFromLine(line)
		applyOps(&cycles, op, val)
	}
	printSprites(cycles)
	// fmt.Printf("The X values after program %+v\n", cycles)
	// return signalStrenght(cycles)
	return 0
}

func printSprites(cycles []int) {
	var res [][]string = make([][]string, 0)
	cycleIndex := 0
	for j := 0; j < 6; j++ {
		var line []string = make([]string, 0)
		for i := 0; i < 40; i++ {
			x := cycles[cycleIndex]
			if (i == x-1) || (i == x) || (i == x+1) {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}
			cycleIndex++
		}
		res = append(res, line)
		fmt.Println(line)
	}
}

func applyOps(cycles *[]int, ops string, val *int) {
	var lastX int = (*cycles)[len(*cycles)-1]
	if ops == "noop" {
		*cycles = append(*cycles, lastX)
	} else if ops == "addx" {
		*cycles = append(*cycles, lastX)
		*cycles = append(*cycles, lastX+*val)
	}

}

func signalStrenght(cycles []int) int {
	totalSum := 0
	index := 19
	for {
		strengthSignal := (index + 1) * cycles[index]
		fmt.Printf("StrengthSignal index %d -> %d  X->%d\n", index, strengthSignal, cycles[index])
		totalSum += strengthSignal
		index = index + 40
		if index > len(cycles) {
			return totalSum
		}

	}
}

func getOpFromLine(line string) (string, *int) {
	var fields []string = strings.Split(line, " ")
	var operation string = fields[0]
	var argument *int = nil
	if len(fields) > 1 {
		value, _ := strconv.Atoi(fields[1])
		argument = &value
	}
	return operation, argument
}

func main() {
	inputLines := readStdin()
	response := runProgram(inputLines)
	fmt.Println("Response ", response)

}

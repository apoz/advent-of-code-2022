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
	fmt.Printf("The X values after program %+v\n", cycles)
	return signalStrenght(cycles)
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

package main

import (
	"bufio"
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

type assignment struct {
	lower  int
	higher int
}

func getAssignments(line string) [2]assignment {
	var assignments [2]assignment
	assignList := strings.Split(line, ",")
	assignments[0] = getAssignment(assignList[0])
	assignments[1] = getAssignment(assignList[1])
	return assignments
}

func getAssignment(ass string) assignment {
	var assign assignment
	numbers := strings.Split(ass, "-")
	assign.lower, _ = strconv.Atoi(numbers[0])
	assign.higher, _ = strconv.Atoi(numbers[1])
	return assign

}

func isContained(ass1 assignment, ass2 assignment) bool {
	if (ass1.lower <= ass2.lower) &&
		(ass1.higher >= ass2.higher) {
		return true
	}
	if (ass2.lower <= ass1.lower) &&
		(ass2.higher >= ass1.higher) {
		return true
	}
	return false
}

func main() {
	totalCount := 0
	inputLines := readStdin()
	for _, line := range inputLines {
		assignments := getAssignments(line)
		if isContained(assignments[0], assignments[1]) {
			totalCount++
		}
	}
	println("Total count for overlappin assignments ", totalCount)
}

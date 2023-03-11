package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var snafu2Dec map[rune]int = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var dec2Snafu map[int]rune = map[int]rune{
	2:  '2',
	1:  '1',
	0:  '0',
	-1: '-',
	-2: '=',
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func addSnafus(inputLines []string) string {
	maxCols := getMaxLength(inputLines)
	inputNumbers := getLinesSplitted(inputLines)
	var snafuTotalDigits []rune = make([]rune, 0)

	var carryOver int = 0
	for col := 0; col < maxCols; col++ {
		columnSum := carryOver
		for _, number := range inputNumbers {
			if col < len(number) {
				columnSum += snafu2Dec[number[len(number)-1-col]]
			}
		}
		// we have the sum of the column, now we have to handle the carry overs
		carryOver = 0
		for columnSum > 2 {
			columnSum -= 5
			carryOver += 1
		}

		for columnSum < -2 {
			columnSum += 5
			carryOver -= 1

		}
		snafuTotalDigits = append(snafuTotalDigits, dec2Snafu[columnSum])

	}
	var rev []rune
	for i := len(snafuTotalDigits) - 1; i >= 0; i-- {
		rev = append(rev, snafuTotalDigits[i])
	}
	return string(rev)
}

func getLinesSplitted(inputLines []string) [][]rune {
	var inputNumbers [][]rune = make([][]rune, 0)
	for _, line := range inputLines {
		lineSlice := []rune(line)
		inputNumbers = append(inputNumbers, lineSlice)
	}
	return inputNumbers
}

func getMaxLength(inputLines []string) int {
	maxLength := 0
	for _, line := range inputLines {
		strings.ReplaceAll(line, "\n", "")
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}
func main() {
	inputLines := readStdin()
	fmt.Printf("The sum is %s\n", addSnafus(inputLines))

}

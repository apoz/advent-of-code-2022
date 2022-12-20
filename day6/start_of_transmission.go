package main

import (
	"bufio"
	"os"
)

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getRunesFromLine(line string) []rune {
	var runes []rune = []rune(line)
	return runes
}

func getFirstMessagePos(runes []rune) int {
	bottomIndex := 0
	headIndex := 4
	for {
		if allRunesDifferent(runes, bottomIndex, headIndex) {
			return headIndex
		}
		bottomIndex++
		headIndex++
		if headIndex == len(runes) {
			break
		}
	}

	return 0
}

func allRunesDifferent(runes []rune, bottomIndex int, headIndex int) bool {
	var aux map[rune]int = make(map[rune]int)
	for i := bottomIndex; i < headIndex; i++ {
		_, ok := aux[runes[i]]
		if ok {
			return false
		}
		aux[runes[i]] = 1
	}
	return true
}

func main() {
	inputLines := readStdin()
	for _, line := range inputLines {
		runes := getRunesFromLine(line)
		response := getFirstMessagePos(runes)
		println("First good rune is  ", response)
	}

}

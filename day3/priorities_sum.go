package main

import (
	"bufio"
	"os"
	"unicode"
)

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func splitRucksack(rucksack string) [2]map[rune]int {
	var first map[rune]int = make(map[rune]int)
	var second map[rune]int = make(map[rune]int)
	var m [2]map[rune]int
	string_length := len(rucksack)
	index_limit := string_length / 2
	for i, item := range rucksack {
		if i < index_limit {
			first[item] += 1
		} else {
			second[item] += 1
		}
	}
	m[0] = first
	m[1] = second
	return m
}

func getPrioPointsForElem(elem rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.

	if unicode.IsLower(elem) {
		return int(elem) - 96 // a is 97 in ASCII
	} else {
		return int(elem) - 38 // A is 65 in ASCII
	}
}

func getRepeatedRune(first map[rune]int, second map[rune]int) rune {
	for k, _ := range first {
		_, ok := second[k]
		if ok {
			// Rune present in both maps
			return k
		}
	}
	return rune(0)
}

func main() {
	totalPrioPoints := 0
	inputLines := readStdin()
	for _, rucksack := range inputLines {
		compartments := splitRucksack(rucksack)
		repeatedItem := getRepeatedRune(compartments[0], compartments[1])
		totalPrioPoints += getPrioPointsForElem(repeatedItem)
	}
	println("Total prio points is ", totalPrioPoints)
}

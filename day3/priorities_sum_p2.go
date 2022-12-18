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

func mapRucksack(rucksack string) map[rune]int {
	var mymap map[rune]int = make(map[rune]int)
	for _, item := range rucksack {
		_, ok := mymap[item]
		if ok {
			mymap[item]++
		} else {
			mymap[item] = 1
		}
	}
	return mymap
}

func getPrioPointsForElem(elem rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.

	if unicode.IsLower(elem) {
		return int(elem) - 96 // a is 97 in ASCII
	}
	return int(elem) - 38 // A is 65 in ASCII
}

func getRepeatedRune(first map[rune]int, second map[rune]int, third map[rune]int) rune {
	for k := range first {
		_, ok := second[k]
		if ok {
			_, ok2 := third[k]
			if ok2 {
				return k
			}
		}
	}
	return rune(0)
}

func main() {
	totalPrioPoints := 0
	inputLines := readStdin()
	for i := 0; i <= len(inputLines)-3; i += 3 {
		compElf1 := mapRucksack(inputLines[i])
		compElf2 := mapRucksack(inputLines[i+1])
		compElf3 := mapRucksack(inputLines[i+2])
		repeatedItem := getRepeatedRune(compElf1, compElf2, compElf3)
		totalPrioPoints += getPrioPointsForElem(repeatedItem)
	}
	println("Total prio points is ", totalPrioPoints)
}

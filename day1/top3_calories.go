package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func min(array [3]int) int {
	var min int = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func replace(array [3]int, newValue int, removeValue int) [3]int {
	for i, value := range array {
		if value == removeValue {
			array[i] = newValue
			break
		}
	}
	return array
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var maxCals [3]int = [3]int{0, 0, 0}
	var currentCals int = 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			minCals := min(maxCals)
			if currentCals > minCals {
				maxCals = replace(maxCals, currentCals, minCals)
			}
			currentCals = 0
		} else {
			cals, err := strconv.Atoi(line)
			if err != nil {
				log.Println("Unexpected content: " + line)
			}
			currentCals += cals
		}
	}
	log.Println("The max number of calories for 3 elfs is: " + strconv.Itoa(maxCals[0]+maxCals[1]+maxCals[2]))
}

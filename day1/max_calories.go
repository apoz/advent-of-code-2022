package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	maxCals := 0
	currentCals := 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			if currentCals > maxCals {
				maxCals = currentCals
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
	log.Println("The max number of calories for one elf is: " + strconv.Itoa(maxCals))
}

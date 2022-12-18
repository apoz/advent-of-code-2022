package main

import (
	"bufio"
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

func getPointsForChoice(choice string) int {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	switch choice {
	case "ROCK":
		return 1
	case "PAPER":
		return 2
	case "SCISSORS":
		return 3
	default:
		return 0
	}
}

func getPointsForRound(elfChoice string, playerChoice string) int {
	// (0 if you lost, 3 if the round was a draw, and 6 if you won).
	output := ""
	switch playerChoice {
	case "ROCK":
		switch elfChoice {
		case "ROCK":
			// tie
			output = "TIE"
		case "PAPER":
			// lost
			output = "LOST"
		case "SCISSORS":
			// win
			output = "WIN"
		}
	case "PAPER":
		switch elfChoice {
		case "ROCK":
			// win
			output = "WIN"
		case "PAPER":
			// tie
			output = "TIE"
		case "SCISSORS":
			// lost
			output = "LOST"
		}
	case "SCISSORS":
		switch elfChoice {
		case "ROCK":
			// lost
			output = "LOST"
		case "PAPER":
			// win
			output = "WIN"
		case "SCISSORS":
			// tie
			output = "TIE"
		}
	}
	switch output {
	case "WIN":
		return 6
	case "TIE":
		return 3
	case "LOST":
		return 0
	default:
		return 0
	}
}

func getRPSValue(line string) string {
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	switch line {
	case "A", "X":
		return "ROCK"
	case "B", "Y":
		return "PAPER"
	case "C", "Z":
		return "SCISSORS"
	default:
		return "UNKNOWN"
	}
}

func main() {
	totalPoints := 0
	inputLines := readStdin()
	for _, line := range inputLines {
		choices := strings.Split(line, " ")
		elfChoice := getRPSValue(choices[0])
		playerChoice := getRPSValue(choices[1])
		roundPoints := getPointsForChoice(playerChoice) +
			getPointsForRound(elfChoice, playerChoice)
		totalPoints += roundPoints
	}
	println("Total points for the player is ", totalPoints)
}

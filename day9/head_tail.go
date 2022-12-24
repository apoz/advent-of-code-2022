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

type pos struct {
	x int
	y int
}

func trackMovements(lines []string) int {
	var headMovs []pos = make([]pos, 0)
	var tailMovs []pos = make([]pos, 0)
	initialPos := pos{
		x: 0,
		y: 0,
	}
	headMovs = append(headMovs, initialPos)
	tailMovs = append(tailMovs, initialPos)
	for _, line := range lines {
		direction, amount := getMovementFromLine(line)
		println("Direction ", direction, "Amount ", amount)
		for i := 0; i < amount; i++ {
			headMovs = headMovement(headMovs, direction)
			tailMovs = tailMovement(tailMovs, headMovs[len(headMovs)-1])
			fmt.Printf(" Head Pos %+v  Tail Pos %+v\n", headMovs[len(headMovs)-1], tailMovs[len(tailMovs)-1])
		}
	}
	fmt.Printf("headMovs -> %+v\n", headMovs)
	fmt.Printf("tailMovs -> %+v\n", tailMovs)
	return countDifferentPos(tailMovs)
}

func countDifferentPos(posList []pos) int {
	var differentPos map[pos]bool = make(map[pos]bool)
	for _, mypos := range posList {
		differentPos[mypos] = true
	}
	return len(differentPos)
}

func getMovementFromLine(line string) (string, int) {
	var direction string
	var amount int
	var fields []string = strings.Split(line, " ")
	direction = fields[0]
	amount, _ = strconv.Atoi(fields[1])
	return direction, amount
}

func headMovement(track []pos, direction string) []pos {
	lastPos := track[len(track)-1]
	var newPos pos
	if direction == "R" {
		newPos = pos{
			x: lastPos.x + 1,
			y: lastPos.y,
		}
	} else if direction == "L" {
		newPos = pos{
			x: lastPos.x - 1,
			y: lastPos.y,
		}
	} else if direction == "U" {
		newPos = pos{
			x: lastPos.x,
			y: lastPos.y + 1,
		}
	} else if direction == "D" {
		newPos = pos{
			x: lastPos.x,
			y: lastPos.y - 1,
		}
	}
	return append(track, newPos)
}

func tailMovement(track []pos, headPos pos) []pos {
	if len(track) == 1 && (headPos.x == 1 || headPos.y == 1) { // initial pos is initialized already
		return track
	}

	tailPos := track[len(track)-1]
	var newTailPos *pos = nil
	if tailPos.x == headPos.x && tailPos.y == headPos.y { //head and tail in the same pos
		return track
	}
	if tailPos.x+2 == headPos.x && tailPos.y == headPos.y { // head 2 X's up
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y,
		}
	} else if tailPos.x-2 == headPos.x && tailPos.y == headPos.y { // head 2 X's down
		newTailPos = &pos{
			x: tailPos.x - 1,
			y: tailPos.y,
		}
	} else if tailPos.x == headPos.x && tailPos.y+2 == headPos.y { // head 2 Y's up
		newTailPos = &pos{
			x: tailPos.x,
			y: tailPos.y + 1,
		}
	} else if tailPos.x == headPos.x && tailPos.y-2 == headPos.y { // head 2 Y's down
		newTailPos = &pos{
			x: tailPos.x,
			y: tailPos.y - 1,
		}
	} else if (tailPos.x+2 == headPos.x && tailPos.y+1 == headPos.y) ||
		(tailPos.x+1 == headPos.x && tailPos.y+2 == headPos.y) { // top right diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y + 1,
		}
	} else if (tailPos.x+2 == headPos.x && tailPos.y-1 == headPos.y) ||
		(tailPos.x+1 == headPos.x && tailPos.y-2 == headPos.y) { // top left diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y - 1,
		}
	} else if (tailPos.x-2 == headPos.x && tailPos.y+1 == headPos.y) ||
		(tailPos.x-1 == headPos.x && tailPos.y+2 == headPos.y) { // down right diagonal
		newTailPos = &pos{
			x: tailPos.x - 1,
			y: tailPos.y + 1,
		}
	} else if (tailPos.x-2 == headPos.x && tailPos.y-1 == headPos.y) ||
		(tailPos.x-1 == headPos.x && tailPos.y-2 == headPos.y) { // down left diagonal
		newTailPos = &pos{
			x: tailPos.x - 1,
			y: tailPos.y - 1,
		}
	} else if (tailPos.x+1 == headPos.x && tailPos.y+2 == headPos.y) ||
		(tailPos.x+2 == headPos.x && tailPos.y+1 == headPos.y) { // down left diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y + 1,
		}
	} else {
		if (tailPos.x+1 == headPos.x) ||
			(tailPos.x-1 == headPos.x) ||
			(tailPos.y-1 == headPos.y) ||
			(tailPos.y+1 == headPos.y) {
			fmt.Printf(" Single separate diagonal, no worries")
		} else {
			fmt.Printf("SHOULD NOT HAPPEN headPos %+v tailPos %+v\n", headPos, track)
		}
	}
	if newTailPos != nil {
		return append(track, *newTailPos)
	}
	return track
}

func main() {
	inputLines := readStdin()
	result := trackMovements(inputLines)
	fmt.Printf("Result -> %d\n", result)
}

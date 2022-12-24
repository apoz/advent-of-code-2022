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
	var Movs [][]pos = make([][]pos, 0)
	for i := 0; i < 10; i++ {
		initialPos := pos{
			x: 0,
			y: 0,
		}
		var myMovs []pos = make([]pos, 0)
		myMovs = append(myMovs, initialPos)
		Movs = append(Movs, myMovs)
	}
	// var headMovs []pos = make([]pos, 0)
	// var Movs []pos = make([]pos, 0)
	// initialPos := pos{
	// 	x: 0,
	// 	y: 0,
	// }
	// headMovs = append(headMovs, initialPos)
	// tailMovs = append(tailMovs, initialPos)
	for _, line := range lines {
		direction, amount := getMovementFromLine(line)
		println("Direction ", direction, "Amount ", amount)
		for i := 0; i < amount; i++ {
			Movs[0] = headMovement(Movs[0], direction)
			for i := 1; i < len(Movs); i++ {
				Movs[i] = tailMovement(Movs[i], Movs[i-1][len(Movs[i-1])-1])
			}
			fmt.Println("===========================")
			for j := 0; j < len(Movs); j++ {
				fmt.Printf("Last pos in Mov[%d] is %+v\n", j, Movs[j][len(Movs[j])-1])
			}
			fmt.Println("===========================")
		}

	}
	return countDifferentPos(Movs[9])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tailMovement(track []pos, headPos pos) []pos {
	if len(track) == 1 && (Abs(headPos.x) < 2 && (Abs(headPos.y) < 2)) { // initial pos is initialized already
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
		(tailPos.x+1 == headPos.x && tailPos.y+2 == headPos.y) ||
		(tailPos.x+2 == headPos.x && tailPos.y+2 == headPos.y) { // top right diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y + 1,
		}
	} else if (tailPos.x+2 == headPos.x && tailPos.y-1 == headPos.y) ||
		(tailPos.x+1 == headPos.x && tailPos.y-2 == headPos.y) ||
		(tailPos.x+2 == headPos.x && tailPos.y-2 == headPos.y) { // top left diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y - 1,
		}
	} else if (tailPos.x-2 == headPos.x && tailPos.y+1 == headPos.y) ||
		(tailPos.x-1 == headPos.x && tailPos.y+2 == headPos.y) ||
		(tailPos.x-2 == headPos.x && tailPos.y+2 == headPos.y) { // down right diagonal
		newTailPos = &pos{
			x: tailPos.x - 1,
			y: tailPos.y + 1,
		}
	} else if (tailPos.x-2 == headPos.x && tailPos.y-1 == headPos.y) ||
		(tailPos.x-1 == headPos.x && tailPos.y-2 == headPos.y) ||
		(tailPos.x-2 == headPos.x && tailPos.y-2 == headPos.y) { // down left diagonal
		newTailPos = &pos{
			x: tailPos.x - 1,
			y: tailPos.y - 1,
		}
	} else if (tailPos.x+1 == headPos.x && tailPos.y+2 == headPos.y) ||
		(tailPos.x+2 == headPos.x && tailPos.y+1 == headPos.y) ||
		(tailPos.x+2 == headPos.x && tailPos.y+2 == headPos.y) { // down left diagonal
		newTailPos = &pos{
			x: tailPos.x + 1,
			y: tailPos.y + 1,
		}
	} else {
		if (tailPos.x+1 == headPos.x) ||
			(tailPos.x-1 == headPos.x) ||
			(tailPos.y-1 == headPos.y) ||
			(tailPos.y+1 == headPos.y) {
			fmt.Printf(" Single separate diagonal, no worries\n")
		} else {
			fmt.Printf("SHOULD NOT HAPPEN headPos %+v prevPos %+v\n", headPos, track[len(track)-1])
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

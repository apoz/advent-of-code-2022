package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Blizzard struct {
	depth     int
	width     int
	direction string
}

type Boards struct {
	boards []Board
	modulo int
}

func (b Boards) printBoards() {
	for i, board := range b.boards {
		fmt.Printf("Time %d\n", i)
		board.printBoard()
	}
}

func (b Boards) getBoardForTime(minute int) Board {
	return b.boards[(minute)%b.modulo]
}

type Board struct {
	board    [][]string
	maxDepth int
	maxWidth int
}

func (b Board) printBoard() {
	for _, line := range b.board {
		for _, elem := range line {
			fmt.Printf(" %s", elem)
		}
		fmt.Printf("\n")
	}
}

func getNextBlizzards(blizzards []Blizzard, maxDepth int, maxWidth int) []Blizzard {
	newBlizzards := make([]Blizzard, 0)
	for _, blizzard := range blizzards {
		newBlizzards = append(newBlizzards, getNextMinute(blizzard, maxDepth, maxWidth))
	}
	return newBlizzards
}

func getNextMinute(blizzard Blizzard, maxDepth int, maxWidth int) Blizzard {
	depth := blizzard.depth
	width := blizzard.width
	direction := blizzard.direction
	if direction == "^" {
		depth--
		if depth == 0 {
			depth = maxDepth - 1
		}
	} else if direction == ">" {
		width++
		if width == maxWidth {
			width = 1
		}
	} else if direction == "<" {
		width--
		if width == 0 {
			width = maxWidth - 1
		}
	} else if direction == "v" {
		depth++
		if depth == maxDepth {
			depth = 1
		}
	}
	return Blizzard{
		depth:     depth,
		width:     width,
		direction: direction,
	}
}

func getBoard(blizzards []Blizzard, depth int, width int) Board {
	board := getEmptyBoard(depth, width)
	for _, blizzard := range blizzards {
		counter, direction := getNumberOfBlizzards(blizzard.depth, blizzard.width, blizzards)
		if counter == 1 {
			board[blizzard.depth][blizzard.width] = direction
		} else {
			board[blizzard.depth][blizzard.width] = strconv.Itoa(counter)
		}
	}
	return Board{
		board:    board,
		maxDepth: depth - 1,
		maxWidth: width - 1,
	}
}

func getNumberOfBlizzards(depth int, width int, blizzards []Blizzard) (int, string) {
	counter := 0
	direction := ""
	for _, blizzard := range blizzards {
		if blizzard.depth == depth && blizzard.width == width {
			counter++
			direction = blizzard.direction
		}
	}
	return counter, direction
}

func getEmptyBoard(depth int, width int) [][]string {
	board := make([][]string, 0)
	for i := 0; i < depth; i++ {
		line := make([]string, width)
		for j := 0; j < width; j++ {
			if i == 0 || i == (depth-1) || j == 0 || j == (width-1) {
				line[j] = "#"
			} else {
				line[j] = "."
			}
		}
		board = append(board, line)
	}
	board[0][1] = "."
	board[depth-1][width-2] = "."
	return board
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getBlizzards(lines []string) []Blizzard {
	var blizzards []Blizzard = make([]Blizzard, 0)
	for depth, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		elems := strings.Split(line, "")
		for width, elem := range elems {
			if isBlizzard(elem) {
				blizzard := Blizzard{
					depth:     depth,
					width:     width,
					direction: elem,
				}
				blizzards = append(blizzards, blizzard)
			}
		}
	}
	return blizzards
}

func getDimensions(lines []string) (int, int) {
	depth := len(lines)
	width := len(strings.Split(lines[0], ""))
	return depth, width
}

func isBlizzard(elem string) bool {
	if elem == "^" ||
		elem == "<" ||
		elem == ">" ||
		elem == "v" {
		return true
	}
	return false
}

func getBoards(blizzards []Blizzard, depth int, width int) Boards {
	boards := make([]Board, 0)
	modulo := (depth - 2) * (width - 2)
	for i := 0; i < modulo; i++ {
		b := getBoard(blizzards, depth, width)
		boards = append(boards, b)
		blizzards = getNextBlizzards(blizzards, depth-1, width-1)
	}
	return Boards{
		boards: boards,
		modulo: modulo,
	}
}

type Status struct {
	time         int
	currentDepth int
	currentWidth int
	goalDepth    int
	goalWidth    int
}

func (s Status) printBoard(board Board) {
	for i, line := range board.board {
		for j, elem := range line {
			if i == s.currentDepth && j == s.currentWidth {
				fmt.Printf(" E")
			} else {
				fmt.Printf(" %s", elem)
			}
		}
		fmt.Printf("\n")
	}

}

func initialStatus(startDepth int, startWidth int, goalDepth int, goalWidth int, t int) Status {
	return Status{
		time:         t,
		currentDepth: startDepth,
		currentWidth: startWidth,
		goalDepth:    goalDepth,
		goalWidth:    goalWidth,
	}
}

func goalInFrontier(frontier map[Status]bool, goalDepth int, goalWidth int) bool {
	for k, _ := range frontier {
		if k.currentDepth == goalDepth && k.currentWidth == goalWidth {
			return true
		}
	}
	return false

}

func BFS(boards Boards, startDepth int, startWidth int, goalDepth int, goalWidth int, startTime int) (int, int) {
	frontier := make(map[Status]bool)
	initialStatus := initialStatus(startDepth, startWidth, goalDepth, goalWidth, startTime)
	fmt.Printf("Starting BFS with Status %+v\n", initialStatus)
	fmt.Printf("Starting Board:\n")
	initialStatus.printBoard(boards.boards[startTime%boards.modulo])
	frontier[initialStatus] = true
	boardTime := startTime + 1 //We're interested in next minute board
	spentTime := 0
	for goalInFrontier(frontier, goalDepth, goalWidth) == false {
		frontier = getNextFrontier(frontier, boards.getBoardForTime(boardTime%boards.modulo))
		spentTime++
		boardTime++
	}
	boardTime-- //it was increased in advance
	for status, _ := range frontier {
		if status.currentDepth == goalDepth && status.currentWidth == goalWidth {
			fmt.Printf("Found solution spentTime %d boardTime %d\n", spentTime, boardTime)
			status.printBoard(boards.boards[boardTime%boards.modulo])
			break
		}
	}
	return spentTime, boardTime
}

func getNextFrontier(frontier map[Status]bool, board Board) map[Status]bool {
	nextFrontier := make(map[Status]bool)
	for status, _ := range frontier {
		options := getOptions(status.currentDepth, status.currentWidth, board)
		for _, point := range options {
			// if point.depth == status.goalDepth && point.width == status.goalWidth {
			// 	fmt.Printf("Found solution, %d\n", status.time+1)
			// 	os.Exit(0)
			// }
			newStatus := Status{
				currentDepth: point.depth,
				currentWidth: point.width,
				goalDepth:    status.goalDepth,
				goalWidth:    status.goalWidth,
				time:         status.time + 1,
			}
			nextFrontier[newStatus] = true
		}
	}
	//fmt.Printf("New frontier is %d\n", len(frontier))
	return nextFrontier
}

func getStatusesForNextMinute(status Status, boards Boards) []Status {
	nextMinuteBoard := boards.getBoardForTime(status.time + 1)
	//fmt.Printf("\nNext minute (%d) board is:\n", status.time+1)
	//nextMinuteBoard.printBoard()
	currentDepth := status.currentDepth
	currentWidth := status.currentWidth
	options := getOptions(currentDepth, currentWidth, nextMinuteBoard)
	statuses := make([]Status, 0)
	for _, point := range options {
		//fmt.Printf("\nNext option %+v\n", point)
		// fmt.Printf("GOAL DEPTH: %d GOAL WIDTH %d\n", status.goalDepth, status.goalWidth)
		if point.depth == status.goalDepth && point.width == status.goalWidth {
			fmt.Printf("Found solution, %d\n", status.time+1)
			os.Exit(0)
		}
		newStatus := Status{
			currentDepth: point.depth,
			currentWidth: point.width,
			goalDepth:    status.goalDepth,
			goalWidth:    status.goalWidth,
			time:         status.time + 1,
		}
		//fmt.Printf("Next STEP BOARD\n")
		//newStatus.printBoard(nextMinuteBoard)
		statuses = append(statuses, newStatus)

	}
	return statuses

}

type Point struct {
	depth int
	width int
}

func getOptions(depth int, width int, board Board) []Point {
	points := make([]Point, 0)
	// move up
	if depth >= 1 && board.board[depth-1][width] == "." {
		points = append(points, Point{depth: depth - 1, width: width})
	}
	// move down
	if depth < len(board.board)-1 && board.board[depth+1][width] == "." {
		points = append(points, Point{depth: depth + 1, width: width})
	}
	// move right
	if board.board[depth][width+1] == "." {
		points = append(points, Point{depth: depth, width: width + 1})
	}
	// move left
	if board.board[depth][width-1] == "." {
		points = append(points, Point{depth: depth, width: width - 1})
	}
	// stand still
	if board.board[depth][width] == "." {
		points = append(points, Point{depth: depth, width: width})
	}
	return points
}

func main() {
	inputLines := readStdin()
	blizzards := getBlizzards(inputLines)
	depth, width := getDimensions(inputLines)
	fmt.Printf("depth %d width %d\n", depth, width)
	boards := getBoards(blizzards, depth, width)
	fmt.Printf("After boards \n")
	spentTime, boardTime := BFS(boards, 0, 1, depth-1, width-2, 0)
	fmt.Printf("First got solved spentTime %d BoardTime %d\n", spentTime, boardTime)
	spentTime2, boardTime := BFS(boards, depth-1, width-2, 0, 1, boardTime)
	fmt.Printf("Second got solved spentTime %d BoardTime %d\n", spentTime2, boardTime)
	spentTime3, boardTime := BFS(boards, 0, 1, depth-1, width-2, boardTime)
	fmt.Printf("Third got solved spentTime %d BoardTime %d\n", spentTime3, boardTime)
	fmt.Printf("Final solution %d\n", spentTime+spentTime2+spentTime3)
}

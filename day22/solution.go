package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	board     [][]string
	posDepth  int
	posWidth  int
	direction string
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getBoardDimensions(lines []string) (int, int) {
	width := 0
	depth := 0
	for count, line := range lines {
		w := len(strings.Split(line, ""))
		if w > width {
			width = w
		}
		depth = count
	}
	return depth + 1, width
}

func getBoard(lines []string) [][]string {
	_, width := getBoardDimensions(lines)
	board := make([][]string, 0)
	for _, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		elems := strings.Split(line, "")
		if len(elems) < width {
			for i := len(elems); i < width; i++ {
				elems = append(elems, " ")
			}
		}
		board = append(board, elems)
	}
	return board
}

func (p *Problem) printBoard(board [][]string) {
	for count, line := range (*p).board {
		fmt.Printf(" [%d]  ", count)
		for i, c := range line {
			if p.posDepth == count && p.posWidth == i {
				fmt.Print(" X")
			} else {
				fmt.Printf(" %s", c)
			}

		}
		fmt.Printf("\n")
	}
}

func (p *Problem) getStartPosition() (int, int) {
	var width int = 0
	for w, c := range (*p).board[0] {
		if c == "." {
			width = w
			break
		}
	}
	return 0, width
}

func (p *Problem) getFirstPositionAndValueInRow(depth int) (int, string) {
	i := 0
	for {
		if (*p).board[depth][i] == " " {
			i++
			continue
		} else {
			return i, (*p).board[depth][i]
		}
	}
}

func (p *Problem) getLastPositionAndValueInRow(depth int) (int, string) {
	i := len((*p).board[depth]) - 1
	for {
		if (*p).board[depth][i] == " " {
			i--
			continue
		} else {
			return i, (*p).board[depth][i]
		}

	}
}

func (p *Problem) getFirstPositionAndValueInColumn(width int) (int, string) {
	i := 0
	for {
		if (*p).board[i][width] == " " {
			i++
			continue
		} else {
			return i, (*p).board[i][width]
		}
	}
}

func (p *Problem) getLastPositionAndValueInColumn(width int) (int, string) {
	i := len((*p).board) - 1
	for {
		if (*p).board[i][width] == " " {
			i--
			continue
		} else {
			return i, (*p).board[i][width]
		}
	}
}

func (p *Problem) updateDirection(dir string) {
	if (*p).direction == "R" && dir == "R" {
		(*p).direction = "B"
	} else if (*p).direction == "B" && dir == "R" {
		(*p).direction = "L"
	} else if (*p).direction == "L" && dir == "R" {
		(*p).direction = "T"
	} else if (*p).direction == "T" && dir == "R" {
		(*p).direction = "R"
	} else if (*p).direction == "R" && dir == "L" {
		(*p).direction = "T"
	} else if (*p).direction == "T" && dir == "L" {
		(*p).direction = "L"
	} else if (*p).direction == "L" && dir == "L" {
		(*p).direction = "B"
	} else if (*p).direction == "B" && dir == "L" {
		(*p).direction = "R"
	} else {
		fmt.Println("Unexpected direction move!")
	}
}

func (p *Problem) makeMoveRight(steps int) {
	for i := 0; i < steps; i++ {
		if (*p).posWidth == len((*p).board[(*p).posDepth])-1 {
			if (*p).board[(*p).posDepth][(*p).posWidth] == "." {
				pos, val := p.getFirstPositionAndValueInRow(p.posDepth)
				if val == "." {
					p.posWidth = pos
				} else {
					break
				}
			}
		} else if (*p).board[(*p).posDepth][(*p).posWidth+1] == "." {
			(*p).posWidth = (*p).posWidth + 1
		} else if (*p).board[(*p).posDepth][(*p).posWidth+1] == "#" {
			break
		} else if (*p).board[(*p).posDepth][(*p).posWidth+1] == " " {
			pos, val := p.getFirstPositionAndValueInRow(p.posDepth)
			if val == "." {
				p.posWidth = pos
			} else {
				break
			}
		}
	}
}

func (p *Problem) makeMoveLeft(steps int) {
	for i := 0; i < steps; i++ {
		if (*p).posWidth == 0 {
			if (*p).board[(*p).posDepth][(*p).posWidth] == "." {
				pos, val := p.getLastPositionAndValueInRow(p.posDepth)
				if val == "." {
					p.posWidth = pos
				} else {
					break
				}
			}
		} else if (*p).board[(*p).posDepth][(*p).posWidth-1] == "." {
			(*p).posWidth = (*p).posWidth - 1
		} else if (*p).board[(*p).posDepth][(*p).posWidth-1] == "#" {
			break
		} else if (*p).board[(*p).posDepth][(*p).posWidth-1] == " " {
			pos, val := p.getLastPositionAndValueInRow(p.posDepth)
			if val == "." {
				p.posWidth = pos
			} else {
				break
			}
		}
	}
}

func (p *Problem) makeMoveTop(steps int) {
	for i := 0; i < steps; i++ {
		if (*p).posDepth == 0 {
			if (*p).board[(*p).posDepth][(*p).posWidth] == "." {
				pos, val := p.getLastPositionAndValueInColumn(p.posWidth)
				if val == "." {
					p.posDepth = pos
				} else {
					break
				}
			}
		} else if (*p).board[(*p).posDepth-1][(*p).posWidth] == "." {
			(*p).posDepth = (*p).posDepth - 1
		} else if (*p).board[(*p).posDepth-1][(*p).posWidth] == "#" {
			break
		} else if (*p).board[(*p).posDepth-1][(*p).posWidth] == " " {
			pos, val := p.getLastPositionAndValueInColumn(p.posWidth)
			if val == "." {
				p.posDepth = pos
			} else {
				break
			}
		}
	}
}

func (p *Problem) makeMoveBottom(steps int) {
	for i := 0; i < steps; i++ {
		if (*p).posDepth == len((*p).board)-1 {
			if (*p).board[(*p).posDepth][(*p).posWidth] == "." {
				pos, val := p.getFirstPositionAndValueInColumn(p.posWidth)
				if val == "." {
					p.posDepth = pos
				} else {
					break
				}
			}
		} else if (*p).board[(*p).posDepth+1][(*p).posWidth] == "." {
			(*p).posDepth = (*p).posDepth + 1
		} else if (*p).board[(*p).posDepth+1][(*p).posWidth] == "#" {
			break
		} else if (*p).board[(*p).posDepth+1][(*p).posWidth] == " " {
			pos, val := p.getFirstPositionAndValueInColumn(p.posWidth)
			if val == "." {
				p.posDepth = pos
			} else {
				break
			}
		}
	}
}

func (p *Problem) printResult() {
	dirs := map[string]int{
		"R": 0,
		"B": 1,
		"L": 2,
		"T": 3,
	}
	response := 1000*((*p).posDepth+1) + 4*((*p).posWidth+1) + dirs[(*p).direction]
	fmt.Printf("Response is %d\n", response)
}

func (p *Problem) makeMove(steps int) {
	if p.direction == "R" {
		p.makeMoveRight(steps)
	} else if p.direction == "L" {
		p.makeMoveLeft(steps)
	} else if p.direction == "T" {
		p.makeMoveTop(steps)
	} else if p.direction == "B" {
		p.makeMoveBottom(steps)
	}
}

func getMovs(line string) ([]int, []string) {
	reg, _ := regexp.Compile(`\d+`)
	reg2, _ := regexp.Compile(`\D+`)
	numbers := reg.FindAllString(line, -1)
	changes := reg2.FindAllString(line, -1)
	integers := make([]int, 0)
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		integers = append(integers, n)
	}
	return integers, changes
}

func main() {

	// var valves map[string]Valve = make(map[string]Valve)
	// var movements string = "10R5L5R10L4R5L5"
	var movements string = "5L37R32R38R28L18R11R37R41R42L8L3R14L22L28R4L32R44R3R45L44L39L27L25L7R32L36R21L29R33R37L37R19L44R36L14R45R13R32L9R16L23L19R1L18R9R48L38L49L9L43R30L13L20R46L46R38R34R9L16R1R8L36L14R9L17R18L19L23R42R43L3R27R9R25R39L40R47L1R17L34R50R43L16L10R28R34R47R33R18R41R44L26L3L40L2L50L7L28R8L21R36R45R16R41L33R34L37R31L40R36R4R5R47R11R6L25L6L9R22R42L4L41R39R8R50R10L38R45R7R36L43R20L18R45L40R30L10L15L44R24R16R36L24L32R47R43R12R31R26R4L38R37L16L7L42L28L42R19R29R4L11L40L8R21L31R5R26R33L44R34R3L6R9R43L36R29L30R32L29R35R28L27R10L28R6R23L29L26R32R28R31L46R30R36L29L7R39R5L11L29R3L3R29R9R11L50R15R6R35R50L18L49R26R24L33L48L30L10R44R1R31L35L35L2L27L6R24R10L34R14R9R26R3L46L7R11R37R1R43R35L48R36L1R13L21L8L31R43R40L1R49L12R33R32R19R38L12R12L47L5L9L29L18R7R10L24L42R47R44R3L22L23R5L2L17R44R30L7L17R26L41L17L37R1L46R26R9R40L50L6R3R48L40L20L34R47R46R19R32L35R22L7R34R8L4R49R47L32R26R12R11R47R4L14R1L24L30L21L5R10L28R5L31L46R33R42R27L31L14L35R41R18L43R45R31R8R32R39L45L14L38L49L34R26L49R37L27L47R7L14L22R22R42L20R36R21R35R10L37L45R35L14R48R10R21L30R27L34R1L46L34R47R22L7L46R41R33L41L39R50L29L50R33L46L36R7R48L1L3R35R24L25R43L7R50L13L41R22R49L23L6R49L12R2R34L29L11L36R8R21L10R50L10L49L42R31L18R25L30R13R40L20L35L17R47L13L27L42L42R5L16L22R30L4L28L18R23L19L16L40L35L6R37L30R14L31L50R1R19L9L10L50R40R23R34R50L8L30L28R10R19L26L13L44R28L19R11L19L30L7R10R46L17R15R27R36R19L19R49R11L37R1R16L36L48L31R41R30R16L18R7R35L34L43L38L34L11L48R17L40R8L31L47R4L36L28R48R34R29L38R21L22L23R3R34R17L32L42L34R29L46L19R32R48L41R46R9L40L49R41R36R23L49L30R1L28L10L14R35R27R7R45L49L35R42R31R17L5L15R21L22R5L16R2L3L19R24L26R31R10L32R10R17R17L20L42R31R11R11R25R22L39R19L4L46L34R44R38L13R40R23L2R16L21R8L7R30L2R45R21L32L42L2L10L45L46R21R31L22L10R42L7R20L12L39R37R37L43L45L1L8L14R44R19R23L37L6L12R27R26R8R12L3R47R35L25R41L39R46R10L32L35L31R12L20R8L19L26L18L5L22L15R24L2R1R38R50L4L45L4L16R16R31R41R5L29R49R25L21R20L15R33R47L7R32L35R5L31L18R42L12R9L21L4L36L14R33L7R21L27R18L20L47L14L37R35L42L22L11R33R13L11R11L7R7R23L47R4R24L10R45L16R45L35R27R43L47R3L42R14R2L44R1L45L25R23R1L47R38L46L18R46R42L1L34R48R42L42L50R39R25R43R19L48R38L20L45L40R28L39R4R16L46R5L41L25L4R50L11R6R34L19L33L33L27L24L16L23L15R22R32R19R7R4R30L1R5R23L46R38L23L42R40L21R45L27L1L41R44R26L47R34R44L37L38L37L38L1R50R30R38R28R5R13L31L11L30R34L16R8R27L19L10R33R49L36R43L5L19L1R24L29R13R19R24R10L48R36R18L40L28R13L36L1R18R27L42R11R16L26L20L46R27R29R26R40R5L31L22L33R37R33R28R22L42R5R4L19R34R26L13R8L46L25R40R39R36L6R9R22R42L25R38L18L10L11R15R25L17L9R36R47L31L5L33R6R2R45R33R26L16R9L21L9L18L4R30L36L35R6R19L42L38L5R9L39L13R3L20R16L48L47R16L15L13R43L17R30R27R16R32R7L11L50R50R20L3L24L44R42L50L22L42R16R37L37R8L41R50R47R5L7L2L38R39R9R36L24L23R21L49L42R18R37L49L37L38L2R48R41L4R28R1R36L6R23R22R22L13L18L35L11R27L16L19R7L5R22L26L23L43R17L1R36L34R45R47R12R34R14L39L16L47L19L12R21L28L14R27R36L3L38R11L23R29R10R34L32L26R7L10L27L10R23R27R49R42R27L49L14R49R38R16R41R25L35L12L9L18R25L13R35L1R8L49L19R39R38R49R32R34L28R15L10R1R23R9L39L29R26L28L27L10L27L1L36R14R48R33R24R37L42R42R32R40R16R2R1L37L32R17L13L41L19R33R18L7L2R3R43L30L21L48R3R31R5R38L47L36L47L30R24R26R37L24R24L39L19L27R9R33R11R37L31L25R8R41L24R10L41L35R16L7L26R42R30L49R34L49R19R43L24L9R11R34L10L21R39L37R6L16R8R1R11L17R14L14R10R35R37R29L10R21L37R41L11L48R33R16L43R43L24L17L12L20R43R25R45R19R45R49L21L42L31L36R16R41R50R43L15L30L26L25R40R34R18R46R24R14L20R12L19R49L1R27L42R11L24R34R35R2L20R26R41R18L42R18R7L4R33L21R13L24R10R15L41L47R34R28L35L8L3R10R14R6R15R2R12R44R41L45L7R3R32R29L26R13R4L38R20L12L26R8L6L15L10R6L45R28L10L26L33L47L19R14R7L13R21R18L17L24L21L24R22R7R40L22R27R37L17R35L22R37R33R6L4L38L4L36L6L5L27L1R49R33R31L3L35L45L43R41R42R44L4R18L29L15L46R14R7R37L16L5R45L41R36L11R18L17L12R29L12L40R28R44L26L9L19R39L36L33L20R42L17L36L21R19R38R33L19R33L25R2L24R6R33R8R43L35L26L23L25L12L10L17R2R45L10R17L7R38L1R2L9R14R35L38L39L38R46R10R28R40L31L34L30R45L2L42L49R28R46R44R11R48R1L38L20L25R8R26L34R3L39L25L48L12R34L25R14R36R39L45R34R25L26L33R22R9L13R26R11L20R20R27R22L26R30R41R38L30R46R25L46R21L11L38L45L47R38R1L15R29R42R6R2L6L32R36R32R22R10L42L31L30R18L21R25L16L9R20L8R38R50L18R3R18R16R16R10L46L14R2R4L48L46L9L18L38L27R29L47R7R13R9R6L26L46R40R36R37R38R12R9L28R36R41L37L35R48L18R17L25L39L12R26L8R31L15R24R43L40R19L32R17R30R24L41R5L23L36R21L39R43L2L33L11R10L38R40R6L11R33L19R24R32L46R18L4R22R32R13L47R8L9R12L18L32L24L19R37R15R32L6R45R31L29L37R26L36L48L46L7L1L2L27R4L45L37R30L49L34L39R36R26R3R1R20R48R35R9L38R14L38L24R7L1R47L14R9L44L11L11R38R14R41R41R1R8R20R46L3R43L49R8R22R27L8R48R7L24R50R23L11R21R14L39L14R3R6R45L42R39L6R25L16L21R44R35L10L30L7L22L13R44R50R5L14L35L19R26R12L26L42R3L27R11R1L10R43L20L28L30L25L21L18R26L19R27L22R36R42R4R42R39R30R5R6L45R10L19R1R18R50L34L2L35R22R45R28R8R45L37L38R7L44L26L36L29R21L29R16R34L40R48R46R30R32R22R40L7R24R24R11R4R5R28R35R29R17L2R42R45L35R14R33L16R45L28R16R23L33R33R41R30L41L25R22R48R42L47R41L17L5R42L27L26R42L5R3R1L16L36R23L31R29L1L4R41R13L3L22R39R17R21L39R29L31L18R31R30L18L11L49L42R21L31L50R50R27R4L42L21R16R13R7L48L41L18R27L10R10R14L31L28L15R37L41L41L13L9L18L25R11R9L1L43R2L7L10L5R15L48L27L19L34R47R21L36R50R49R12R49L25R6R44L45R40R32R3L49L37L31R6L37L32R2R35L49L50R43R37R24L42R43L13L10L38L4L36R41L21R44R17R18R40R9R15L2L31R23R40R33R40L26L1L33R35L24R10R32L10R33L45R2R37R47L30L28R49R21L24R25R46R44R25R40R7R9R41R33L15L12L37R15L4L19L13R2L11L10L41R2R1L15L26L19R45R41L45L18L1R30R37R42L40L50L31R2R36L31L15R19L36R22L11L5R22L50R33R42L42R32L40R17L10R9L14R16L9L22R33L49R14R4R41L17R8R25L41L22L46L32R40L28L48R38L41R26R24R23R25R47R22R45R27R46L4L24L1L49L20L7R11R38R30R7L16L2R16L46L41L16R31R7R5R14R17L21R8L30R27R17R38L15R39L1R45R2L49R10R23L38R22R44R41L25R13L11L17L32L32R21L17L43R35R10L29L32L36R26L4R50L32R6L16L10L16L21R45L35L39L43L4R42R38L37R34R32L9L26R50L47L25R18R46R47R42L35R8"
	inputLines := readStdin()
	_ = movements
	board := getBoard(inputLines)
	p := &Problem{
		board:     board,
		posDepth:  0,
		posWidth:  0,
		direction: "R",
	}
	depth, width := p.getStartPosition()
	p.posDepth = depth
	p.posWidth = width
	p.printBoard(board)
	fmt.Printf("Board start pos depth %d width %d \n", p.posDepth, p.posWidth)
	numbers, directions := getMovs(movements)
	fmt.Printf("Numbers %+v  Directions %+v\n", numbers, directions)
	for i, steps := range numbers {
		p.makeMove(steps)
		// fmt.Printf("After move %d steps\n", steps)
		// p.printBoard(p.board)
		if i < len(directions) {
			p.updateDirection(directions[i])
		}
		// fmt.Printf("Updated directions to %s\n", p.direction)
	}
	p.printResult()
}

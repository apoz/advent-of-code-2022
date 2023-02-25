package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxWidth int = 0
var maxDepth int = 0

const widthOffset int = 150
const depthOffset int = 150

type Dir struct {
	basicDirection string
	dirs           []string
}

var basicDirections []*Dir = []*Dir{
	&Dir{
		basicDirection: "N",
		dirs:           []string{"NE", "N", "NW"},
	},
	&Dir{
		basicDirection: "S",
		dirs:           []string{"SE", "S", "SW"},
	},
	&Dir{
		basicDirection: "W",
		dirs:           []string{"NW", "W", "SW"},
	},
	&Dir{
		basicDirection: "E",
		dirs:           []string{"NE", "E", "SE"},
	},
}

var directions map[string]Position = map[string]Position{
	"N":  Position{depth: -1, width: 0},
	"NE": Position{depth: -1, width: 1},
	"E":  Position{depth: 0, width: 1},
	"SE": Position{depth: 1, width: 1},
	"S":  Position{depth: 1, width: 0},
	"SW": Position{depth: 1, width: -1},
	"W":  Position{depth: 0, width: -1},
	"NW": Position{depth: -1, width: -1},
}

func rotateBasicDirections() {
	newDirections := basicDirections[1:]
	newDirections = append(newDirections, basicDirections[0])
	basicDirections = newDirections
}

type Position struct {
	depth int
	width int
}

func (p Position) allNeighbours() []Position {
	var posList = make([]Position, 0)
	for k := range directions {
		posList = append(posList, p.add(directions[k]))
	}
	return posList
}

func (p Position) someNeighbours(dirs []string) []Position {
	var posList = make([]Position, 0)
	for _, k := range dirs {
		posList = append(posList, p.add(directions[k]))
	}
	return posList
}

func (p Position) add(myPos Position) Position {
	return Position{
		depth: p.depth + myPos.depth,
		width: p.width + myPos.width,
	}
}

func (p Position) print() {
	//fmt.Printf("Position (depth: %d, width: %d)\n", p.depth, p.width)
}

type Elf struct {
	identifier  int
	currentPos  Position
	proposedPos Position
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func fillUp(lines []string, rounds int) []string {
	widthFill := ""
	for i := 0; i < rounds; i++ {
		widthFill += "."
	}
	currentWidth := len(strings.Split(lines[0], ""))
	emptyLine := widthFill
	for i := 0; i < currentWidth; i++ {
		emptyLine += "."
	}
	emptyLine += widthFill

	newLines := make([]string, 0)
	for i := 0; i < rounds; i++ {
		newLines = append(newLines, emptyLine)
	}
	for _, line := range lines {
		newLines = append(newLines, widthFill+line+widthFill)
	}
	for i := 0; i < rounds; i++ {
		newLines = append(newLines, emptyLine)
	}
	return newLines
}

func contains(s []Position, p Position) bool {
	for _, v := range s {
		if v.depth == p.depth && v.width == p.width {
			return true
		}
	}
	return false
}

func anyAdjacentElf(id int, elfs []*Elf) bool {
	var currentPos Position = getCurrentPositionForElf(id, elfs)
	allNeigh := currentPos.allNeighbours()
	for _, elf := range elfs {
		if elf.identifier == id {
			continue
		}
		if contains(allNeigh, elf.currentPos) {
			return true
		}
	}
	return false
}

func getCurrentPositionForElf(id int, elfs []*Elf) Position {
	for _, elf := range elfs {
		if (*elf).identifier == id {
			return (*elf).currentPos
		}
	}
	return Position{}
}

func cleanProposals(elfs []*Elf) {
	for _, elf := range elfs {
		(*elf).proposedPos = Position{}
	}
}

func getMinWidth(elves []*Elf) int {
	minWidth := 999999
	for _, elf := range elves {
		if (*elf).currentPos.width < minWidth {
			minWidth = (*elf).currentPos.width
		}
	}
	return minWidth
}

func getMaxWidth(elves []*Elf) int {
	maxWidth := 0
	for _, elf := range elves {
		if (*elf).currentPos.width > maxWidth {
			maxWidth = (*elf).currentPos.width
		}
	}
	return maxWidth
}

func getMinDepth(elves []*Elf) int {
	minDepth := 999999
	for _, elf := range elves {
		if (*elf).currentPos.depth < minDepth {
			minDepth = (*elf).currentPos.depth
		}
	}
	return minDepth
}

func getMaxDepth(elves []*Elf) int {
	maxDepth := 0
	for _, elf := range elves {
		if (*elf).currentPos.depth > maxDepth {
			maxDepth = (*elf).currentPos.depth
		}
	}
	return maxDepth
}

func score(elves []*Elf) int {
	minWidth := getMinWidth(elves)
	maxWidth := getMaxWidth(elves)
	minDepth := getMinDepth(elves)
	maxDepth := getMaxDepth(elves)

	totalTiles := (maxWidth + 1 - minWidth) * (maxDepth + 1 - minDepth)
	elfCount := len(elves)
	return totalTiles - elfCount
}

func anyElfInThesePositions(elfs []*Elf, positions []Position) bool {
	for _, elf := range elfs {
		if contains(positions, (*elf).currentPos) {
			return true
		}
	}
	return false
}

func otherElfProposedSame(id int, elfs []*Elf) bool {
	proposedPos := Position{}
	for _, elf := range elfs {
		if (*elf).identifier == id {
			proposedPos = (*elf).proposedPos
		}
	}

	for _, elf := range elfs {
		if (*elf).identifier == id {
			continue
		}
		if (*elf).proposedPos.depth == proposedPos.depth &&
			(*elf).proposedPos.width == proposedPos.width {
			return true
		}
	}
	return false
}

func getElfs(inputLines []string) []*Elf {
	maxDepth = len(inputLines)
	maxWidth = len(strings.Split(inputLines[0], ""))
	elves := make([]*Elf, 0)
	i := 1
	for depth, line := range inputLines {
		for width, c := range strings.Split(line, "") {
			if c == "#" {
				currentPos := Position{
					width: width + widthOffset,
					depth: depth + depthOffset,
				}
				elf := Elf{
					identifier:  i,
					currentPos:  currentPos,
					proposedPos: Position{},
				}
				i++
				elves = append(elves, &elf)
			}
		}
	}
	return elves
}

func printScenario(elves []*Elf) {
	matrix := make([][]string, 0)
	maxWidth := getMaxWidth(elves)
	maxDepth := getMaxDepth(elves)
	for i := 0; i < maxDepth; i++ {
		line := make([]string, 0)
		for j := 0; j < maxWidth; j++ {
			line = append(line, ".")
		}
		matrix = append(matrix, line)
	}
	for _, elf := range elves {
		matrix[elf.currentPos.depth][elf.currentPos.width] = "#"
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%s ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func printScenarioHighlightElf(elves []*Elf, identifier int) {
	matrix := make([][]string, 0)
	maxWidth := getMaxWidth(elves)
	maxDepth := getMaxDepth(elves)

	for i := 0; i < maxDepth; i++ {
		line := make([]string, 0)
		for j := 0; j < maxWidth; j++ {
			line = append(line, ".")
		}
		matrix = append(matrix, line)
	}
	for _, elf := range elves {
		if (*elf).identifier == identifier {
			matrix[elf.currentPos.depth][elf.currentPos.width] = "X"
		} else {
			matrix[elf.currentPos.depth][elf.currentPos.width] = "#"
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%s ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func printScenarioHighlightElfProposed(elves []*Elf, identifier int) {
	matrix := make([][]string, 0)
	printElves(elves)
	maxWidth := getMaxWidth(elves)
	maxDepth := getMaxDepth(elves)

	for i := 0; i < maxDepth; i++ {
		line := make([]string, 0)
		for j := 0; j < maxWidth; j++ {
			line = append(line, ".")
		}
		matrix = append(matrix, line)
	}
	for _, elf := range elves {
		if (*elf).identifier == identifier {
			matrix[elf.currentPos.depth][elf.currentPos.width] = "X"
			matrix[elf.proposedPos.depth][elf.proposedPos.width] = "?"
		} else {
			matrix[elf.currentPos.depth][elf.currentPos.width] = "#"
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%s ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func printElves(elves []*Elf) {
	for _, elf := range elves {
		fmt.Printf("Elf %d  depth %d  width %d \n", (*elf).identifier, (*elf).currentPos.depth, (*elf).currentPos.width)
	}
}

func iterate(elfs []*Elf) int {
	for _, elf := range elfs {
		(*elf).proposedPos = (*elf).currentPos
		// Check if elf has immediate neighbours
		if anyAdjacentElf((*elf).identifier, elfs) == false {
			continue
		}

		//else
		for _, myDir := range basicDirections {
			posToCheck := (*elf).currentPos.someNeighbours(myDir.dirs)
			if anyElfInThesePositions(elfs, posToCheck) == false {
				(*elf).proposedPos = (*elf).currentPos.add(directions[myDir.basicDirection])
				break
			}
		}
	}

	movingElfs := 0
	// make proposed pos the current ones
	for _, elf := range elfs {
		if otherElfProposedSame((*elf).identifier, elfs) == false &&
			(((*elf).currentPos.depth != (*elf).proposedPos.depth) ||
				(*elf).currentPos.width != (*elf).proposedPos.width) {
			(*elf).currentPos = (*elf).proposedPos
			movingElfs++
		}
	}

	cleanProposals(elfs)
	return movingElfs
}

func main() {
	inputLines := readStdin()
	elves := getElfs(inputLines)
	var moved_elves int = 0
	var rounds int = 1
	for true {
		fmt.Printf("Rounds %d\n", rounds)
		moved_elves = iterate(elves)
		fmt.Printf("Moved elves %d\n", moved_elves)
		if moved_elves == 0 {
			fmt.Printf("The number of rounds is %d", rounds)
			os.Exit(0)
		}
		rotateBasicDirections()
		rounds++
	}
	fmt.Printf("The score is %d\n", score(elves))

}

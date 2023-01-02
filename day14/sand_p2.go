package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	matrix           [][]string // depth, dist
	distBase         int
	dataPoints       [][]datapoint
	maxDepth         int
	maxDist          int
	sandDroppingDist int
}

func (p *Problem) setSandDroppingDist(dist int) {
	p.sandDroppingDist = dist
}

func (p *Problem) getBaseMatrix(distLimit int, depthLimit int) [][]string {
	fmt.Printf("Base matrix is distLimit %d depthLimit %d\n", distLimit, depthLimit)
	var matrix [][]string = make([][]string, 0)
	for depth := 0; depth < depthLimit; depth++ {
		row := make([]string, 0)
		for dist := 0; dist < distLimit; dist++ {
			row = append(row, ".")
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func (p *Problem) getDistLimits(dataPoints [][]datapoint) (int, int) {
	var minDist int = 9999999
	var maxDist int = 0
	for depth := 0; depth < len(dataPoints); depth++ {
		for dist := 0; dist < len(dataPoints[depth]); dist++ {
			dp := dataPoints[depth][dist]
			if dp.dist < minDist {
				minDist = dp.dist
			}
			if dp.dist > maxDist {
				maxDist = dp.dist
			}
		}
	}
	return minDist, maxDist
}

func (p *Problem) countSandBlocks() int {
	var count int = 0
	for depth := 0; depth < len(p.matrix); depth++ {
		for dist := 0; dist < len(p.matrix[depth]); dist++ {
			if p.matrix[depth][dist] == "O" {
				count++
			}
		}
	}
	return count
}

func (p *Problem) getDepthLimits(dataPoints [][]datapoint) (int, int) {
	var minDepth int = 9999999
	var maxDepth int = 0
	for depth := 0; depth < len(dataPoints); depth++ {
		for dist := 0; dist < len(dataPoints[depth]); dist++ {
			if dataPoints[depth][dist].depth < minDepth {
				minDepth = dataPoints[depth][dist].depth
			}
			if dataPoints[depth][dist].depth > maxDepth {
				maxDepth = dataPoints[depth][dist].depth
			}
		}
	}
	return minDepth, maxDepth
}

type datapoint struct {
	dist  int
	depth int
}

func (p *Problem) getDataPoints(lines []string) [][]datapoint {
	var dataPoints [][]datapoint = make([][]datapoint, 0)
	for _, line := range lines {
		items := strings.Split(line, " -> ")
		lineDatapoint := make([]datapoint, 0)
		for _, item := range items {
			dataPoint := datapoint{}
			fmt.Printf("item %s\n", item)
			point := strings.Split(item, ",")
			fmt.Printf("%s\n", point)
			dist, _ := strconv.Atoi(string(point[0]))
			depth, _ := strconv.Atoi(string(point[1]))
			dataPoint.dist = dist
			dataPoint.depth = depth
			lineDatapoint = append(lineDatapoint, dataPoint)
		}
		dataPoints = append(dataPoints, lineDatapoint)
	}

	return dataPoints
}

func (p *Problem) printMatrix() {
	fmt.Printf("distBase %d\n", p.distBase)
	s := strconv.Itoa(p.distBase)
	println("==========================")
	for i := 0; i < (len(s)); i++ {
		fmt.Printf("    %c\n", s[i])
	}
	for depth := 0; depth < len(p.matrix); depth++ {
		fmt.Printf("%d  ", depth)
		for dist := 0; dist < len(p.matrix[depth]); dist++ {
			fmt.Printf(" %s", p.matrix[depth][dist])
		}
		print("\n")
	}
	println("==========================")
}

func (p *Problem) sandBlockFall() bool {
	currentDist := p.sandDroppingDist - p.distBase
	currentDepth := 0
	var fallDown bool = false
	fmt.Printf("HERE starting with currentDist %d, currentDepth %d\n", currentDist, currentDepth)
	canMove := true
	for canMove == true {
		fmt.Printf("BBBBIn the for loop currentDist %d currentDepth %d len(p.matrix) %d\n", currentDist, currentDepth, len(p.matrix))
		if p.matrix[0][p.sandDroppingDist-p.distBase] == "O" {
			canMove = false
			return true
		}
		if currentDepth == len(p.matrix)-1 || currentDist == len(p.matrix[0])-1 {
			fmt.Println("Moving out here\n")
			canMove = false
			fallDown = true
		} else if p.matrix[currentDepth+1][currentDist] == "." {
			fmt.Printf("Moving one down \n")
			currentDepth++
		} else if p.matrix[currentDepth+1][currentDist] == "#" || p.matrix[currentDepth+1][currentDist] == "O" {
			fmt.Printf("Next vertical move is blocked \n")
			if currentDist == len(p.matrix[0]) || currentDepth == len(p.matrix) {
				canMove = false
				fallDown = true
			}
			if p.matrix[currentDepth+1][currentDist-1] == "." {
				fmt.Printf("I can move to the left \n")
				currentDist--
			} else if p.matrix[currentDepth+1][currentDist+1] == "." {
				fmt.Printf("I can move to the right \n")
				currentDist++
			} else {
				fmt.Printf("I cannot move, setting the sand\n")
				p.setMatrixDatapoint(currentDist, currentDepth, "O")
				canMove = false
				return false
			}
			currentDepth++
		}
	}
	return fallDown
}

func (p *Problem) setFloor() {
	for dist := 0; dist < len(p.matrix[len(p.matrix)-1]); dist++ {
		p.setMatrixDatapoint(dist, len(p.matrix)-1, "#")
	}
}

func (p *Problem) getMatrix(lines []string) [][]string {
	if p.matrix != nil {
		return p.matrix
	}
	datapoints := p.getDataPoints(lines)
	fmt.Printf("Got the datapoints %+v\n", datapoints)
	p.dataPoints = datapoints
	minDist, maxDist := p.getDistLimits(p.dataPoints)
	fmt.Printf("minDist %d  maxDist %d\n", minDist, maxDist)
	//p.distBase = minDist - 5000
	p.distBase = 0
	minDepth, maxDepth := p.getDepthLimits(p.dataPoints)
	fmt.Printf("minDepth %d maxDepth %d\n", minDepth, maxDepth)
	distLimit := maxDist + 5000 - minDist + 5000 + 1
	depthLimit := maxDepth + 1
	matrix := p.getBaseMatrix(distLimit, depthLimit+2)
	p.matrix = matrix
	p.setFloor()
	p.printMatrix()
	p.processDatapoints()
	p.setSandDroppingDist(500)
	p.setMatrixDatapoint(500, 0, "+")
	return matrix
}

func (p *Problem) processDatapoints() {
	for i := 0; i < len(p.dataPoints); i++ {
		for j := 0; j < (len(p.dataPoints[i]) - 1); j++ {
			dps := p.interpolateDatapoints(p.dataPoints[i][j], p.dataPoints[i][j+1])
			for _, dp := range dps {
				p.setMatrixDatapoint(dp.dist, dp.depth, "#")
			}
		}
	}

}

func (p *Problem) interpolateDatapoints(orig datapoint, dest datapoint) []datapoint {
	var points []datapoint = make([]datapoint, 0)
	if orig.dist == dest.dist {
		for i := min(orig.depth, dest.depth); i <= max(orig.depth, dest.depth); i++ {
			dp := datapoint{}
			dp.dist = orig.dist
			dp.depth = i
			points = append(points, dp)
		}
	} else if orig.depth == dest.depth {
		for i := min(orig.dist, dest.dist); i <= max(orig.dist, dest.dist); i++ {
			dp := datapoint{}
			dp.depth = orig.depth
			dp.dist = i
			points = append(points, dp)
		}
	} else {
		fmt.Printf("Something is wrong")
	}
	return points
}

func (p *Problem) setMatrixDatapoint(dist int, depth int, value string) {
	if dist < p.distBase {
		p.matrix[depth][dist] = value
		return
	}
	p.matrix[depth][dist-p.distBase] = value
	return
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	inputLines := readStdin()
	prob := &Problem{}
	prob.getMatrix(inputLines)
	prob.printMatrix()
	fallOut := prob.sandBlockFall()
	for fallOut == false {
		fallOut = prob.sandBlockFall()
	}
	fmt.Println("FINAL STATUS")
	prob.printMatrix()
	fmt.Printf("The number of sand blocks is %d\n", prob.countSandBlocks())
}

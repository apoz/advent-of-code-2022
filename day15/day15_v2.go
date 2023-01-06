package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Board struct {
	matrix        []string // depth, dist
	dataPoints    []datapoint
	minDepth      int
	maxDepth      int
	minDist       int
	maxDist       int
	responseDepth int
}

type datapoint struct {
	sensorDepth   int
	sensorDist    int
	beaconDepth   int
	beaconDist    int
	manhattanDist int
}

func (dp *datapoint) print() {
	fmt.Printf("Sensor (depth=%d,dist=%d)   Beacon(depth=%d, dist=%d) manhattanDist=%d\n", dp.sensorDepth, dp.sensorDist, dp.beaconDepth, dp.beaconDist, dp.manhattanDist)
}

func (b *Board) getDepthLimits() (int, int) {
	minDepth := 9999999
	maxDepth := 0
	for _, dp := range b.dataPoints {
		if dp.sensorDepth-dp.manhattanDist < minDepth {
			minDepth = dp.sensorDepth - dp.manhattanDist
		}
		if dp.sensorDepth+dp.manhattanDist > maxDepth {
			maxDepth = dp.sensorDepth + dp.manhattanDist
		}
	}
	return minDepth, maxDepth
}

func (b *Board) getDistLimits() (int, int) {
	minDist := 9999999
	maxDist := 0
	for _, dp := range b.dataPoints {
		if dp.sensorDist-dp.manhattanDist < minDist {
			minDist = dp.sensorDist - dp.manhattanDist
		}
		if dp.sensorDist+dp.manhattanDist > maxDist {
			maxDist = dp.sensorDist + dp.manhattanDist
		}
	}
	return minDist, maxDist
}

func (b *Board) getBaseMatrix() []string {
	distLen := b.maxDist + 1 - b.minDist
	fmt.Printf("distLen is %d\n", distLen)
	// depthLen := b.maxDepth + 1 - b.minDepth

	row := make([]string, 0)
	for dist := 0; dist < distLen; dist++ {
		row = append(row, ".")
	}
	return row
}

func (b *Board) printMatrix() {
	fmt.Println("==========================")
	fmt.Printf("%d  ", b.responseDepth)
	for dist := 0; dist < len(b.matrix); dist++ {
		fmt.Printf(" %s", b.matrix[dist])
	}
	fmt.Printf("\n")
	fmt.Println("==========================")
}

func (b *Board) getDataPoint(line string) datapoint {
	fmt.Println(line)
	r := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	res := r.FindStringSubmatch(line)
	dp := datapoint{}
	dp.sensorDist, _ = strconv.Atoi(res[1])
	dp.sensorDepth, _ = strconv.Atoi(res[2])
	dp.beaconDist, _ = strconv.Atoi(res[3])
	dp.beaconDepth, _ = strconv.Atoi(res[4])
	fmt.Printf("sensorDist %d sensorDepth %d beaconDist %d beaconDepth %d\n", dp.sensorDist, dp.sensorDepth, dp.beaconDist, dp.beaconDepth)
	dp.manhattanDist = b.calculateDistance(dp.sensorDist, dp.sensorDepth, dp.beaconDist, dp.beaconDepth)
	return dp
}

func (p *Board) getDataPoints(lines []string) []datapoint {
	var dataPoints []datapoint = make([]datapoint, 0)
	for _, line := range lines {
		lineDatapoint := p.getDataPoint(line)
		// if (lineDatapoint.sensorDepth-lineDatapoint.manhattanDist < p.responseDepth) &&
		// 	(lineDatapoint.sensorDepth+lineDatapoint.manhattanDist > p.responseDepth) {
		dataPoints = append(dataPoints, lineDatapoint)
		// } else {
		// 	fmt.Print("Discarding datapoint!")
		// 	lineDatapoint.print()

		// }
	}
	return dataPoints
}

func (b *Board) setMatrixCell(dist int, s string) {
	fmt.Printf("setting depth=%d dist=%d to %s\n", b.responseDepth, dist, s)
	matrixDist := dist - b.minDist
	b.matrix[matrixDist] = s
}

func (p *Board) calculateDistance(sensorDist int, sensorDepth int, beaconDist int, beaconDepth int) int {
	return abs(sensorDist-beaconDist) + abs(sensorDepth-beaconDepth)
}

func (p *Board) getMatrixCell(dist int) string {
	matrixDist := dist - p.minDist
	return p.matrix[matrixDist]

}

func (p *Board) getNumberOfPositions(s string) int {
	counter := 0
	for dist := 0; dist < len(p.matrix); dist++ {
		if p.matrix[dist] == s {
			counter++
		}
	}
	return counter
}

func (p *Board) getMatrix(lines []string) []string {
	p.dataPoints = p.getDataPoints(lines)
	fmt.Printf("Got the datapoints %+v\n", p.dataPoints)
	p.minDist, p.maxDist = p.getDistLimits()
	p.minDepth, p.maxDepth = p.getDepthLimits()
	fmt.Printf("minDist %d  maxDist %d\n", p.minDist, p.maxDist)
	fmt.Printf("minDepth %d  maxDepth %d\n", p.minDepth, p.maxDepth)
	p.matrix = p.getBaseMatrix()
	for _, dp := range p.dataPoints {
		if dp.sensorDepth == p.responseDepth {
			p.setMatrixCell(dp.sensorDist, "S")
		}
		if dp.beaconDepth == p.responseDepth {
			p.setMatrixCell(dp.beaconDist, "B")
		}
	}
	fmt.Println("After setting S and B")
	// p.printMatrix()

	fmt.Println("Up to here we're OK")
	for _, dp := range p.dataPoints {
		fmt.Printf("HELLOO len(p.matrix)=%d\n", len(p.matrix))
		fmt.Printf("Processing the following datapoint\n")
		dp.print()
		for dist := p.minDist; dist < p.maxDist; dist++ {
			manhattanDist := p.calculateDistance(dp.sensorDist, dp.sensorDepth, dist, p.responseDepth)
			// fmt.Printf("Checking depth=%d dist=%d. The distance is %d vs %d \n", p.responseDepth, dist, manhattanDist, dp.manhattanDist)
			if manhattanDist <= dp.manhattanDist {
				// fmt.Printf(
				// 	"Distance is smaller for sensor depth=%d dist=%d for depth=%d dist=%d   (%d <= %d)\n",
				// 	dp.sensorDepth,
				// 	dp.sensorDist,
				// 	p.responseDepth,
				// 	dist,
				// 	manhattanDist,
				// 	dp.manhattanDist,
				// )
				if p.getMatrixCell(dist) == "." {
					p.setMatrixCell(dist, "#")
				}
			}
		}
	}
	// p.setMatrixCell(dp.sensorDist, dp.sensorDepth, "S")
	// p.setMatrixCell(dp.beaconDist, dp.beaconDepth, "B")
	// p.printMatrix()

	// p.distBase = minDist
	// minDepth, maxDepth := p.getDepthLimits(p.dataPoints)
	// fmt.Printf("minDepth %d maxDepth %d\n", minDepth, maxDepth)
	// distLimit := maxDist - minDist + 1
	// depthLimit := maxDepth + 1

	// p.matrix = matrix
	// p.printMatrix()
	// fmt.Printf("Before process datapoints\n")
	// p.processDatapoints()
	// fmt.Printf("After process datapoints\n")
	// p.setSandDroppingDist(500)
	// p.setMatrixDatapoint(500, 0, "+")
	for _, dp := range p.dataPoints {
		if dp.sensorDepth == p.responseDepth {
			p.setMatrixCell(dp.sensorDist, "S")
		}
		if dp.beaconDepth == p.responseDepth {
			p.setMatrixCell(dp.beaconDist, "B")
		}
	}
	return nil
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	inputLines := readStdin()
	prob := &Board{}
	// prob.responseDepth = 10
	prob.responseDepth = 2000000
	prob.getMatrix(inputLines)
	fmt.Printf("The number of impossible positions for depth=%d is %d\n", prob.responseDepth, prob.getNumberOfPositions("#"))
	// prob.printMatrix()
	// fallOut := prob.sandBlockFall()
	// for fallOut == false {
	// 	// 	fmt.Printf("AAAIn the for loop\n")
	// 	// 	// prob.printMatrix()
	// 	fallOut = prob.sandBlockFall()
	// 	fmt.Printf("In the for loop after sandBlock fall\n")
	// 	prob.printMatrix()
	// }
	// // fmt.Println("FINAL STATUS")
	// prob.printMatrix()
	// fmt.Printf("The number of sand blocks is %d\n", prob.countSandBlocks())
}

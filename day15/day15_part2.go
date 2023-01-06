package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Board struct {
	matrix          []string // depth, dist
	dataPoints      []datapoint
	minDepth        int
	maxDepth        int
	minDist         int
	maxDist         int
	responseDepth   int
	perimeterPoints []point
}

type datapoint struct {
	sensorDepth   int
	sensorDist    int
	beaconDepth   int
	beaconDist    int
	manhattanDist int
}

type point struct {
	depth int
	dist  int
}

func (dp *datapoint) print() {
	fmt.Printf("Sensor (depth=%d,dist=%d)   Beacon(depth=%d, dist=%d) manhattanDist=%d\n", dp.sensorDepth, dp.sensorDist, dp.beaconDepth, dp.beaconDist, dp.manhattanDist)
}

func (b *Board) getPerimeterPointsForSensor(s datapoint) []point {
	topDepth := s.sensorDepth - (s.manhattanDist + 1)
	topDist := s.sensorDist
	bottomDepth := s.sensorDepth + (s.manhattanDist + 1)
	bottomDist := s.sensorDist
	leftDepth := s.sensorDepth
	leftDist := s.sensorDist - (s.manhattanDist + 1)
	rightDepth := s.sensorDepth
	rightDist := s.sensorDist + (s.manhattanDist + 1)
	fmt.Printf("Top (depth=%d, dist=%d)\n", topDepth, topDist)
	fmt.Printf("Bottom (depth=%d, dist=%d)\n", bottomDepth, bottomDist)
	fmt.Printf("Left (depth=%d, dist=%d)\n", leftDepth, leftDist)
	fmt.Printf("Rigth (depth=%d, dist=%d)\n", rightDepth, rightDist)

	points := make([]point, 0)
	points1 := b.getPointsInLineGrowingSlope(leftDepth, leftDist, topDepth, topDist)
	for _, p := range points1 {
		points = append(points, p)
	}
	points2 := b.getPointsInLineGrowingSlope(bottomDepth, bottomDist, rightDepth, rightDist)
	for _, p := range points2 {
		points = append(points, p)
	}
	points3 := b.getPointsInLineDecreasingSlope(topDepth, topDist, rightDepth, rightDist)
	for _, p := range points3 {
		points = append(points, p)
	}
	points4 := b.getPointsInLineDecreasingSlope(leftDepth, leftDist, bottomDepth, bottomDist)
	for _, p := range points4 {
		points = append(points, p)
	}
	return points
}

func (b *Board) getPointsInLineGrowingSlope(origDepth int, origDist int, destDepth int, destDist int) []point {
	points := make([]point, 0)
	for i := 0; i <= (destDist - origDist); i++ {
		point := point{
			depth: origDepth - i,
			dist:  origDist + i,
		}
		points = append(points, point)
	}
	return points

}

func (b *Board) getPointsInLineDecreasingSlope(origDepth int, origDist int, destDepth int, destDist int) []point {
	points := make([]point, 0)
	for i := 0; i <= (destDepth - origDepth); i++ {
		point := point{
			depth: origDepth + i,
			dist:  origDist + i,
		}
		points = append(points, point)
	}
	return points

}

func (b *Board) getPerimeterPoints() []point {
	points := make([]point, 0)
	for _, dp := range b.dataPoints {
		fmt.Println("Getting perimeter for: ")
		dp.print()
		sensorPoints := b.getPerimeterPointsForSensor(dp)
		for _, p := range sensorPoints {
			manDist := b.calculateDistance(dp.sensorDist, dp.sensorDepth, p.dist, p.depth)
			if manDist != dp.manhattanDist+1 {
				fmt.Printf("Something is wrong! (depth=%d, dist=%d) is not on the perimeter. Manhattan distance is %d\n", p.depth, p.dist, manDist)
			} else {
				if p.depth > 0 && p.dist > 0 {
					points = append(points, p)
				}

			}
		}

	}
	return points
}

func (b *Board) isOutsideOfAllSensors(p point) bool {
	for _, dp := range b.dataPoints {
		if b.calculateDistance(dp.sensorDist, dp.sensorDepth, p.dist, p.depth) <= dp.manhattanDist {
			return false
		}
	}
	return true
}

func (b *Board) isOnSquare(pp point) bool {
	if pp.depth > 0 && pp.depth <= 4000000 && pp.dist > 0 && pp.dist < 4000000 {
		// if pp.depth > 0 && pp.depth <= 20 && pp.dist > 0 && pp.dist <= 20 {
		return true
	}
	return false
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
		dataPoints = append(dataPoints, lineDatapoint)
	}
	return dataPoints
}

func (p *Board) calculateDistance(sensorDist int, sensorDepth int, beaconDist int, beaconDepth int) int {
	return abs(sensorDist-beaconDist) + abs(sensorDepth-beaconDepth)
}

func (p *Board) getMatrix(lines []string) []string {
	p.dataPoints = p.getDataPoints(lines)
	fmt.Printf("Got the datapoints %+v\n", p.dataPoints)
	p.minDist, p.maxDist = p.getDistLimits()
	p.minDepth, p.maxDepth = p.getDepthLimits()
	fmt.Printf("minDist %d  maxDist %d\n", p.minDist, p.maxDist)
	fmt.Printf("minDepth %d  maxDepth %d\n", p.minDepth, p.maxDepth)
	p.perimeterPoints = p.getPerimeterPoints()
	counter := 0
	for _, pp := range p.perimeterPoints {
		counter++
		if p.isOnSquare(pp) && p.isOutsideOfAllSensors(pp) {
			fmt.Printf("Found point dist=%d depth=%d\n", pp.dist, pp.depth)
			fmt.Printf("Tuning Freq %d\n", pp.dist*4000000+pp.depth)
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
	//fmt.Printf("The number of impossible positions for depth=%d is %d\n", prob.responseDepth, prob.getNumberOfPositions("#"))
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

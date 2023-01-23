package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type GemInventory struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type RobotInventory struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type Status struct {
	robots RobotInventory
	gems   GemInventory
}

type Path struct {
	path []Status
}

type Blueprint struct {
	robotCost map[string]GemInventory
	id        int
	paths     []Path
	maxOreRobots int
	maxClayRobots int
	maxObsidianRobots int
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getInitialStatus() Status {
	robots := RobotInventory{
		ore:      1,
		clay:     0,
		obsidian: 0,
		geode:    0,
	}
	gems := GemInventory{
		ore:      0,
		clay:     0,
		obsidian: 0,
		geode:    0,
	}
	status := Status{
		robots: robots,
		gems:   gems,
	}
	return status
}

func getBlueprintForLine(line string) Blueprint {

	r := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	res := r.FindStringSubmatch(line)

	id, _ := strconv.Atoi(res[1])

	robotCost := make(map[string]GemInventory)

	oreRobotCost := GemInventory{}
	cost, _ := strconv.Atoi(res[2])
	oreRobotCost.ore = cost
	robotCost["ore"] = oreRobotCost

	clayRobotCost := GemInventory{}

	cost, _ = strconv.Atoi(res[3])
	clayRobotCost.ore = cost
	robotCost["clay"] = clayRobotCost

	obsidianRobotCost := GemInventory{}

	cost, _ = strconv.Atoi(res[4])
	obsidianRobotCost.ore = cost
	cost, _ = strconv.Atoi(res[5])
	obsidianRobotCost.clay = cost
	robotCost["obsidian"] = obsidianRobotCost

	geodeRobotCost := GemInventory{}
	cost, _ = strconv.Atoi(res[6])
	geodeRobotCost.ore = cost
	cost, _ = strconv.Atoi(res[7])
	geodeRobotCost.obsidian = cost
	robotCost["geode"] = geodeRobotCost

	paths := make([]Path, 0)
	path := Path{
		path: make([]Status, 0),
	}
	path.path = append(path.path, getInitialStatus())
	paths = append(paths, path)

	bp := Blueprint{
		robotCost: robotCost,
		paths:     paths,
		id:        id,
	}

	return bp
}

func getOptionsForPath(path Path, robotCost map[string]GemInventory) []string {
	operations := make([]string, 0)
	lastStatus := path.path[len(path.path)-1]
	operations = append(operations, "do_nothing")
	if canBuildRobot(lastStatus, robotCost["ore"]) {
		operations = append(operations, "build_ore_robot")
	}
	if canBuildRobot(lastStatus, robotCost["clay"]) {
		operations = append(operations, "build_clay_robot")
	}
	if canBuildRobot(lastStatus, robotCost["obsidian"]) {
		operations = append(operations, "build_obsidian_robot")
	}
	if canBuildRobot(lastStatus, robotCost["geode"]) {
		operations = append(operations, "build_geode_robot")
	}
	return operations
}

func applyOperationToPath(operation string, path Path, cost map[string]GemInventory) Path {
	// fmt.Printf("Applying operation %s to path %+v\n", operation, path)
	status := path.path[len(path.path)-1]
	newStatus := copyStatus(status)
	if operation == "do_nothing" {
		return path
	} else if operation == "build_ore_robot" {
		newStatus.gems.clay -= cost["ore"].clay
		newStatus.gems.obsidian -= cost["ore"].obsidian
		newStatus.gems.ore -= cost["ore"].ore
		newStatus.gems.geode -= cost["ore"].geode
		newStatus.robots.ore++
	} else if operation == "build_clay_robot" {
		newStatus.gems.clay -= cost["clay"].clay
		newStatus.gems.obsidian -= cost["clay"].obsidian
		newStatus.gems.ore -= cost["clay"].ore
		newStatus.gems.geode -= cost["clay"].geode
		newStatus.robots.clay++
	} else if operation == "build_obsidian_robot" {
		newStatus.gems.clay -= cost["obsidian"].clay
		newStatus.gems.obsidian -= cost["obsidian"].obsidian
		newStatus.gems.ore -= cost["obsidian"].ore
		newStatus.gems.geode -= cost["obsidian"].geode
		newStatus.robots.obsidian++
	} else if operation == "build_geode_robot" {
		newStatus.gems.clay -= cost["geode"].clay
		newStatus.gems.obsidian -= cost["geode"].obsidian
		newStatus.gems.ore -= cost["geode"].ore
		newStatus.gems.geode -= cost["geode"].geode
		newStatus.robots.geode++
	}
	newPath := Path{
		path: append(path.path, newStatus),
	}
	// fmt.Printf("Path after %+v\n", newPath)
	return newPath
}

func canBuildRobot(status Status, cost GemInventory) bool {
	if status.gems.clay >= cost.clay &&
		status.gems.obsidian >= cost.obsidian &&
		status.gems.ore >= cost.ore {
		return true
	}
	return false
}

func copyStatus(status Status) Status {
	robots := RobotInventory{
		ore:      status.robots.ore,
		clay:     status.robots.clay,
		obsidian: status.robots.obsidian,
		geode:    status.robots.geode,
	}
	gemInventory := GemInventory{
		ore:      status.gems.ore,
		clay:     status.gems.clay,
		obsidian: status.gems.obsidian,
		geode:    status.gems.geode,
	}
	newStatus := Status{
		robots: robots,
		gems:   gemInventory,
	}
	return newStatus
}

func copyPath(path Path) Path {
	statusList := make([]Status, 0)
	for _, status := range path.path {
		newStatus := copyStatus(status)
		statusList = append(statusList, newStatus)
	}
	newPath := Path{
		path: statusList,
	}
	return newPath
}
func getInventoyIncrease(robots RobotInventory, gems GemInventory) GemInventory {
	newGems := GemInventory{
		clay:     robots.clay + gems.clay,
		obsidian: robots.obsidian + gems.obsidian,
		ore:      robots.ore + gems.ore,
		geode:    robots.geode + gems.geode,
	}
	return newGems
}

func max(a int, b int) {
	if a > b {
		return a
	}
	return b
}

func DFS(time int, path Path, robotCost map[string]GemInventory) []Path {
	paths := []Path{path}

	// fmt.Printf("Calling DFS time=%d with path %+v\n", time, path)
	options := getOptionsForPath(path, robotCost)
	// fmt.Printf("The options for the current path are %+v\n", options)
	for _, operation := range options {
		newTime := time - 1
		if time <= 0 {
			continue
		}
		newPath := copyPath(path)
		currentStatus := newPath.path[len(newPath.path)-1]
		currentRobots := currentStatus.robots
		newPath = applyOperationToPath(operation, path, robotCost)

		newPath.path[len(newPath.path)-1].gems.ore += currentRobots.ore
		newPath.path[len(newPath.path)-1].gems.clay += currentRobots.clay
		newPath.path[len(newPath.path)-1].gems.obsidian += currentRobots.obsidian
		newPath.path[len(newPath.path)-1].gems.geode += currentRobots.geode
		// fmt.Printf("After increase path is %+v\n", newPath.path[len(newPath.path)-1])
		paths = append(paths, DFS(newTime, newPath, robotCost)...)

	}
	return paths
}

// func DFS(current string, time int, path Path, visited map[string]bool) []Path {
// 	paths := []Path{path}

// 	for _, next := range nonZeroValves {
// 		newTime := time - distances[current][next] - 1
// 		if visited[next] || newTime <= 0 {
// 			continue
// 		}
// 		newMap := copyMap(visited)
// 		newMap[next] = true
// 		newPath := path.copy()
// 		flowRate := valves[next].flowRate
// 		// fmt.Printf("Flowrate for %s is %d\n", next, flowRate)
// 		newPath.addToPath(newTime*flowRate, next)
// 		paths = append(paths, DFS(next, newTime, newPath, newMap)...)
// 	}

// 	return paths
// }

func main() {

	// var valves map[string]Valve = make(map[string]Valve)

	inputLines := readStdin()
	blueprints := make([]Blueprint, 0)

	for i, line := range inputLines {
		bp := getBlueprintForLine(line)
		bp.id = i + 1

		blueprints = append(blueprints, bp)
	}
	fmt.Printf("Blueprints %+v\n", blueprints)
	// for _, bp := range blueprints {
	bp := blueprints[0]
	fmt.Printf("Processing bp %d\n", bp.id)
	path := Path{
		path: make([]Status, 0),
	}
	path.path = append(path.path, getInitialStatus())
	bp.paths = DFS(24, path, bp.robotCost)
	max := 0
	for _, path := range bp.paths {
		if path.path[len(path.path)-1].gems.geode > max {
			max = path.path[len(path.path)-1].gems.geode
		}
	}
	fmt.Printf("Max is %d\n", max)
	// }
}

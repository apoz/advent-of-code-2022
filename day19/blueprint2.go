package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var maxValFound int = 0

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
	path              []Status
	maxOreRobots      int
	maxClayRobots     int
	maxObsidianRobots int
}

type Blueprint struct {
	robotCost         map[string]GemInventory
	id                int
	paths             []Path
	maxOreRobots      int
	maxClayRobots     int
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

	maxOreRobots := max(robotCost["ore"].ore, max(robotCost["clay"].ore, max(robotCost["obsidian"].ore, robotCost["geode"].ore)))
	maxObsidianRobots := max(robotCost["ore"].obsidian, max(robotCost["clay"].obsidian, max(robotCost["obsidian"].obsidian, robotCost["geode"].obsidian)))
	maxClayRobots := max(robotCost["ore"].clay, max(robotCost["clay"].clay, max(robotCost["obsidian"].clay, robotCost["geode"].clay)))
	// fmt.Printf("MaxOreRobots %d\n", maxOreRobots)
	// fmt.Printf("MaxObsidianRobots %d\n", maxObsidianRobots)
	// fmt.Printf("MaxClayRobots %d\n", maxClayRobots)

	paths := make([]Path, 0)
	path := Path{
		path:              make([]Status, 0),
		maxOreRobots:      maxOreRobots,
		maxClayRobots:     maxClayRobots,
		maxObsidianRobots: maxObsidianRobots,
	}
	path.path = append(path.path, getInitialStatus())
	paths = append(paths, path)

	bp := Blueprint{
		robotCost:         robotCost,
		paths:             paths,
		id:                id,
		maxOreRobots:      maxOreRobots,
		maxClayRobots:     maxClayRobots,
		maxObsidianRobots: maxObsidianRobots,
	}

	return bp
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getOptionsForPath(path Path, robotCost map[string]GemInventory) []string {
	operations := make([]string, 0)
	lastStatus := path.path[len(path.path)-1]
	// fmt.Printf("lastStatus %+v\n", lastStatus)
	operations = append(operations, "do_nothing")
	if canBuildRobot(lastStatus, robotCost["geode"]) {
		operations = append(operations, "build_geode_robot")
	}
	if canBuildRobot(lastStatus, robotCost["ore"]) && lastStatus.robots.ore < path.maxOreRobots {
		operations = append(operations, "build_ore_robot")
	}
	if canBuildRobot(lastStatus, robotCost["clay"]) && lastStatus.robots.clay < path.maxClayRobots {
		operations = append(operations, "build_clay_robot")
	}
	if canBuildRobot(lastStatus, robotCost["obsidian"]) && lastStatus.robots.obsidian < path.maxObsidianRobots {
		operations = append(operations, "build_obsidian_robot")
	}

	return operations
}

func applyOperationToPath(operation string, path Path, cost map[string]GemInventory) Path {
	fmt.Printf("\nAAAAAApplying operation %s to path %+v\n\n", operation, path)
	status := path.path[len(path.path)-1]
	newStatus := copyStatus(status)
	if operation == "do_nothing" {
		return path
	} else if operation == "build_ore_robot" {
		fmt.Printf("Executing build_ore_robot current gems clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
		newStatus.gems.clay -= cost["ore"].clay
		newStatus.gems.obsidian -= cost["ore"].obsidian
		newStatus.gems.ore -= cost["ore"].ore
		newStatus.gems.geode -= cost["ore"].geode
		newStatus.robots.ore++
		fmt.Printf("Executing build_ore_robot current gems after clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
	} else if operation == "build_clay_robot" {
		fmt.Printf("Executing build_clay_robot current gems clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
		newStatus.gems.clay -= cost["clay"].clay
		newStatus.gems.obsidian -= cost["clay"].obsidian
		newStatus.gems.ore -= cost["clay"].ore
		newStatus.gems.geode -= cost["clay"].geode
		newStatus.robots.clay++
		fmt.Printf("Executing build_clay_robot current gems after clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
	} else if operation == "build_obsidian_robot" {
		fmt.Printf("Executing build_obsidian_robot current gems clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
		newStatus.gems.clay -= cost["obsidian"].clay
		newStatus.gems.obsidian -= cost["obsidian"].obsidian
		newStatus.gems.ore -= cost["obsidian"].ore
		newStatus.gems.geode -= cost["obsidian"].geode
		newStatus.robots.obsidian++
		fmt.Printf("Executing build_obsidian_robot current gems after clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
	} else if operation == "build_geode_robot" {
		fmt.Printf("Executing build_geode_robot current gems clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
		newStatus.gems.clay -= cost["geode"].clay
		newStatus.gems.obsidian -= cost["geode"].obsidian
		newStatus.gems.ore -= cost["geode"].ore
		newStatus.gems.geode -= cost["geode"].geode
		newStatus.robots.geode++
		fmt.Printf("Executing build_geode_robot current gems after clay=%d obsidian=%d ore=%d geode=%d\n", newStatus.gems.clay, newStatus.gems.obsidian, newStatus.gems.ore, newStatus.gems.geode)
	}
	newPath := Path{
		path:              append(path.path, newStatus),
		maxOreRobots:      path.maxOreRobots,
		maxObsidianRobots: path.maxObsidianRobots,
		maxClayRobots:     path.maxClayRobots,
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
	robots := RobotInventory{}
	robots.ore = status.robots.ore
	robots.clay = status.robots.clay
	robots.obsidian = status.robots.obsidian
	robots.geode = status.robots.geode

	gemInventory := GemInventory{}
	gemInventory.ore = status.gems.ore
	gemInventory.clay = status.gems.clay
	gemInventory.obsidian = status.gems.obsidian
	gemInventory.geode = status.gems.geode

	newStatus := Status{}
	newStatus.robots = robots
	newStatus.gems = gemInventory
	return newStatus
}

func copyPath(path Path) Path {
	statusList := make([]Status, 0)
	for _, status := range path.path {
		newStatus := copyStatus(status)
		statusList = append(statusList, newStatus)
	}
	newPath := Path{
		path:              statusList,
		maxOreRobots:      path.maxOreRobots,
		maxClayRobots:     path.maxClayRobots,
		maxObsidianRobots: path.maxObsidianRobots,
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

func DFS(time int, path Path, robotCost map[string]GemInventory) []Path {
	paths := []Path{path}

	fmt.Printf("\nCalling DFS time=%d with path %+v\n", time, path)
	options := getOptionsForPath(path, robotCost)
	fmt.Printf("The options for the current path are %+v\n", options)
	for _, operation := range options {
		newTime := time - 1
		if time <= 0 {
			if path.path[len(path.path)-1].gems.geode > maxValFound {
				maxValFound = path.path[len(path.path)-1].gems.geode
				fmt.Printf("New max geodes -> %d\n", maxValFound)
				fmt.Printf("Path -> %+v\n", path)
			}
			continue
		}
		newPath := copyPath(path)
		currentStatus := newPath.path[len(newPath.path)-1]
		currentRobots := currentStatus.robots
		fmt.Printf("Applying %s operation to current status %+v\n", operation, currentStatus)
		newPath = applyOperationToPath(operation, path, robotCost)
		fmt.Printf("After operation %+v\n", newPath.path[len(newPath.path)-1])

		newPath.path[len(newPath.path)-1].gems.ore += currentRobots.ore
		newPath.path[len(newPath.path)-1].gems.clay += currentRobots.clay
		newPath.path[len(newPath.path)-1].gems.obsidian += currentRobots.obsidian
		newPath.path[len(newPath.path)-1].gems.geode += currentRobots.geode
		fmt.Printf("\nAfter increase path is %+v\n", newPath.path[len(newPath.path)-1])
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
	// fmt.Printf("Blueprints %+v\n", blueprints)
	// for _, bp := range blueprints {
	bp := blueprints[0]
	fmt.Printf("Processing bp %d %+v\n", bp.id, bp)
	path := Path{
		path:              make([]Status, 0),
		maxOreRobots:      bp.maxOreRobots,
		maxClayRobots:     bp.maxClayRobots,
		maxObsidianRobots: bp.maxObsidianRobots,
	}
	path.path = append(path.path, getInitialStatus())
	bp.paths = DFS(24, path, bp.robotCost)
	maxVal := 0
	for _, path := range bp.paths {
		if path.path[len(path.path)-1].gems.geode > maxVal {
			maxVal = path.path[len(path.path)-1].gems.geode
		}
	}
	fmt.Printf("Max is %d\n", maxVal)
	// }
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// https://github.com/mebeim/aoc/blob/master/2022/README.md#day-19---not-enough-minerals

type Blueprint struct {
	id        int
	robotCost map[string]GemInventory
}

type Status struct {
	time int
	gems GemInventory
	robs RobotInventory
}

func newStatus(time int, g GemInventory, r RobotInventory) Status {
	s := Status{}
	s.time = time
	gI := GemInventory{}
	gI.clay = g.clay
	gI.ore = g.ore
	gI.obsidian = g.obsidian
	gI.geode = g.geode
	s.gems = gI
	rI := RobotInventory{}
	rI.clay = r.clay
	rI.ore = r.ore
	rI.geode = r.geode
	rI.obsidian = r.obsidian
	s.robs = rI
	return s
}

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

type StatusQueue struct {
	statuses []Status
}

func (s *StatusQueue) enqueue(status Status) {
	s.statuses = append(s.statuses, status)
}

func (s *StatusQueue) dequeue() Status {
	if len(s.statuses) > 0 {
		elem := s.statuses[0]
		s.statuses = s.statuses[1:]
		return elem
	}
	return Status{}
}

func (s *StatusQueue) isEmpty() bool {
	if len(s.statuses) == 0 {
		return true
	}
	return false
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getBlueprintForLine(line string) Blueprint {

	r := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	res := r.FindStringSubmatch(line)

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

	bp := Blueprint{
		robotCost: robotCost,
	}

	return bp
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func collectRobotResults(r RobotInventory, g GemInventory) GemInventory {
	newGI := GemInventory{}
	newGI.ore = g.ore + r.ore
	newGI.clay = g.clay + r.clay
	newGI.obsidian = g.obsidian + r.obsidian
	newGI.geode = g.geode + r.geode
	return newGI

}
func initialStatus() Status {
	gI := GemInventory{
		ore:      0,
		clay:     0,
		obsidian: 0,
		geode:    0,
	}
	rI := RobotInventory{
		ore:      1,
		clay:     0,
		obsidian: 0,
		geode:    0,
	}
	s := Status{
		time: 24,
		robs: rI,
		gems: gI,
	}
	return s
}

func canBuildGeodeRobot(g GemInventory, cost GemInventory) bool {
	if g.ore >= cost.ore && g.obsidian >= cost.obsidian {
		return true
	}
	return false
}

func canBuildObsidianRobot(g GemInventory, cost GemInventory) bool {
	if g.ore >= cost.ore && g.clay >= cost.clay {
		return true
	}
	return false
}

func canBuildClayRobot(g GemInventory, cost GemInventory) bool {
	if g.ore >= cost.ore {
		return true
	}
	return false
}

func canBuildOreRobot(g GemInventory, cost GemInventory) bool {
	if g.ore >= cost.ore {
		return true
	}
	return false
}

func search(bp Blueprint) int {

	time := 24
	best := 0
	visited := make(map[Status]bool)
	counter := 0

	max_ore_needed := max(
		bp.robotCost["ore"].ore,
		max(bp.robotCost["obsidian"].ore,
			max(bp.robotCost["geode"].ore, bp.robotCost["clay"].ore)))
	max_clay_needed := bp.robotCost["obsidian"].clay
	max_obs_needed := bp.robotCost["geode"].obsidian
	fmt.Printf("Max_ore_needed=%d max_clay_needed=%d max_obs_needed=%d\n", max_ore_needed, max_clay_needed, max_obs_needed)
	statuses := make([]Status, 0)
	q := &StatusQueue{
		statuses: statuses,
	}
	q.enqueue(initialStatus())
	for q.isEmpty() == false {
		qLen := len(*&q.statuses)
		if (qLen % 1000000) == 0 {
			fmt.Printf("Queue len %d\n", len(*&q.statuses))
		}

		status := q.dequeue()
		time = status.time
		// fmt.Printf("Processing status minute=%d  %+v\n", 24-status.time+1, status)

		_, ok := visited[status]
		if ok { // state was already visited
			// fmt.Println("This status was visited already\n")
			continue
		}
		visited[status] = true

		// Collect robot results
		newGems := collectRobotResults(status.robs, status.gems)
		// fmt.Printf("Right After collect robot results minute=%d gems=%+v\n", 24-time+1, newGems)
		time -= 1

		// If time == 0 we´re in a goal state and check geodes count
		if time == 0 {
			//fmt.Printf("Time <=0 \n")
			counter++
			if newGems.geode > best {
				best = newGems.geode
				fmt.Printf("New max for geodes minute=%d robots: %+v gems: %+v\n", 24-time+1, status.robs, newGems)
			}
			best = max(best, newGems.geode)
			if (counter % 1000000) == 0 {
				fmt.Printf("Tested %d goal status\n", counter)

			}
			continue
		}

		//Options
		newStat := newStatus(time, newGems, status.robs)
		// Do nothing
		if (status.gems.ore < max_ore_needed) || (status.robs.clay > 0 && status.gems.clay < max_clay_needed) || (status.robs.obsidian > 0 && status.gems.obsidian < max_obs_needed) {
			// Do nothing
			// fmt.Printf("Added do_nothing action minute=%d %+v\n", 24-time+1, newStat)
			q.enqueue(newStat)
		}
		// Have enough material for geode robot? build it
		if canBuildGeodeRobot(status.gems, bp.robotCost["geode"]) {
			// fmt.Printf("Build geodes robot minute=%d\n", 24-time+1)
			nS := newStatus(time, newGems, status.robs)
			// fmt.Printf("Status before %+v\n", nS)
			nS.gems.ore -= bp.robotCost["geode"].ore
			nS.gems.obsidian -= bp.robotCost["geode"].obsidian
			nS.robs.geode++
			// fmt.Printf("Status after %+v\n", nS)
			q.enqueue(nS)
		}

		// Have enough material for obsidian robot? build it
		if canBuildObsidianRobot(status.gems, bp.robotCost["obsidian"]) {
			if status.robs.obsidian < max_obs_needed {
				// fmt.Printf("Build obsidian robot minute=%d\n", 24-time+1)
				nS := newStatus(time, newGems, status.robs)
				// fmt.Printf("Status before %+v\n", nS)
				nS.gems.ore -= bp.robotCost["obsidian"].ore
				nS.gems.clay -= bp.robotCost["obsidian"].clay
				nS.robs.obsidian++
				//fmt.Printf("Added build_obs action %+v\n", nS)
				// fmt.Printf("Status after %+v\n", nS)
				q.enqueue(nS)
			}
		}

		// Have enough material for ore robot? build it
		if canBuildOreRobot(status.gems, bp.robotCost["ore"]) {
			if status.robs.ore < max_ore_needed {
				// fmt.Printf("Build ore robot minute=%d\n", 24-time+1)
				nS := newStatus(time, newGems, status.robs)
				// fmt.Printf("Status before %+v\n", nS)
				nS.gems.ore -= bp.robotCost["ore"].ore
				nS.robs.ore++
				// fmt.Printf("Status after %+v\n", nS)
				//fmt.Printf("Added build_ore action %+v\n", nS)
				q.enqueue(nS)
			}
		}

		// Have enough material for clay robot? build it
		if canBuildClayRobot(status.gems, bp.robotCost["clay"]) {
			if status.robs.clay < max_clay_needed {
				// fmt.Printf("Build clay robot minute=%d\n", 24-time+1)
				nS := newStatus(time, newGems, status.robs)
				// fmt.Printf("Status before %+v\n", nS)
				nS.gems.ore -= bp.robotCost["clay"].ore
				nS.robs.clay++
				// fmt.Printf("Status after %+v\n", nS)
				//fmt.Printf("Added build_clay action %+v\n", nS)
				q.enqueue(nS)
			}
		}

	}
	return best

}

func main() {

	// var valves map[string]Valve = make(map[string]Valve)

	inputLines := readStdin()
	blueprints := make([]Blueprint, 0)

	total := 0
	for i, line := range inputLines {
		bp := getBlueprintForLine(line)
		bp.id = i + 1
		blueprints = append(blueprints, bp)
		qaLevel := bp.id * search(bp)
		total += qaLevel
	}

	fmt.Printf("The result is %d\n", total)
}

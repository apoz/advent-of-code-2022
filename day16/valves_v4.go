package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var valves map[string]Valve
var distances map[string]map[string]int
var nonZeroValves []string

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

type Path struct {
	flow    int
	visited []string
}

func newPath() Path {
	return Path{0, []string{}}
}

func (p *Path) addToPath(flow int, valve string) {
	p.flow += flow
	p.visited = append(p.visited, valve)
}

func (p Path) copy() Path {
	vis := make([]string, len(p.visited))
	copy(vis, p.visited)
	return Path{p.flow, vis}
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getValveForLine(line string) Valve {

	r := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
	res := r.FindStringSubmatch(line)
	v := Valve{}
	v.name = res[1]
	v.flowRate, _ = strconv.Atoi(res[2])
	v.tunnels = strings.Split(res[3], ", ")
	return v
}

func getValvesDict(lines []string) map[string]Valve {
	valves := make(map[string]Valve, 0)
	for _, line := range lines {
		valve := getValveForLine(line)
		valves[valve.name] = valve
	}
	return valves
}

func getNonZeroValves(v map[string]Valve) []string {
	nonZero := make([]string, 0)
	for val := range v {
		if v[val].flowRate > 0 {
			nonZero = append(nonZero, val)
		}
	}
	return nonZero
}

func main() {

	// var valves map[string]Valve = make(map[string]Valve)

	inputLines := readStdin()

	valves = getValvesDict(inputLines)
	// fmt.Printf("valves dictionary -> %+v\n", valves)
	nonZeroValves = append(nonZeroValves, getNonZeroValves(valves)...)
	// fmt.Printf("NonZero valves -> %+v\n", nonZeroValves)
	// dat, _ := os.ReadFile("./short-input.txt")
	// rv, err := regexp.Compile(`[A-Z]{2}`)
	// check(err)
	// rf, err := regexp.Compile(`\d+`)
	// check(err)

	// valves = make(map[string]Valve)
	// for _, line := range strings.Split(strings.TrimSpace(string(dat)), "\n") {
	// 	valve := rv.FindAllString(line, -1)
	// 	flow, _ := strconv.Atoi(rf.FindString(line))
	// 	v := Valve{flow, valve[1:]}
	// 	valves[valve[0]] = v
	// 	if flow > 0 {
	// 		non_zero = append(non_zero, valve[0])
	// 	}
	// }

	distances = floydWarshall(valves)

	// fmt.Printf("distances dictionary -> %+v\n", distances)
	// t1 := time.Now()
	mypaths := DFS("AA", 30, newPath(), make(map[string]bool))
	// fmt.Printf("Mypaths: %+v\n", mypaths)
	sort.Slice(mypaths, func(i, j int) bool { return mypaths[i].flow > mypaths[j].flow })
	fmt.Printf("Max flow: %d\n", mypaths[0].flow)
	// fmt.Printf("Time elapsed: %s\n", time.Since(t1))

	// t2 := time.Now()
	fmt.Printf("Second part: %d\n", partTwo())
	// fmt.Printf("Time elapsed: %s\n", time.Since(t2))
}

func partTwo() int {
	p2 := DFS("AA", 26, newPath(), make(map[string]bool))

	max := 0
	for _, a := range p2 {
		if len(a.visited) != 0 {
			m := make(map[string]bool)
			for _, v := range a.visited {
				m[v] = true
			}
			for _, b := range p2 {
				f := a.flow + b.flow
				if f > max && len(b.visited) > 0 && allValvesDifferent(m, b.visited) {
					max = f
				}
			}
		}
	}
	return max
}

func DFS(current string, time int, path Path, visited map[string]bool) []Path {
	paths := []Path{path}

	for _, next := range nonZeroValves {
		newTime := time - distances[current][next] - 1
		if visited[next] || newTime <= 0 {
			continue
		}
		newMap := copyMap(visited)
		newMap[next] = true
		newPath := path.copy()
		flowRate := valves[next].flowRate
		// fmt.Printf("Flowrate for %s is %d\n", next, flowRate)
		newPath.addToPath(newTime*flowRate, next)
		paths = append(paths, DFS(next, newTime, newPath, newMap)...)
	}

	return paths
}

func floydWarshall(valves map[string]Valve) map[string]map[string]int {
	var dist map[string]map[string]int = make(map[string]map[string]int)

	for i := range valves {
		for j := range valves {
			if _, ok := dist[i]; !ok {
				dist[i] = make(map[string]int)
			}
			if i == j {
				dist[i][j] = 0
			} else if contains(valves[i].tunnels, j) {
				dist[i][j] = 1
			} else {
				dist[i][j] = 999999
			}
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func copyMap(m map[string]bool) map[string]bool {
	mcopy := make(map[string]bool)
	for k, v := range m {
		mcopy[k] = v
	}
	return mcopy
}

func allValvesDifferent(m map[string]bool, brr []string) bool {
	for _, v := range brr {
		if m[v] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

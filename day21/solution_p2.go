package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Monkey struct {
	name       string
	value      int
	finalValue bool
	op         string
	monkey1    string
	monkey2    string
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getSingleNumber(line string) (bool, string, int) {
	singleNumberRegex := regexp.MustCompile(`^(\w+): (\d+)$`)
	res := singleNumberRegex.FindStringSubmatch(line)
	if res == nil {
		return false, "", 0
	}
	val, _ := strconv.Atoi(res[2])
	return true, res[1], val
}

func getOperationMonkey(line string) (bool, string, string, string, string) {
	operationRegex := regexp.MustCompile(`^(\w+): (\w+) ([\/*+-]) (\w+)$`)
	res := operationRegex.FindStringSubmatch(line)
	if res == nil {
		return false, "", "", "", ""
	}
	return true, res[1], res[2], res[4], res[3]
}

func evalOperation(op string, m1 int, m2 int) int {
	if op == "+" {
		return m1 + m2
	} else if op == "-" {
		return m1 - m2
	} else if op == "*" {
		return m1 * m2
	} else if op == "/" {
		return m1 / m2
	} else if op == "=" {
		if m1 == m2 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func evaluateDictionary(monkeys map[string]*Monkey) bool {
	somethingChanged := false
	for _, monkey := range monkeys {
		if (*monkey).finalValue == false && (*monkey).op != "" {
			if (*monkeys[(*monkey).monkey1]).finalValue && (*monkeys[(*monkey).monkey2]).finalValue {
				(*monkey).value = evalOperation((*monkey).op, (*monkeys[(*monkey).monkey1]).value, (*monkeys[(*monkey).monkey2]).value)
				(*monkey).finalValue = true
				somethingChanged = true
			}
		}
	}
	return somethingChanged
}

func resetDictFinalStatus(monkeys map[string]*Monkey) {
	for _, mon := range monkeys {
		if (*mon).op != "" {
			(*mon).finalValue = false
		}
	}
}

func getDictFromLines(lines []string) map[string]*Monkey {
	monkeyList := make(map[string]*Monkey, 0)

	for _, line := range lines {
		single_number, monkey_name, value := getSingleNumber(line)
		if single_number == true {
			mon := Monkey{
				name:       monkey_name,
				value:      value,
				finalValue: true,
			}
			if monkey_name == "humn" {
				mon.value = 0
				mon.finalValue = false
			}
			monkeyList[monkey_name] = &mon

		} else {
			isOp, monkey_name, m1, m2, op := getOperationMonkey(line)
			if isOp == false {
				fmt.Println("Something went wrong")
				os.Exit(1)
			}
			mon := Monkey{
				name:       monkey_name,
				value:      0,
				finalValue: false,
				monkey1:    m1,
				monkey2:    m2,
				op:         op,
			}
			if mon.name == "root" {
				mon.op = "="
			}
			monkeyList[monkey_name] = &mon
		}
	}
	changed := true
	for changed {
		changed = evaluateDictionary(monkeyList)
	}
	return monkeyList
}

func findNewRequiredResult(currentRes int, operation string, operand1 int, operand2 int) int {
	if operation == "+" {
		if operand1 > 0 {
			return currentRes - operand1
		} else {
			return currentRes - operand2
		}
	}
	if operation == "*" {
		if operand1 > 0 {
			return currentRes / operand1
		} else {
			return currentRes / operand2
		}
	}
	if operation == "-" {
		if operand1 > 0 {
			return operand1 - currentRes
		} else {
			return currentRes + operand2
		}
	}
	if operation == "/" {
		if operand1 > 0 {
			return operand1 / currentRes
		} else {
			return currentRes * operand2
		}
	}
	return 0

}

func main() {

	// var valves map[string]Valve = make(map[string]Valve)

	inputLines := readStdin()
	monkeys := getDictFromLines(inputLines)

	// changed := true
	// for changed {
	// 	changed = evaluateDictionary(monkeys)
	// }
	fmt.Printf("The dictionary is %+v\n", monkeys)

	found := false
	mon := monkeys["root"]
	res := 0
	if (*monkeys[(*mon).monkey1]).finalValue == false {
		res = (*monkeys[(*mon).monkey2]).value
	} else {
		res = (*monkeys[(*mon).monkey1]).value
	}
	for !found {
		fmt.Printf("Studying Monkey %+v  ", (*mon))
		if (*monkeys[(*mon).monkey1]).finalValue == false {
			fmt.Printf(".The expected result is %d.  ", res)
			fmt.Printf(". The other operand is %d\n", (*monkeys[(*mon).monkey2]).value)
			mon = monkeys[(*mon).monkey1]
			res = findNewRequiredResult(res, (*mon).op, (*monkeys[(*mon).monkey1]).value, (*monkeys[(*mon).monkey2]).value)
		} else if (*monkeys[(*mon).monkey2]).finalValue == false {
			fmt.Printf(".The expected result is %d.  ", res)
			fmt.Printf(". The other operand is %d\n", (*monkeys[(*mon).monkey2]).value)
			mon = monkeys[(*mon).monkey2]
			res = findNewRequiredResult(res, (*mon).op, (*monkeys[(*mon).monkey1]).value, (*monkeys[(*mon).monkey2]).value)
		}
		if mon.name == "humn" {
			fmt.Printf("FOUND HUMN %+v. RESPONSE %d\n", mon, res)
			found = true
			os.Exit(0)
		}
	}
}

package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	roundItems     []int
	operation      string
	secondOperator int
	testDivisible  int
	trueMonkey     int
	falseMonkey    int
	inspectCounter int
}

func (m *Monkey) appendItemRound(item int) {
	m.roundItems = append(m.roundItems, item)
}

func (m *Monkey) calculateRound(monkeys []*Monkey) {
	for len(m.roundItems) > 0 {
		item := 0
		if len(m.roundItems) > 1 {
			item, m.roundItems = m.roundItems[0], m.roundItems[1:]
		} else {
			item, m.roundItems = m.roundItems[0], make([]int, 0)
		}
		fmt.Printf("Evaluating item %d\n", item)
		m.inspect()
		opValue := m.calculateOperation(item)
		opValue = int(opValue) / 3
		testResult := opValue%m.testDivisible == 0
		if testResult {
			fmt.Printf("Adding %d to Monkey %d\n", opValue, m.trueMonkey)
			monkeys[m.trueMonkey].appendItemRound(opValue)
		} else {
			fmt.Printf("Adding %d to Monkey %d\n", opValue, m.falseMonkey)
			monkeys[m.falseMonkey].appendItemRound(opValue)
		}
	}

}

func (m *Monkey) inspect() {
	m.inspectCounter++
}

func (m *Monkey) calculateOperation(item int) int {
	fmt.Printf("calculateOperation item %d\n", item)
	sOp := 0
	if m.secondOperator == 0 {
		sOp = item
	} else {
		sOp = m.secondOperator
	}
	if m.operation == "+" {
		return item + sOp
	} else if m.operation == "*" {
		fmt.Printf("Operation result %d * %d = %d\n", item, sOp, item*sOp)
		return item * sOp
	} else {
		print("SHOULD NOT HAPPEN")
		return 0
	}
}

func main() {
	var Monkeys []*Monkey
	m0Items := make([]int, 0)
	m0Items = append(m0Items, 66)
	m0Items = append(m0Items, 71)
	m0Items = append(m0Items, 94)
	Mon0 := &Monkey{
		roundItems:     m0Items,
		operation:      "*",
		secondOperator: 5,
		testDivisible:  3,
		trueMonkey:     7,
		falseMonkey:    4,
		inspectCounter: 0,
	}
	m1Items := make([]int, 0)
	m1Items = append(m1Items, 70)
	Mon1 := &Monkey{
		roundItems:     m1Items,
		operation:      "+",
		secondOperator: 6,
		testDivisible:  17,
		trueMonkey:     3,
		falseMonkey:    0,
		inspectCounter: 0,
	}
	m2Items := make([]int, 0)
	m2Items = append(m2Items, 62)
	m2Items = append(m2Items, 68)
	m2Items = append(m2Items, 56)
	m2Items = append(m2Items, 65)
	m2Items = append(m2Items, 94)
	m2Items = append(m2Items, 78)
	Mon2 := &Monkey{
		roundItems:     m2Items,
		operation:      "+",
		secondOperator: 5,
		testDivisible:  2,
		trueMonkey:     3,
		falseMonkey:    1,
		inspectCounter: 0,
	}
	m3Items := make([]int, 0)
	m3Items = append(m3Items, 89)
	m3Items = append(m3Items, 94)
	m3Items = append(m3Items, 94)
	m3Items = append(m3Items, 67)
	Mon3 := &Monkey{
		roundItems:     m3Items,
		operation:      "+",
		secondOperator: 2,
		testDivisible:  19,
		trueMonkey:     7,
		falseMonkey:    0,
		inspectCounter: 0,
	}
	m4Items := make([]int, 0)
	m4Items = append(m4Items, 71)
	m4Items = append(m4Items, 61)
	m4Items = append(m4Items, 73)
	m4Items = append(m4Items, 65)
	m4Items = append(m4Items, 98)
	m4Items = append(m4Items, 98)
	m4Items = append(m4Items, 63)
	Mon4 := &Monkey{
		roundItems:     m4Items,
		operation:      "*",
		secondOperator: 7,
		testDivisible:  11,
		trueMonkey:     5,
		falseMonkey:    6,
		inspectCounter: 0,
	}
	m5Items := make([]int, 0)
	m5Items = append(m5Items, 55)
	m5Items = append(m5Items, 62)
	m5Items = append(m5Items, 68)
	m5Items = append(m5Items, 61)
	m5Items = append(m5Items, 60)
	Mon5 := &Monkey{
		roundItems:     m5Items,
		operation:      "+",
		secondOperator: 7,
		testDivisible:  5,
		trueMonkey:     2,
		falseMonkey:    1,
		inspectCounter: 0,
	}
	m6Items := make([]int, 0)
	m6Items = append(m6Items, 93)
	m6Items = append(m6Items, 91)
	m6Items = append(m6Items, 69)
	m6Items = append(m6Items, 64)
	m6Items = append(m6Items, 72)
	m6Items = append(m6Items, 89)
	m6Items = append(m6Items, 50)
	m6Items = append(m6Items, 71)
	Mon6 := &Monkey{
		roundItems:     m6Items,
		operation:      "+",
		secondOperator: 1,
		testDivisible:  13,
		trueMonkey:     5,
		falseMonkey:    2,
		inspectCounter: 0,
	}
	m7Items := make([]int, 0)
	m7Items = append(m7Items, 76)
	m7Items = append(m7Items, 50)
	Mon7 := &Monkey{
		roundItems:     m7Items,
		operation:      "*",
		secondOperator: 0,
		testDivisible:  7,
		trueMonkey:     4,
		falseMonkey:    6,
		inspectCounter: 0,
	}
	Monkeys = append(Monkeys, Mon0)
	Monkeys = append(Monkeys, Mon1)
	Monkeys = append(Monkeys, Mon2)
	Monkeys = append(Monkeys, Mon3)
	Monkeys = append(Monkeys, Mon4)
	Monkeys = append(Monkeys, Mon5)
	Monkeys = append(Monkeys, Mon6)
	Monkeys = append(Monkeys, Mon7)
	// sm0Items := make([]int, 0)
	// sm0Items = append(sm0Items, 79)
	// sm0Items = append(sm0Items, 98)
	// shortMon0 := &Monkey{
	// 	roundItems:     sm0Items,
	// 	operation:      "*",
	// 	secondOperator: 19,
	// 	testDivisible:  23,
	// 	trueMonkey:     2,
	// 	falseMonkey:    3,
	// 	inspectCounter: 0,
	// }
	// sm1Items := make([]int, 0)
	// sm1Items = append(sm1Items, 54)
	// sm1Items = append(sm1Items, 65)
	// sm1Items = append(sm1Items, 75)
	// sm1Items = append(sm1Items, 74)
	// shortMon1 := &Monkey{
	// 	roundItems:     sm1Items,
	// 	operation:      "+",
	// 	secondOperator: 6,
	// 	testDivisible:  19,
	// 	trueMonkey:     2,
	// 	falseMonkey:    0,
	// 	inspectCounter: 0,
	// }
	// sm2Items := make([]int, 0)
	// sm2Items = append(sm2Items, 79)
	// sm2Items = append(sm2Items, 60)
	// sm2Items = append(sm2Items, 97)
	// shortMon2 := &Monkey{
	// 	roundItems:     sm2Items,
	// 	operation:      "*",
	// 	secondOperator: 0,
	// 	testDivisible:  13,
	// 	trueMonkey:     1,
	// 	falseMonkey:    3,
	// 	inspectCounter: 0,
	// }
	// sm3Items := make([]int, 0)
	// sm3Items = append(sm3Items, 74)
	// shortMon3 := &Monkey{
	// 	roundItems:     sm3Items,
	// 	operation:      "+",
	// 	secondOperator: 3,
	// 	testDivisible:  17,
	// 	trueMonkey:     0,
	// 	falseMonkey:    1,
	// 	inspectCounter: 0,
	// }
	// Monkeys = append(Monkeys, shortMon0)
	// Monkeys = append(Monkeys, shortMon1)
	// Monkeys = append(Monkeys, shortMon2)
	// Monkeys = append(Monkeys, shortMon3)

	for rounds := 0; rounds < 20; rounds++ {
		for k, mon := range Monkeys {
			fmt.Println("Evaluating round for Monkey ", k)
			mon.calculateRound(Monkeys)
		}
		println("After Round ", rounds)
		for i, mon := range Monkeys {
			fmt.Printf("Monkey %d: %+v\n", i, mon)
		}
	}

	resMonkeys := make([]int, 0)
	for i, mon := range Monkeys {
		fmt.Println("Monkey ", i, " ->", mon.inspectCounter)
		resMonkeys = append(resMonkeys, mon.inspectCounter)
	}
	sort.Ints(resMonkeys)
	fmt.Printf("RESULT->%d\n", resMonkeys[len(resMonkeys)-1]*resMonkeys[len(resMonkeys)-2])
	// inputLines := readStdin()
	// result := trackMovements(inputLines)
	// fmt.Printf("Result -> %d\n", result)
}

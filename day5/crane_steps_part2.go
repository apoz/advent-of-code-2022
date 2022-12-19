package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type stack interface {
	getStackTop() rune
	addRuneToTop(rune)
	removeStackTop() rune
	getStack() []rune
	removeElementsFromTop(int) []rune
	addElementsToTop([]rune)
}

type myStack struct {
	pile string
}

func (x *myStack) getStackTop() rune {
	runeList := []rune(x.pile)
	return runeList[len(runeList)-1]
}

func (x *myStack) addRuneToTop(y rune) {
	runeList := []rune(x.pile)
	runeList = append(runeList, y)
	x.pile = string(runeList)
}

func (x *myStack) removeStackTop() rune {
	runeList := []rune(x.pile)
	lastIndex := len(runeList) - 1
	last := runeList[lastIndex]
	runeList = runeList[:lastIndex]
	x.pile = string(runeList)
	return last
}

func (x *myStack) getStack() []rune {
	return []rune(x.pile)
}

func (x *myStack) removeElementsFromTop(num int) []rune {
	runeList := []rune(x.pile)
	lastIndex := len(runeList) - num
	newRuneList := runeList[lastIndex:]
	runeList = runeList[:lastIndex]
	x.pile = string(runeList)
	return newRuneList
}

func (x *myStack) addElementsToTop(elemList []rune) {
	runeList := []rune(x.pile)
	for _, elem := range elemList {
		runeList = append(runeList, elem)
	}
	x.pile = string(runeList)
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getValuesFromLine(line string) (int, int, int) {
	raws := strings.Split(line, " ")
	amount, _ := strconv.Atoi(raws[1])
	origin, _ := strconv.Atoi(raws[3])
	destination, _ := strconv.Atoi(raws[5])
	return amount, origin, destination
}

func printStatus(stacks []stack) {
	println("------------------------")
	for i := 0; i < len(stacks); i++ {
		println("Stack ", i, " ->", string(stacks[i].getStack()))
	}
	println("------------------------")
}

func main() {
	var stacks []stack = make([]stack, 0)
	stack1 := &myStack{
		pile: "NDMQBPZ",
	}
	stack2 := &myStack{
		pile: "CLZQMDHV",
	}
	stack3 := &myStack{
		pile: "QHRDVFZG",
	}
	stack4 := &myStack{
		pile: "HGDFN",
	}
	stack5 := &myStack{
		pile: "NFQ",
	}
	stack6 := &myStack{
		pile: "DQVZFBT",
	}
	stack7 := &myStack{
		pile: "QMTZDVSH",
	}
	stack8 := &myStack{
		pile: "MGFPNQ",
	}
	stack9 := &myStack{
		pile: "BWRM",
	}
	stacks = append(stacks, stack1)
	stacks = append(stacks, stack2)
	stacks = append(stacks, stack3)
	stacks = append(stacks, stack4)
	stacks = append(stacks, stack5)
	stacks = append(stacks, stack6)
	stacks = append(stacks, stack7)
	stacks = append(stacks, stack8)
	stacks = append(stacks, stack9)

	println("Len stacks ", len(stacks))
	inputLines := readStdin()
	for _, line := range inputLines {
		amount, origin, destination := getValuesFromLine(line)
		// printStatus(stacks)
		// println("Amount ", amount, "Origin ", origin, "Destination ", destination)
		elems := stacks[origin-1].removeElementsFromTop(amount)
		stacks[destination-1].addElementsToTop(elems)
	}
	var finalTops []rune = make([]rune, 0)
	for j := 0; j < len(stacks); j++ {
		finalTops = append(finalTops, stacks[j].getStackTop())
	}
	println("TOPS AT THE END ", string(finalTops))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getListFromLines(lines []string) []*Node {
	nodeList := make([]*Node, 0)

	val, err := strconv.Atoi(lines[0])
	if err != nil {
		os.Exit(0)
	}
	firstNode := &Node{}
	(*firstNode).value = val
	nodeList = append(nodeList, firstNode)
	previousNode := firstNode
	var newNode *Node
	for _, line := range lines[1:] {
		val, err := strconv.Atoi(line)
		if err != nil {
			os.Exit(0)
		}
		newNode = &Node{}
		(*newNode).value = val
		(*newNode).prev = previousNode
		(*previousNode).next = newNode
		nodeList = append(nodeList, newNode)
		previousNode = newNode
	}
	(*newNode).next = firstNode
	(*firstNode).prev = newNode
	return nodeList
}

func moveNodeRight(node *Node) {
	//fmt.Printf("Moving node %s right\n", strconv.Itoa((*node).value))
	//fmt.Printf("LIST BEFORE\n")
	//printList(node, 7)
	previous := (*node).prev
	next := (*node).next

	(*previous).next = next

	(*next).prev = previous

	(*node).prev = next
	(*node).next = (*next).next

	(*next).next.prev = node

	(*next).next = node
	//fmt.Printf("LIST AFTER\n")
	//printList(node, 7)
}

func moveNodeLeft(node *Node) {
	//fmt.Printf("Moving node %s left\n", strconv.Itoa((*node).value))
	//fmt.Printf("LIST BEFORE\n")
	//printList(node, 7)
	previous := (*node).prev
	next := (*node).next

	(*node).next = previous
	(*node).prev = (*previous).prev

	(*previous).prev.next = node

	(*previous).next = next
	(*previous).prev = node

	(*next).prev = previous
	//fmt.Printf("LIST AFTER\n")
	//printList(node, 7)
}

func decrypt(nodeList []*Node, decryptionKey int) {
	listLength := len(nodeList)
	for _, node := range nodeList {
		(*node).value = (*node).value * decryptionKey
	}
	for i := 0; i < 10; i++ {
		for _, node := range nodeList {
			if (*node).value < 0 {
				for i := 0; i < absolute((*node).value, listLength); i++ {
					moveNodeLeft(node)
				}
			} else {
				for i := 0; i < absolute((*node).value, listLength); i++ {
					moveNodeRight(node)
				}
			}
			//fmt.Printf("After processing %s list is\n", strconv.Itoa((*node).value))
			//printList(node, len(nodeList))
		}
	}
}

func absolute(val int, length int) int {
	modVal := val % (length - 1)
	if modVal < 0 {
		return -1 * modVal
	}
	return modVal
}

func calculateOrder(node *Node, value int, order int) int {
	for {
		if (*node).value == value {
			break
		} else {
			node = (*node).next
		}
	}
	fmt.Printf("Found %s value\n", strconv.Itoa(value))
	for i := 0; i < order; i++ {
		node = (*node).next
	}
	return (*node).value
}

func printList(node *Node, length int) {
	for i := 0; i < length; i++ {
		fmt.Printf("Node value %s  Node next Value %s  Node prev value %s \n", strconv.Itoa((*node).value), strconv.Itoa((*(*node).next).value), strconv.Itoa((*(*node).prev).value))
		node = node.next
	}
	fmt.Printf("\n")
}

func main() {

	// var valves map[string]Valve = make(map[string]Valve)
	decryptionKey := 811589153
	inputLines := readStdin()
	nodeList := getListFromLines(inputLines)
	node := nodeList[0]
	fmt.Printf("Original list\n")
	printList(nodeList[0], len(nodeList))
	decrypt(nodeList, decryptionKey)

	_1000Value := calculateOrder(node, 0, 1000)
	_2000Value := calculateOrder(node, 0, 2000)
	_3000Value := calculateOrder(node, 0, 3000)
	fmt.Printf("_1000_value := %s\n", strconv.Itoa(_1000Value))
	fmt.Printf("_2000_value := %s\n", strconv.Itoa(_2000Value))
	fmt.Printf("_3000_value := %s\n", strconv.Itoa(_3000Value))
	fmt.Printf("Response := %s\n", strconv.Itoa(_1000Value+_2000Value+_3000Value))
}

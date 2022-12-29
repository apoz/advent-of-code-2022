package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func getPacket(line string) []interface{} {
	var myPackage []interface{}
	e := json.Unmarshal([]byte(line), &myPackage)
	if e != nil {
		fmt.Println("Error marshalling: ", e)
	}
	//fmt.Printf("Element %+v\n", myPackage)
	return myPackage
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func comparison(i, int) bool {
// 	res := elemsInRightOrder(items[i], items[j])
// 	if res == "right" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func elemsInRightOrder(left interface{}, right interface{}) string {
	leftType := fmt.Sprintf("%T", left)
	rightType := fmt.Sprintf("%T", right)
	// fmt.Printf("elemsInRightOrder left %+v leftType %s right %+v rightType %s\n", left, leftType, right, rightType)
	if leftType == "float64" && rightType == "float64" { // 2 integers
		if left.(float64) < right.(float64) {
			return "right"
		} else if left.(float64) > right.(float64) {
			return "wrong"
		} else {
			return "continue"
		}
	} else if leftType == "float64" && (rightType == "[]interface {}" || rightType == "[]float64") {
		var leftList []interface{} = make([]interface{}, 0)
		leftList = append(leftList, left)
		return elemsInRightOrder(leftList, right)

	} else if (leftType == "[]float64" || leftType == "[]interface {}") && rightType == "float64" {
		var rightList []interface{} = make([]interface{}, 0)
		rightList = append(rightList, right)
		return elemsInRightOrder(left, rightList)

	} else if leftType == "[]interface {}" && rightType == "[]interface {}" {
		for i := 0; i < min(len(left.([]interface{})), len(right.([]interface{}))); i++ {
			res := elemsInRightOrder(left.([]interface{})[i], right.([]interface{})[i])
			if res == "continue" {
				continue
			} else {
				return res
			}
		}
		if len(left.([]interface{})) > len(right.([]interface{})) { // right ran out of elems
			return "wrong"
		} else if len(left.([]interface{})) < len(right.([]interface{})) { // left ran out of elems
			return "right"
		} else {
			return "continue"
		}

	} else if leftType == "[]interface {}" && rightType == "[]float64" {
		for i := 0; i < min(len(left.([]interface{})), len(right.([]interface{}))); i++ {
			res := elemsInRightOrder(left.([]interface{})[i], right.([]interface{})[i])
			if res == "continue" {
				continue
			} else {
				return res
			}
		}
		if len(left.([]interface{})) > len(right.([]interface{})) { // right ran out of elems
			return "wrong"
		} else if len(left.([]interface{})) < len(right.([]interface{})) { // left ran out of elems
			return "right"
		} else {
			return "continue"
		}

	}
	return "UNKOWN"
}

func findDivider(items []interface{}, divider interface{}) int {
	divString := fmt.Sprintf("%+v", divider)
	for i := 0; i < len(items); i++ {
		aString := fmt.Sprintf("%+v", items[i])
		if aString == divString {
			return i + 1
		}
	}
	return -1
}

func main() {
	inputLines := readStdin()
	var items []interface{} = make([]interface{}, 0)
	// var result []int = make([]int, 0)
	for i := 0; i < len(inputLines); i++ {
		item := getPacket(inputLines[i])
		items = append(items, item)
	}
	dividerA := getPacket("[[2]]")
	items = append(items, dividerA)
	dividerB := getPacket("[[6]]")
	items = append(items, dividerB)

	sort.Slice(items, func(i, j int) bool {
		res := elemsInRightOrder(items[i], items[j])
		if res == "right" {
			return true
		} else {
			return false
		}
	})

	dividerAPos := findDivider(items, dividerA)
	dividerBPos := findDivider(items, dividerB)

	fmt.Printf("Response %d\n", dividerAPos*dividerBPos)
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

func elemsInRightOrder(left interface{}, right interface{}) string {
	leftType := fmt.Sprintf("%T", left)
	rightType := fmt.Sprintf("%T", right)
	fmt.Printf("elemsInRightOrder left %+v leftType %s right %+v rightType %s\n", left, leftType, right, rightType)
	if leftType == "float64" && rightType == "float64" { // 2 integers
		println("A 2 float64!")
		if left.(float64) < right.(float64) {
			return "right"
		} else if left.(float64) > right.(float64) {
			return "wrong"
		} else {
			return "continue"
		}
	} else if leftType == "float64" && (rightType == "[]interface {}" || rightType == "[]float64") {
		println("B left float64 right list!")
		var leftList []interface{} = make([]interface{}, 0)
		leftList = append(leftList, left)
		return elemsInRightOrder(leftList, right)

	} else if (leftType == "[]float64" || leftType == "[]interface {}") && rightType == "float64" {
		println("C left list right float64!")
		var rightList []interface{} = make([]interface{}, 0)
		rightList = append(rightList, right)
		return elemsInRightOrder(left, rightList)

	} else if leftType == "[]interface {}" && rightType == "[]interface {}" {
		println("D left list interface{} right list interface {}!")
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
		println("D left list interface{} right list interface {}!")
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

func main() {
	inputLines := readStdin()
	var results []bool = make([]bool, 0)
	// var result []int = make([]int, 0)
	for i := 0; i < len(inputLines); i += 3 {
		left := getPacket(inputLines[i])
		right := getPacket(inputLines[i+1])

		fmt.Printf("Left elem %+v\n", left)
		fmt.Printf("Right elem %+v\n", right)
		resp := elemsInRightOrder(left, right)
		if resp == "right" {
			results = append(results, true)
		} else if resp == "wrong" {
			results = append(results, false)
		} else {
			fmt.Println("UNKNOWNNNNNNNNN")
		}
	}
	total := 0
	for i, val := range results {
		if val {
			total = total + (i + 1)
		}
	}
	fmt.Printf("RESULT-> %d\n", total)

}

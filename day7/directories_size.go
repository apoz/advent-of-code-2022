package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStdin() []string {
	var lines []string = []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

var smallerSizes []int = make([]int, 0)

type file struct {
	Name string
	Size int
}

type directory struct {
	Name            string
	Path            string
	ParentDirectory *directory
	Files           []*file
	Dirs            []*directory
}

func (x *directory) totalSize() int {
	var totalSize int = 0
	for _, fil := range x.Files {
		totalSize += fil.Size
	}
	for _, dir := range x.Dirs {
		totalSize += dir.totalSize()
	}
	return totalSize
}

func applyLine(line string, currDir *directory) *directory {
	// fmt.Printf("Apllying Line %s\n", line)
	// fmt.Printf("CurrentDir %+v\n", *currDir)

	cmdParts := strings.Split(line, " ")
	if cmdParts[0] == "$" { //its a command
		if cmdParts[1] == "cd" {
			if cmdParts[2] == ".." {
				return currDir.ParentDirectory
			} else if cmdParts[2] == "/" {
				for currDir.ParentDirectory != nil {
					currDir = currDir.ParentDirectory
				}
				return currDir
			} else { //changing current dir
				for _, dir := range currDir.Dirs {
					if dir.Name == cmdParts[2] {
						return dir
					}
				}
				print("SHOULD NOT HAPPEN1")
				return currDir
			}
		} else if cmdParts[1] == "ls" {
			return currDir
		}

	} else {
		if cmdParts[0] == "dir" { // its a dir
			newDir := &directory{
				Name:            cmdParts[1],
				Path:            currDir.Path + currDir.Name + "/",
				ParentDirectory: currDir,
				Files:           make([]*file, 0),
				Dirs:            make([]*directory, 0),
			}
			fmt.Printf("Adding new DIR %+v\n", *newDir)
			currDir.Dirs = append(currDir.Dirs, newDir)
			return currDir
		} else { // its a file
			filename := cmdParts[1]
			size, _ := strconv.Atoi(cmdParts[0])
			newFile := &file{
				Name: filename,
				Size: size,
			}
			currDir.Files = append(currDir.Files, newFile)
			return currDir
		}
	}
	println("SHOULD NOT HAPPEN! ", line)
	return currDir
}

func applyLines(lines []string) *directory {
	rootDir := &directory{
		Name:            "",
		Path:            "",
		ParentDirectory: nil,
		Files:           make([]*file, 0),
		Dirs:            make([]*directory, 0),
	}

	currentDir := rootDir
	for _, line := range lines {
		currentDir = applyLine(line, currentDir)
	}
	return rootDir
}

var allDirs map[string]*directory = make(map[string]*directory)

func getAllDirs(rootDir *directory) {
	allDirs[rootDir.Path+"/"+rootDir.Name] = rootDir
	for _, dir := range rootDir.Dirs {
		getAllDirs(dir)
	}

	// for k, v := range allDirs {
	// 	fmt.Printf("Key ->%s\n", k)
	// 	fmt.Printf("Value ->%+v\n", *v)
	// }
}

func sumAllDirsSmallerThan(rootDir *directory, sizeLimit int) int {
	var totalSize int = 0
	getAllDirs(rootDir)
	for _, dir := range allDirs {
		dirSize := dir.totalSize()
		if dirSize <= sizeLimit {
			println(dir.Path, "/", dir.Name, " is smaller than ", sizeLimit, " (", dirSize, ")")
			totalSize += dirSize
		}
	}
	return totalSize
}

func main() {
	inputLines := readStdin()
	rootDir := applyLines(inputLines)
	totalSize := sumAllDirsSmallerThan(rootDir, 100000)
	println("TotalSize ->", totalSize)

}

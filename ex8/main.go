package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"

	"strings"

	"aggelos.com/go/aoc/ex8/nodeUtils"
)


func main() {
	data := readData()
	nodeMap := formatData(data)

	resultPart1 := solverPart1(nodeMap)
	fmt.Println("The result (PART1) : ", resultPart1)

	/*
	start := time.Now()
	res := nodeUtils.LCM([]int {20221, 16343, 13019, 14681, 21883, 16897})
	elapsed := time.Since(start)
	fmt.Printf("Res <%d> -- Time took %s\n", res, elapsed)
	*/

	resultPart2 := solverPart2(nodeMap)
	fmt.Println("The result (PART2) : ", resultPart2)
	if resultPart1 == 14681 && resultPart2 == 14321394058031 {
		fmt.Println(" N I C E ")
	} else {
		fmt.Println(" B A D ")
	}
}


func getRootPath() string {
	path, _ := os.Getwd()
	return path + "\\data\\input.txt"
}


func readData() string {

	file, err := os.Open(getRootPath())
	if err != nil {
		fmt.Println("Error opening the file : ", err)
	}

	scanner := bufio.NewScanner(file)

	data := ""
	for scanner.Scan() {
		line := scanner.Text()
		data += line + "\n"
	}
	return data
}


func formatData(text string) nodeUtils.NodeMap {
	lines := strings.Split(text, "\n")
	nodesMap := nodeUtils.NodeMap {Instructions: lines[0], Paths: map[string]nodeUtils.Node{}}
	for idx := 1; idx < len(lines); idx++ {
		if lines[idx] != "" {
			source, node := formatLine(lines[idx])
			nodesMap.Paths[source] = node
		}
	}
	return nodesMap
}


func formatLine(text string) (string, nodeUtils.Node) {
	parts := strings.Split(text, " = ")
	nodeParts := strings.Split(strings.TrimSpace(parts[1]), ",")
	newNode := nodeUtils.Node{
		Left: strings.TrimSpace(nodeParts[0][1:]),
		Right: strings.TrimSpace(nodeParts[1][:len(nodeParts[1])-1])}
	return strings.TrimSpace(parts[0]), newNode
}


func solverPart1(nodeMap nodeUtils.NodeMap) int {
	return nodeMap.WalkMap("AAA", false)
}


func solverPart2(nodeMap nodeUtils.NodeMap) int {
	return nodeMap.WalkMapPart2(true)
}
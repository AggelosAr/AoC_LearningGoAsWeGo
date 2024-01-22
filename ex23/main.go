package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
	"time"

	utils "aggelos.com/go/aoc/ex23/pointUtils"
)


func main() {
	grid := readData()
	
	start := time.Now()
	resultPart1 := solve1(grid)
	elapsed := time.Since(start)	
	fmt.Println("The result (PART1) : ", resultPart1)
	fmt.Printf("Time took %s\n", elapsed)

	start = time.Now()
	resultPart2 := solve2(grid)
	elapsed = time.Since(start)
	fmt.Println("The result (PART2) : ", resultPart2)
	fmt.Printf("Time took %s\n", elapsed)

	if resultPart1 == 2110 && resultPart2 == 0 {
		fmt.Println(" N I C E ")
	} else {
		fmt.Println(" B A D ")
	}
}


func getRootPath() string {
	path, _ := os.Getwd()
	return path + "\\data\\input.txt"
}


func readData() [][]string {

	file, err := os.Open(getRootPath())

	if err != nil {
		fmt.Println("Error opening the file : ", err)
	}

	scanner := bufio.NewScanner(file)
	
	grid := [][]string {}

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}


func getStartingPosition(grid [][]string) utils.Point {
	startingPoint := utils.Point {}

	for idx, el := range grid[0] {
		if el == "." {
			startingPoint = utils.Point{X:0, Y: idx}
		}
	}

	return startingPoint
}


func solve1(grid [][]string) int {

	maxDistance := 0
	sPoint := getStartingPosition(grid)
	visitedMap := map[utils.Point]struct{} {}
	maxDistance = max(maxDistance, travelDFS(sPoint, grid, visitedMap))

	return maxDistance
}



func travelDFS(point utils.Point, grid [][]string, visitedMap map[utils.Point]struct{}) int {
	
	rows := len(grid)
	cols := len(grid[0])
	// If we reached last row END it 
	// Edge case : We may need to travel to full left or right in that case
	// In this test cases we don't
	if point.X == rows - 1 {
		return len(visitedMap)
	}

	visitedMap[point] = struct{}{}
	nextDirs := utils.AvailableDirections[grid[point.X][point.Y]]
	distance := 0

	for _, direction := range nextDirs {

		offset := utils.DirectionsToOffset[direction]

		nextPoint := point.Add(offset)
		if nextPoint.Validate(rows, cols) != nil {
			continue
		}
		if grid[nextPoint.X][nextPoint.Y] == "#" {
			continue
		}
		_, isVisited := visitedMap[nextPoint]
		if isVisited {
			continue
		}
		
		distance = max(distance, travelDFS(nextPoint, grid, visitedMap))
	}

	delete(visitedMap, point) // backtrack
	return distance
}


func findTargetPoint(grid [][]string) utils.Point {
	targetPoint := utils.Point{}

	for idx, el := range grid[len(grid) - 1] {
		if el == "." {
			targetPoint.X = len(grid) - 1
			targetPoint.Y = idx
		}
	}
	return targetPoint
}


func solve2(grid [][]string) int {
	
	sPoint := getStartingPosition(grid)
	graph := utils.GetGraphFromGrid(grid, sPoint)
	
	maxDistance := 0
	target := findTargetPoint(grid)

	visitedMap := map[utils.Point]struct{} {}
	maxDistance = max(maxDistance, graph.Dfs(sPoint, target, visitedMap))
	
	return maxDistance
}


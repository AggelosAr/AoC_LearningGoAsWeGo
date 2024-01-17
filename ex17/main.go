package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex17/utils"
)


func main() {

	grid := getData()

	start := time.Now()
	part1result := solvePart1(grid)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)

	start = time.Now()
	part2result := solvePart2(grid)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 967 && part2result == 1101 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
	
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() [][]int {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := [][]int {}

	for scanner.Scan() {
		line := scanner.Text()

		tempLine := []int {}
		for _, el := range strings.Split(line, "") {
			number, _ := strconv.Atoi(el)
			tempLine = append(tempLine, number)
		}
		grid = append(grid, tempLine)

	}

	return grid
}
	


func solvePart1(grid [][]int) int {
	return findAll(grid, 1, 3)
}

func solvePart2(grid [][]int) int {
	return findAll(grid, 4, 10)
}


func findAll(grid [][]int, minMoves, maxMoves int ) int {
	
	rows := len(grid)
	cols := len(grid[0])

	queue := utils.PQ[utils.State]{}
	queue.GPush(utils.State{History: "R", Streak: 0}, 0)
	queue.GPush(utils.State{History: "D", Streak: 0}, 0)
	
	visitedStates := map[utils.State]struct{} {}
	
	for len(queue) > 0 {

		state, cost := queue.GPop()
		if state.Point.IsTarget(rows, cols) {
			if state.Streak >= minMoves{
				return cost
			}
		}

		_, seenState := visitedStates[state]
		if seenState {
			continue
		}
		visitedStates[state] = struct{} {}
		
		for _, newState := range state.GetNextStates(minMoves, maxMoves, rows, cols) {
			
			_, seenState := visitedStates[newState]
			if seenState {
				continue
			}
			queue.GPush(newState, grid[newState.Point.X][newState.Point.Y] + cost)	
		}
		
	}
	return 0
}


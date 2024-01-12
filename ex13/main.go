package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex13/utils"
)


func main() {

	records := getData()
	
	start := time.Now()
	part1result := solvePart1(records)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)

	start = time.Now()
	part2result := solvePart2(records)
	elapsed = time.Since(start)
	
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 30487 && part2result == 31954 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []utils.Grid {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grids := []utils.Grid {}
	grid := [][]string {}

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			grids = append(grids, utils.Grid{Data: grid, Rows: len(grid), Cols: len(grid[0])})
			grid = [][]string {}
		} else {
			grid = append(grid, strings.Split(line, ""))
		}
	}
	// no new line at end  * don't forget the last one
	grids = append(grids, utils.Grid{Data: grid, Rows: len(grid), Cols: len(grid[0])})
	return grids
}


func solvePart1(grids []utils.Grid) int {
	result := 0
	
	for _, grid := range grids {
		mid1, _, reflectionType := grid.FindMirrorLine(false)
		// ! ! We suppose it must be 1 of the two cases ! !
		switch reflectionType {
			// vertical
			// number of columns left (it is the idx of the lowest part)
			case "Vertical":
				// 1 indexed count
				result += (mid1 + 1)
			// horizontal
			// + 100 * ( number of rows above )
			case "Horizontal":
				result += 100 * (mid1 + 1)
		}
	}

	return result
}


func solvePart2(grids []utils.Grid) int {
	result := 0
	
	for _, grid := range grids {
		// prevMid1 used to not count same reflection as original
		mid1, _, reflectionType := grid.FindMirrorLine(true)

		switch reflectionType {
			case "Vertical":
				// 1 indexed count
				result += (mid1 + 1)
			case "Horizontal":
				result += 100 * (mid1 + 1)
		}
	}

	return result
}


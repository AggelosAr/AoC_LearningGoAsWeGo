package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex20/utils"
)


func main() {


	
	grid := getData()
	
	start := time.Now()
	part1result := solvePart1(grid)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)
	
	
	grid = getData()
	start = time.Now()
	part2result := solvePart2()
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	

	if part1result == 3820 && part2result == 0 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() [][]string { 

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := [][]string {}

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}


func formatLine(line string) (string, []string) {
	parts := strings.Split(line, "->")
	source := strip(parts[0])
	destinationsMustStrip := strings.Split(strip(parts[1]), ",")
	destinations := []string {}
	for _, dest := range destinationsMustStrip {
		destinations = append(destinations, strip(dest))
	}
	return source, destinations
}


func strip(text string) string {
	return strings.TrimRight(strings.TrimLeft(text, " "), " ")
}


func solvePart1(grid [][]string) int {
	return runBFS(grid)
}


func solvePart2() int {
	
	res_6 := runBFSINF(500)
	res_6++

	
	

	

	return 0
}
	



func printData(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("______________________")
}






func runBFS(grid [][]string) int {
	

	steps := 0
	maxSteps := 64 + 1
	gardenPlots := 1


	startingPoint := getStartingPoint(grid)
	grid[startingPoint.X][startingPoint.Y] = "."

	rows := len(grid)
	cols := len(grid[0])

	q := utils.NewPointQ()
	q.Add(startingPoint)

	visitedSet := map[utils.Point]struct{} {}
	
	for q.HasNext() && steps < maxSteps {

		nextQ := utils.NewPointQ()	
		currentGardenPlots := 0

		for q.HasNext() {
			point, err := q.PopLeft()
			if err != nil {
				break
			}

			_, isVisited := visitedSet[point]
			if isVisited {
				continue
			}

			if grid[point.X][point.Y] == "." {
				currentGardenPlots++
			}

			
			
			visitedSet[point] = struct{} {}
			
			for _, offset := range utils.DirectionsToOffset {
				nextPoint := point.Add(offset)
				if nextPoint.Validate(rows, cols) != nil {
					continue
				}
				if grid[nextPoint.X][nextPoint.Y] == "#" {
					continue
				}
				
				nextQ.Add(nextPoint)
			}
		}

		//reset the visitedMap 
		visitedSet = map[utils.Point]struct{} {}
		gardenPlots = currentGardenPlots
		steps++
		q.List = nextQ.List
	}

	//printData(grid)
	return gardenPlots
}



func runBFSINF(maxSteps int) int {

	originalGrid := getData()
	grid := expandGrid(originalGrid)
	startingPoint := getStartingPointOnExpanded(grid)
	
	steps := 0
	maxSteps += 1
	gardenPlots := 0

	rows := len(grid)
	cols := len(grid[0])

	q := utils.NewPointQ()
	q.Add(startingPoint)

	visitedSet := map[utils.Point]struct{} {}
	
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 0
	}
	defer file.Close()
	

	for q.HasNext() && steps < maxSteps {
		
		nextQ := utils.NewPointQ()
		currentGardenPlots := 0
		

		for q.HasNext() {
			point, _ := q.PopLeft()
			
			
			_, isVisited := visitedSet[point]
			if isVisited {
				continue
			}	
			visitedSet[point] = struct{} {}

			currentGardenPlots++

			for _, offset := range utils.DirectionsToOffset {

				nextPoint := point.Add(offset)
				
				if nextPoint.Validate(rows, cols) != nil {

					inRangeX, inRangeY := Unwrap(rows, cols, nextPoint.X, nextPoint.Y)
					if grid[inRangeX][inRangeY] == "#" {
						continue
					}
				} else {
					if grid[nextPoint.X][nextPoint.Y] == "#" {
						continue
					}
				}
				
				nextQ.Add(nextPoint)
			}
		}
		
		visitedSet = map[utils.Point]struct{} {}
		gardenPlots = currentGardenPlots
		
		//fmt.Printf("STEP <%d> -> CELLS <%d>\n", steps, currentGardenPlots)
		line := fmt.Sprintf("STEP <%d> -> CELLS <%d>\n", steps, currentGardenPlots)
		file.WriteString(line)
		steps++
		q.List = nextQ.List
	}

	return gardenPlots
}




func Unwrap(rows, cols, current_X, current_Y int) (int, int) {
    current_X = (current_X % rows + rows ) % rows
    current_Y = (current_Y % cols + cols ) % cols
    return current_X, current_Y
}


func expandGrid(originalGrid [][]string) [][]string {
	newGrid := [][]string {}

	for _, row := range originalGrid {
		repeatedRow := []string {}
		for i := 0; i < 3; i++ {
			repeatedRow = append(repeatedRow, row...)
		}
		newGrid = append(newGrid, repeatedRow)
	}


	expandedGrid := [][]string {}

	for i := 0; i < 3 ; i++ {
		for _, row := range newGrid {
			expandedGrid = append(expandedGrid, row)
		}
	}

	
	return expandedGrid
}


func getStartingPoint(grid [][]string) utils.Point {
	startingPoint := utils.Point{}
	for i, row := range grid {
		for j, char := range row {
			if char == "S" {
				startingPoint.X = i
				startingPoint.Y = j
			}
		}
	}
	return startingPoint
}



func getStartingPointOnExpanded(grid [][]string) utils.Point {
	count := 0
	startingPoint := utils.Point{}
	for i, row := range grid {
		for j, char := range row {
			if count == 4 {
				startingPoint.X = i
				startingPoint.Y = j
			}
			if char == "S" {
				count++
				grid[startingPoint.X][startingPoint.Y] = "."
				
			}
			
		}
	}
	return startingPoint
}
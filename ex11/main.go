package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"strings"

	"aggelos.com/go/aoc/ex11/pointUtils"
)

func main() {

	grid := getData()
	expandedGrid := expandGrid(grid)
	

	start := time.Now()
	part1result := solvePart1(expandedGrid)
	elapsed := time.Since(start)

	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)


	start = time.Now()
	//part2result := solvePart2(grid, 2)
	part2result := solvePart2(grid, 1000000)
	elapsed = time.Since(start)

	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 9550717 && part2result == 648458253817 {
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

	data := [][]string {}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, ""))	
	}

	return data

}

func printGrid(grid [][]string ) {
	for _, row := range grid {
		fmt.Println(row)
	}
}


func expandGrid(grid [][]string) [][]string{
	
	emptyRows, emptyCols := getEmptyLines(grid)

	expandedRows := [][]string {}
	for i, row := range grid {
		expandedRows = append(expandedRows, row)
		_, isEmpty := emptyRows[i]
		if isEmpty {
			expandedRows = append(expandedRows, row)
		}
	}

	expandedCols := [][]string {}
	for _, row := range expandedRows {
		newRow := []string {}
		for idx, el := range row {
			newRow = append(newRow, el)
			_, mustDup := emptyCols[idx]
			if mustDup {
				newRow = append(newRow, ".")
			}
		}
		expandedCols = append(expandedCols, newRow)
	}
	return expandedCols
}


func getEmptyLines(grid [][]string) (map[int]struct{}, map[int]struct{}) {
	
	notEmptyCols := map[int]struct{} {}

	emptyRows := map[int]struct{} {}
	emptyCols := map[int]struct{} {}

	for i, row := range grid {
		isEmpty := true
		for j, char := range row {
			if char == "#" {
				notEmptyCols[j] = struct{}{}
				isEmpty = false
				
			}
		}
		if isEmpty {
			emptyRows[i] = struct{}{}
		}
	}

	for col := 0; col < len(grid[0]); col++ {
		_, isNotEmpty := notEmptyCols[col]
		if !isNotEmpty {
			emptyCols[col] = struct{}{}
		}
	}

	return emptyRows, emptyCols
}


func getStartingPoints(grid [][]string) []pointUtils.Point{
	startingPoints := []pointUtils.Point {}
	for i, row := range grid {
		for j, char := range row {
			if char == "#" {
				startingPoint := pointUtils.Point{X: i, Y: j}
				startingPoints = append(startingPoints, startingPoint)
			}
		}
	}
	return startingPoints
}


// instead of doing 93096...
// do 432
// run a full bfs for each starting point and mark all the distances...
type VisitedSet = map[pointUtils.Point]map[pointUtils.Point]struct{}
func solvePart1(grid [][]string) int {
	
	totalLengths := 0
	visitedSet := VisitedSet{}
	galaxies := getStartingPoints(grid)

	for _, startGalaxy := range galaxies {

		currentCombinations := BFS(startGalaxy, grid)

		for endGalaxy, length := range currentCombinations {
			
			if !startGalaxy.Compare(endGalaxy) {
				isEndInStart := pointInPointMap(startGalaxy, endGalaxy, visitedSet)
				isStartInEnd := pointInPointMap(endGalaxy, startGalaxy, visitedSet)
				
				if isEndInStart || isStartInEnd{
					//pass
				} else {
					//here count gaps then mu,tltpytl
					totalLengths += length 
				}
			}
				
		}
	}
	
	return totalLengths
}
func solvePart2(grid [][]string, modifier int) int {
	
	
	totalLengths := 0
	visitedSet := VisitedSet{}
	galaxies := getStartingPoints(grid)

	emptyRows, emptyCols := getEmptyLines(grid)

	for _, startGalaxy := range galaxies {

		currentCombinations := BFS(startGalaxy, grid)

		for endGalaxy, length := range currentCombinations {
			
			if !startGalaxy.Compare(endGalaxy) {
				isEndInStart := pointInPointMap(startGalaxy, endGalaxy, visitedSet)
				isStartInEnd := pointInPointMap(endGalaxy, startGalaxy, visitedSet)
				
				if isEndInStart || isStartInEnd{
					//pass
				} else {
					//here count gaps then mu,tltpytl
					rows, cols := startGalaxy.Diff(endGalaxy)
					overlaps := countOverlap(rows, emptyRows) + countOverlap(cols, emptyCols)
					totalLengths += length + overlaps *( modifier - 1 )
					/*
					fmt.Println("START : ", startGalaxy)
					fmt.Println("END   : ", endGalaxy)
					fmt.Println("ROWS DIF  : ", rows)
					fmt.Println("COLS DIF  : ", cols)
					fmt.Println("OVERLAPS : ", overlaps)
					fmt.Println("________________________")
					*/
				}
			}
				
		}
	}
	
	return totalLengths
}


func countOverlap(pRange [2]int, pMap map[int]struct{}) int {

	overlaps := 0

	for x := pRange[0] + 1; x < pRange[1]; x++{
		_, exists := pMap[x]
		if exists {
			overlaps++
		}
	}
	return overlaps
}

// check for e.g. PointA - PointB
func pointInPointMap(pointA pointUtils.Point, pointB pointUtils.Point, visitedSet VisitedSet) bool {
	startMap, startExists := visitedSet[pointA]
	if !startExists {
		visitedSet[pointA] = map[pointUtils.Point]struct{} {}
		startMap = visitedSet[pointA]
		
	}
	_, isBInA := startMap[pointB]//startEndMap
	if !isBInA {
		startMap[pointB] = struct{} {}
	}

	return isBInA
}


func BFS(start pointUtils.Point, grid [][]string) map[pointUtils.Point]int {
	
	pathLentgh := 0

	sizeX := len(grid) 
	sizeY := len(grid[0]) 

	visitedSet := map[pointUtils.Point]struct{} {}

	combinations := map[pointUtils.Point]int {}

	q := pointUtils.GetDoublyList()
	q.PushBack(start)
	

	for q.HasMore() {

		nextQ := pointUtils.GetDoublyList()
		
		for q.HasMore() {
			currentPoint := q.PopLeft()
			
			_, isVisited := visitedSet[currentPoint]
			if isVisited {
				continue
			}

			if grid[currentPoint.X][currentPoint.Y] == "#" {
				combinations[currentPoint] = pathLentgh
			}

			visitedSet[currentPoint] = struct{}{}
			
			for _, nextPoint := range currentPoint.GetCrossPoints(sizeX, sizeY) {
				nextQ.PushBack(nextPoint)
			}
			
		}

		q = nextQ
		pathLentgh += 1
	}
	
	return combinations
}


// found how many gaps are in rows 
// in vcols
// find curent shortst path for each row and col in gap between ther 2 points
// add 1 m,ore step (1 step is 1000000)

// POINT a POIUNT b 
// has X rows gaps 
/// has Y cols gaps

// c curent len  + X*MULT + Y*MULT
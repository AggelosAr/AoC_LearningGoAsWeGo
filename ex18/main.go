package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"time"

	"aggelos.com/go/aoc/ex18/utils"
)


func main() {

	start := time.Now()
	part1result := solvePart1(getData(1))
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)

	
	start = time.Now()
	part2result := solvePart2(getData(2))
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 50746 && part2result == 70086216556038 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
	
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData(decoder int) []utils.Cell {

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cells := []utils.Cell {}

	for scanner.Scan() {
		line := scanner.Text()

		cell := utils.Cell{}
		if decoder == 1 {
			cell = utils.GetNewCell(line)
		} else {
			cell = utils.GetNewCell2(line)
		}
		cells = append(cells, cell)
	}

	return cells
}


func solvePart1(cells []utils.Cell) int {
	//visualizeData(cells)
	return shoelace(cells)
}
func solvePart2(cells []utils.Cell) int {
	//visualizeData(cells)2 big
	return shoelace(cells)
}


func shoelace(cells []utils.Cell) int {
	
	
	prevPoint := utils.Point{X: 0, Y: 0}

	perimeter := 0
	areaInShape := 0
	
	for _, cell := range cells {

		currentOffset := utils.Offsets[cell.Direction]
		currentPoint := prevPoint

		for step := 0; step < cell.Distance; step++ {
			currentPoint.AddInPlace(currentOffset)
			perimeter++
		}

		areaInShape += prevPoint.CrossPoints(currentPoint)
		prevPoint = currentPoint
	}

	newArea := math.Abs(float64(areaInShape)) / 2
	// pick's theorem - find the number of points in a shape given its area
	pickArea := newArea - (0.5 * float64(perimeter) ) + 1

	return perimeter + int(pickArea)
}





///////////////////////////////////////////////////////////////

func makeGrid(cells []utils.Cell) [][]string {
	
	minX, maxX, minY, maxY := findGridLimits(cells)
	
	rows := maxX - minX + 1
	cols := maxY - minY + 1

	adjustedOffsets := utils.Point{
		X: int(math.Abs(float64(minX))),
		Y: int(math.Abs(float64(minY))),
	}

	
	//fmt.Println("ROWS, COLS : ", rows, cols )
	//return [][]string {}
	grid := createNewGrid(rows, cols)
	
	movingPoint := utils.Point{X: 0, Y: 0}
	movingPoint.AddInPlace(adjustedOffsets) 

	for _, cell := range cells {
		currentOffset := utils.Offsets[cell.Direction]
		for step := 0; step < cell.Distance; step++ {
			movingPoint.AddInPlace(currentOffset) 
			
			grid[movingPoint.X][movingPoint.Y] = "#"
		}
	}
	return grid
}         



func findGridLimits(cells []utils.Cell) (int, int, int, int){

	minX := 1000000000
	maxX := -1000000000

	minY := 10000000000
	maxY := -1000000000
	

	minStep := 10000000000
	maxStep := -1000000000
	//sumulate the process of traversing like it wa a grid
	movingPoint := utils.Point{X: 0, Y: 0}
	for _, cell := range cells {
		currentOffset := utils.Offsets[cell.Direction]
		
		for step := 0; step < cell.Distance; step++ {
			movingPoint.AddInPlace(currentOffset)
			maxX = max(maxX, movingPoint.X)
			minX = min(minX, movingPoint.X)
			maxY = max(maxY, movingPoint.Y)
			minY = min(minY, movingPoint.Y)
		}
		minStep = min(minStep, cell.Distance)
		maxStep = max(maxStep, cell.Distance)
	}

	//fmt.Println("MIN/MAX distance : ", minStep, maxStep)
	return minX, maxX, minY, maxY
}


func createNewGrid(rows, cols int) [][] string {
	grid := [][]string {} 
	for i := 0; i < rows; i++ {
		newRow := []string {}
		for j := 0; j < cols; j++ {
			newRow = append(newRow, ".")
		}
		grid = append(grid, newRow)
	}
	return grid
}

func printGrid(grid [][] string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("_______________________")
}

func visualizeData(cells []utils.Cell) {
	printGrid(makeGrid(cells))
}

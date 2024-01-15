package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	if part1result == 7939 && part2result == 8318 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}
	
	start = time.Now()
	part2resultGoRoutines := solvePart2WithGoRoutines(grid)
	elapsed = time.Since(start)

	fmt.Printf("Time took with goroutines %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2resultGoRoutines)
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
	

type Beam struct {
	Point Point
	Dir string
}


type Point struct {
	X, Y int
}


func (p Point) Add(pointB Point) Point {
	return Point{X: p.X + pointB.X, Y: p.Y + pointB.Y}
}


func (p Point) IsValid(sizeX, sizeY int) bool {
	return (-1 < p.X && p.X < sizeX) && (-1 < p.Y && p.Y < sizeY)
}


var offsets = map[string]Point{
	"UP":    {X: -1, Y: 0},
	"DOWN":  {X: 1, Y: 0},
	"LEFT":  {X: 0, Y: -1},
	"RIGHT": {X: 0, Y: 1},
}


func energizeGrid(grid [][]string, startingBeam Beam) int {
	
	sizeX := len(grid)
	sizeY := len(grid[0])

	visitedBeams := map[Beam]struct{} {}
	// we will use another one visited set only for the points
	visitedPoints := map[Point]struct{} {}

	q := list.New()
	q.PushBack(startingBeam)

	
	for q.Len() > 0 {
		nextQ := list.New()

		for q.Len() > 0 {

			poppedEl := q.Front()
			q.Remove(poppedEl)

			currBeam := poppedEl.Value.(Beam)

			_, seenBeam := visitedBeams[currBeam]
			if seenBeam {
				continue
			}
			visitedBeams[currBeam] = struct{} {}
			visitedPoints[currBeam.Point] = struct{} {}

			x := currBeam.Point.X
			y := currBeam.Point.Y
			nextDirs := handleBeam(grid[x][y], currBeam.Dir)

			for _, nextDir := range nextDirs {

				offset := offsets[nextDir]
				nextPoint := currBeam.Point.Add(offset)
				
				if nextPoint.IsValid(sizeX, sizeY) {
					nextBeam := Beam{Point: nextPoint, Dir: nextDir}
					q.PushBack(nextBeam)
				}
			}
		}
		q = nextQ
	}
	return len(visitedPoints)
}


func energizeGridGo(grid [][]string, startingBeam Beam, c chan int) {
	
	sizeX := len(grid)
	sizeY := len(grid[0])

	visitedBeams := map[Beam]struct{} {}
	// we will use another one visited set only for the points
	visitedPoints := map[Point]struct{} {}
	
	q := list.New()
	q.PushBack(startingBeam)
	
	for q.Len() > 0 {
		nextQ := list.New()

		for q.Len() > 0 {

			poppedEl := q.Front()
			q.Remove(poppedEl)

			currBeam := poppedEl.Value.(Beam)

			_, seenBeam := visitedBeams[currBeam]
			if seenBeam {
				continue
			}
			visitedBeams[currBeam] = struct{} {}
			visitedPoints[currBeam.Point] = struct{} {}

			x := currBeam.Point.X
			y := currBeam.Point.Y
			nextDirs := handleBeam(grid[x][y], currBeam.Dir)

			for _, nextDir := range nextDirs {

				offset := offsets[nextDir]
				nextPoint := currBeam.Point.Add(offset)
				
				if nextPoint.IsValid(sizeX, sizeY) {
					nextBeam := Beam{Point: nextPoint, Dir: nextDir}
					q.PushBack(nextBeam)
				}
			}
		}
		q = nextQ
	}
	c <- len(visitedPoints)
}


func handleBeam(contraption, direction string) []string{

	nextDirections := []string {}
	if contraption == "." {
		nextDirections = append(nextDirections, direction)
		return nextDirections
	}

	if direction == "UP" {
		if contraption == "/" {
			nextDirections = append(nextDirections, "RIGHT")
		} else if contraption == "-" {
			nextDirections = append(nextDirections, "LEFT", "RIGHT")
		} else if contraption == "\\" {
			nextDirections = append(nextDirections, "LEFT")
		}
	} else if direction == "DOWN" {
		if contraption == "/" {
			nextDirections = append(nextDirections, "LEFT")
		} else if contraption == "-" {
			nextDirections = append(nextDirections, "LEFT", "RIGHT")
		} else if contraption == "\\" {
			nextDirections = append(nextDirections, "RIGHT")
		}
	} else if direction == "LEFT" {
		if contraption == "|" {
			nextDirections = append(nextDirections, "UP", "DOWN")
		} else if contraption == "/" {
			nextDirections = append(nextDirections, "DOWN")
		} else if contraption == "\\" {
			nextDirections = append(nextDirections, "UP")
		}
	} else if direction == "RIGHT" {
		if contraption == "|" {
			nextDirections = append(nextDirections, "UP", "DOWN")
		} else if contraption == "/" {
			nextDirections = append(nextDirections, "UP")
		} else if contraption == "\\" {
			nextDirections = append(nextDirections, "DOWN")
		}
	}

	if len(nextDirections) == 0 {
		nextDirections = append(nextDirections, direction)
	}
	return nextDirections
}


func solvePart1(grid [][]string) int {
	startingBeam := Beam{Point: Point{X: 0, Y: 0}, Dir: "RIGHT"}
	return energizeGrid(grid, startingBeam)
}


// This can become better
// We can use the Beam Map to store and find the path of the current beam
// instead of recalculating many times.
// e.g. we trace each bream from start to end
// if it splits we now have 2 seperate beams with same start but different paths 
// when the beam loops or ends we add to the map 
func solvePart2(grid [][]string) int {

	result := 0
	rows := len(grid)
	cols := len(grid[0])

	for col := 0; col < cols; col++ {
		// 1st row -> Dir Down
		startingBeam := Beam{Point: Point{X: 0, Y: col}, Dir: "DOWN"}
		result = max(result, energizeGrid(grid, startingBeam))
		// Last row -> Dir Up
		startingBeam = Beam{Point: Point{X: rows - 1, Y: col}, Dir: "UP"}
		result = max(result, energizeGrid(grid, startingBeam))
	}
	
	for row := 0; row < rows; row++ {
		// 1st Col -> Dir Right
		startingBeam := Beam{Point: Point{X: row, Y: 0}, Dir: "RIGHT"}
		result = max(result, energizeGrid(grid, startingBeam))
		// Last col -> Dir Left
		startingBeam = Beam{Point: Point{X: row, Y: cols - 1}, Dir: "LEFT"}
		result = max(result, energizeGrid(grid, startingBeam))
	}
	
	return result
}


func solvePart2WithGoRoutines(grid [][]string) int {

	result := 0
	rows := len(grid)
	cols := len(grid[0])
	
	c := make(chan int, cols*2 + rows*2)
	routinesCount := 0

	for col := 0; col < cols; col++ {
		// 1st row -> Dir Down
		startingBeam := Beam{Point: Point{X: 0, Y: col}, Dir: "DOWN"}
		go energizeGridGo(grid, startingBeam, c)
		// Last row -> Dir Up
		startingBeam = Beam{Point: Point{X: rows - 1, Y: col}, Dir: "UP"}
		go energizeGridGo(grid, startingBeam, c)
		routinesCount += 2
	}
	
	for row := 0; row < rows; row++ {
		// 1st Col -> Dir Right
		startingBeam := Beam{Point: Point{X: row, Y: 0}, Dir: "RIGHT"}
		go energizeGridGo(grid, startingBeam, c)
		// Last col -> Dir Left
		startingBeam = Beam{Point: Point{X: row, Y: cols - 1}, Dir: "LEFT"}
		go energizeGridGo(grid, startingBeam, c)
		routinesCount += 2
	}
	
	// Receive results from goroutines.
	for i := 0; i < routinesCount; i++ {
		result = max(result, <-c)
	}

	// Close the channel.
	close(c)
	fmt.Println("Total : ", routinesCount)

	return result
}



func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("__________________")
}

func printGridTemp(grid [][]string, visitedSetPoints map[Point]struct{}) {
	gridCopy := copyArr(grid)
	for k := range visitedSetPoints {
		gridCopy[k.X][k.Y] = "#"
	}
	printGrid(gridCopy)
}

func copyArr(arr [][]string) [][]string {
	newArr := [][]string {}
	for i := 0; i < len(arr); i++ {
		row := []string {}
		for j := 0; j < len(arr[0]); j++ {
			row = append(row, arr[i][j])
		}
		newArr = append(newArr, row)
	}
	return newArr
}


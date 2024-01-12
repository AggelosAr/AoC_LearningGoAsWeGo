package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex10/pointUtils"
)


func main() {
	rawData := readData()
	data := formatData(rawData)

	dataOriginal := formatData(rawData)

	start := time.Now()
	steps, _ := runBFS(data)
	elapsed := time.Since(start)

	
	fmt.Println("The result (PART1) : ", steps)
	fmt.Printf("Time took %s\n", elapsed)


	start = time.Now()
	bordersWithDirs := getBordersDirs(dataOriginal)
	resultPart2 := countIslands(data, bordersWithDirs)
	elapsed = time.Since(start)
	fmt.Println("The result (PART2) : ", resultPart2)
	fmt.Printf("Time took %s\n", elapsed)
	if steps == 7086 && resultPart2 == 0 {
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

func formatData(text string) [][] string{

	lines := strings.Split(text, "\n")
	data := [][] string {}

	for idx := 0; idx < len(lines); idx++ {
		if lines[idx] != "" {
			data = append(data, strings.Split(lines[idx], ""))
		}
	}
	return data
}


// returns the steps, all the borders found
func runBFS(data [][]string) (int, []pointUtils.Point) {
	steps := 0

	startingPoint := getStartingPoint(data)
	q := pointUtils.NewPointQ()
	q.Add(startingPoint)

	nextQ := pointUtils.NewPointQ()

	borders := []pointUtils.Point {startingPoint}

	for {
		for {
			currentPoint, err := q.PopLeft()
			if err != nil {
				break
			}
			
			for _, nextPoint := range getNextPoints(currentPoint, data) {
				if checkCompatibility(currentPoint, nextPoint, data){
					nextQ.Add(nextPoint)
					borders = append(borders, nextPoint)
				}
			}
			// Mark the current point as visited so we don't have to look at it again
			data[currentPoint.X][currentPoint.Y] = "X"
		}

		q.List = nextQ.List
		nextQ = pointUtils.NewPointQ()
		if q.List.Len() == 0 {
			break
		}
		steps++
	}

	// Last item is added twice
	// check is there are at least 3 items .....don't
	size := len(borders)
	if borders[size - 1].Compare(borders[size - 2]) {
		borders = borders[: size - 1]
	}
	// len borders should be  ==  2 * steps
	return steps, borders
}


func getStartingPoint(data [][]string) pointUtils.Point {
	startingPoint := pointUtils.Point{}
	for i, row := range data {
		for j, char := range row {
			if char == "S" {
				startingPoint.X = i
				startingPoint.Y = j
				startingPoint.Symbol = "S"
			}
		}
	}
	return startingPoint
}


func getNextPoint(point pointUtils.Point, offset pointUtils.Offset, data [][]string) (pointUtils.Point, error) {
	nextPoint := point.Add(offset)
	return nextPoint, nextPoint.Validate(len(data), len(data[0]))
}


func getNextPoints(point pointUtils.Point, data [][]string) []pointUtils.Point {

	nextPoints := []pointUtils.Point {}
	direction := data[point.X][point.Y]
	
	for _, offset := range pointUtils.DirectionsMap[direction] {

		nextPoint, err := getNextPoint(point, pointUtils.DirectionsToOffset[offset], data)
		if err != nil {
			continue
		}
		nextPoint.Symbol = data[nextPoint.X][nextPoint.Y]
		// Moved this here so i can reuse the getNextPoint func for Part 2
		_, exists := pointUtils.DirectionsMap[data[nextPoint.X][nextPoint.Y]]
		if !exists {
			continue
		}
		nextPoints = append(nextPoints, nextPoint)
	
	}
	return nextPoints
}


// We can simply check if we can go from pointA to B and vice versa
// Since we already know we can go from point A to B 
// We just need to check if we can go from point B to A
func checkCompatibility(currentPoint pointUtils.Point, nextPoint pointUtils.Point, data [][]string) bool{
	for _, nextPoint := range getNextPoints(nextPoint, data) {
		if nextPoint.Compare(currentPoint) {
			return true
		}
	}
	return false
}




func countIslands(data [][] string, bordersWithDirs pointUtils.BorderPoints) int {

	startingPoint := pointUtils.Point {}
	islands := 0

	for i, row := range data {
		for j, char := range row {
			if char != "X" && char != "+" && char != "-"{
				startingPoint.X = i
				startingPoint.Y = j

				isIsland := runSimpleBFS(data, startingPoint, bordersWithDirs)
				
				if isIsland {
					islands++
				}
			}
		}
	}
	
	enclosedTiles := printData(data, islands)
	return enclosedTiles
}



// TODO BETTER
func getBordersDirs(data [][]string) pointUtils.BorderPoints {
	bordersPoints := pointUtils.BorderPoints {Points: map[string]pointUtils.BorderPoint{}}

	currentPoint := getStartingPoint(data)
	
	visited := map[string]struct{}{}
	
	line := pointUtils.GetNewLine()
	line.Add(currentPoint)

	foundNextPoint := true

	for foundNextPoint {

		visited[currentPoint.ToKey()] = struct{}{}
		nextPoint := pointUtils.Point{}
		foundNextPoint = false		
		/* DEBUG END 
		if currentPoint.X == 1 && currentPoint.Y == 3 {
			fmt.Println(currentPoint)
		}
		if currentPoint.X == 2 && currentPoint.Y == 2 {
			fmt.Println(currentPoint)
		}
		*/
		for _, nextPoint = range getNextPoints(currentPoint, data) {
			
			_, isVisited := visited[nextPoint.ToKey()]
			if isVisited {
				continue
			}
			visited[nextPoint.ToKey()] = struct{}{}
			
			if checkCompatibility(currentPoint, nextPoint, data){
				foundNextPoint = true
				break
			}
		}

		line.Add(nextPoint)
		// we found S again 
		// we ended the loop // TODO
		if !foundNextPoint {
			break
		}
		
		if line.IsComplete() {
			bordersPoints.Add(line.GetBorderPoints())
			line.Clear()
			line.Add(nextPoint)
			currentPoint = nextPoint
		}		
	}
	/* DEBUG END 
	fmt.Println(bordersPoints.Points["(1-3)"])
	fmt.Println(bordersPoints.Points["(2-2)"])
	fmt.Println(bordersPoints.Points["(2-4)"])
	fmt.Println(bordersPoints.Points["(3-3)"])
	*/
	return bordersPoints
}
 

// count islands
func runSimpleBFS(data [][]string, startingPoint pointUtils.Point, borders pointUtils.BorderPoints) bool{
	
	q := pointUtils.NewPointQ()
	q.Add(startingPoint)
	nextQ := pointUtils.NewPointQ()

	canInclude := true
	//  To include the current island there are 2 checks that must pass/
	// 1. We must not encounter an out of border error
	// 2. The island must be enclosed by border:
	// 		a. Only by the inner part of the borders

	currentIsland := []pointUtils.Point {startingPoint}
	
	rotations := [4]pointUtils.Offset {{X:-1, Y:0}, {X:1, Y:0}, {X:0, Y:-1}, {X:0, Y:1}}

	for {
		for {
			currentPoint, err := q.PopLeft()
			if err != nil {
				break
			}
			
			for _, rot := range rotations {

				nextPoint, err := getNextPoint(currentPoint, rot, data)
				// CASE 1
				if err != nil {
					canInclude = false
					continue
				}
				// Move conditions
				// If the next Point is a border
				if data[nextPoint.X][nextPoint.Y] == "X" {
					// If we already determined that we can't include this island -> SKIP
					if !canInclude{
						continue
					}
					// CASE 2
					// e.g. The current Point must be in the In points
					//		Else is is on the outside   


					// if we find at least 1 point in the border . In we can 
					// include the island 
					if canInclude {
						continue
					}

					isIn := false
					
					//TODO CHECK THIS MUST BE CORRECT !! 
					for _, point := range borders.Points[nextPoint.ToKey()].In {
						if point.Compare(currentPoint) {
							isIn = true
						}
					}
					
					if isIn {
						canInclude = true
					}
					continue
				}
				// Can't move on previously seen tile
				if data[nextPoint.X][nextPoint.Y] == "@"{
					continue
				}
				//TEMP +  means the tile is included
				if data[nextPoint.X][nextPoint.Y] == "+"{
					continue
				}
				//TEMP -  means the tile is excluded
				if data[nextPoint.X][nextPoint.Y] == "-"{
					continue
				}
				
				nextQ.Add(nextPoint)
				// Add this point to the island
				currentIsland = append(currentIsland, nextPoint)
			
			
			}
			// Mark the current point as visited so we don't have to look at it again
			data[currentPoint.X][currentPoint.Y] = "@"
		}

		q.List = nextQ.List
		nextQ = pointUtils.NewPointQ()
		if q.List.Len() == 0 {
			break
		}
	}

	markData(data, currentIsland, canInclude) 
	return len(currentIsland) > 0
}


func markData(data [][]string, island []pointUtils.Point, canInclude bool) {
	mark := "-"
	if canInclude {
		mark = "+"
	}
	for _, point := range island{
			data[point.X][point.Y] = mark
		}
}


func printData(data [][]string, islands int) int {
	for _, row := range data {
		fmt.Println(row)
	}

	enclosedTiles := 0
	totalTiles := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == "+" {
				enclosedTiles++
				totalTiles++
			} else if data[i][j] == "-" {
				totalTiles++
			}
		}
	}

	fmt.Println("Islands Found : ", islands)
	fmt.Println("Tiles found enclosed : ", enclosedTiles)
	fmt.Println("Total found : ", totalTiles)
	fmt.Println("\n\n_____________________________________")
	return enclosedTiles
}

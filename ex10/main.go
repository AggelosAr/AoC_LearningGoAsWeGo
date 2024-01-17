package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"aggelos.com/go/aoc/ex10/pointUtils"
)


func main() {
	rawData := readData()
	

	start := time.Now()
	steps, _ := runBFS(formatData(rawData))
	elapsed := time.Since(start)	
	fmt.Println("The result (PART1) : ", steps)
	fmt.Printf("Time took %s\n", elapsed)

	start = time.Now()
	resultPart2 := shoelace(formatData(rawData))
	elapsed = time.Since(start)
	fmt.Println("The result (PART2) : ", resultPart2)
	fmt.Printf("Time took %s\n", elapsed)
	if steps == 7086 && resultPart2 == 317 {
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
	if size > 3 {
		if borders[size - 1].Compare(borders[size - 2]) {
			borders = borders[: size - 1]
		}
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


func printData(data [][]string) {
	for _, row := range data {
		fmt.Println(row)
	}
}


// returns the steps, all the borders found
func getBordersInCorrectOrder(data [][]string) (int, []pointUtils.Point) {
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
					break
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
	if size > 3 {
		if borders[size - 1].Compare(borders[size - 2]) {
			borders = borders[: size - 1]
		}
	}
	
	// len borders should be  ==  2 * steps
	return steps, borders
}


func shoelace(data [][]string) int {
	
	
	_, borders := getBordersInCorrectOrder(data)
	
	prevPoint := borders[0]

	perimeter := len(borders)
	areaInShape := 0
	
	for idx := 1; idx < len(borders); idx++ {

		
		currentPoint := borders[idx]
		//fmt.Println(prevPoint, currentPoint)
		areaInShape += prevPoint.CrossPoints(currentPoint)
		prevPoint = currentPoint
	}

	currentPoint := borders[0]
	areaInShape += prevPoint.CrossPoints(currentPoint)

	//fmt.Println(prevPoint, currentPoint)
	newArea := math.Abs(float64(areaInShape)) / 2
	// pick's theorem - find the number of points in a shape given its area
	pickArea := newArea - (0.5 * float64(perimeter) ) + 1

	//fmt.Println("perimeter : ", perimeter)
	//fmt.Println("areaInShape : ", areaInShape)
	//fmt.Println("PICK AREA : ", pickArea)
	//fmt.Println("newArea : ", newArea)
	return int(pickArea)//+ perimeter
}

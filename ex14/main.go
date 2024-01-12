package main

import (
	"bufio"
	"fmt"
	"math"
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
	rotations   := 180
	part2result := solvePart2(grid, rotations)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s for <%d> rotations\n", elapsed, rotations)
	fmt.Println("PART2 - RESULT - ", part2result)

	if part1result == 102497 && part2result == 105008 {
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
	


func solvePart1(grid [][]string) int {
	tiltGridUp(grid)
	return findTotalLoad(grid)
}



// here we solved using the result of the grid
// it is not certain tho ?! or maybe it is 
// instead we must store all the grids in an array and find repetition there...
func solvePart2(grid [][]string, times int) int {
	copyGrid := copyArr(grid)
	// Given a decent number of rotations in the initial we should have a pattern 
	// e.g. 100-200 times should be enough ^^ totations 
	// If we can't find a result or panic increase them
	loads := rotateManyTimes(copyGrid, times)
	return parseResults(loads, grid)
}


func tiltGridUp(grid [][]string) {
	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {

			rowStart := row
			rocksFound := 0
			
			for row < len(grid) {
				if grid[row][col] == "O" {
					rocksFound++
				} else if grid[row][col] == "#" {
					break
				}
				row++
			}
			
			for x := rowStart; x < row; x++ {
				if rocksFound > 0 {
					grid[x][col] = "O"
					rocksFound--
				} else {
					grid[x][col] = "."
				}
			}
		}
	}
}


func tiltGridLeft(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == "O" {
				// Find the position to move the rock to the left
				destCol := col
				for destCol > 0 && grid[row][destCol-1] == "." {
					destCol--
				}
				// Move the rock to the left
				grid[row][col], grid[row][destCol] = grid[row][destCol], grid[row][col]
			}
		}
	}
}


func tiltGridRight(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for col := len(grid[0]) - 1; col >= 0; col-- {
			if grid[row][col] == "O" {
				// Find the position to move the rock to the right
				destCol := col
				for destCol < len(grid[0])-1 && grid[row][destCol+1] == "." {
					destCol++
				}
				// Move the rock to the right
				grid[row][col], grid[row][destCol] = grid[row][destCol], grid[row][col]
			}
		}
	}
}

func tiltGridDown(grid [][]string) {
	for col := 0; col < len(grid[0]); col++ {
		for row := len(grid) - 1; row >= 0; row-- {
			if grid[row][col] == "O" {
				// Find the position to move the rock downward
				destRow := row
				for destRow < len(grid)-1 && grid[destRow+1][col] == "." {
					destRow++
				}
				// Move the rock downward
				grid[row][col], grid[destRow][col] = grid[destRow][col], grid[row][col]
			}
		}
	}
}



func rotateManyTimes(grid [][]string, times int) []int{
	totalLoads := []int {}
	for i := 0; i < times; i++ {
		tiltGridUp(grid)
		tiltGridLeft(grid)
		tiltGridDown(grid)
		tiltGridRight(grid)
		totalLoads = append(totalLoads, findTotalLoad(grid))
	}
	return totalLoads
}


func findTotalLoad(grid [][]string) int {
	totalLoad := 0
	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			if grid[row][col] == "O" {
				totalLoad += len(grid) - row 
			}	
		}
	}
	return totalLoad
}


// e.g. res X X X X X X X -> 1 2 3 4 5 <- 1 2 3 4 5 1 2 3 

// need to find the repeating pattern in the result array 
// Then calculate which element has the 1 000 000 000 positition.
func parseResults(results []int, grid[][] string) int {

	lastPattern := findPattern(results)
	lastPatternMap := convertArrToMap(lastPattern)

	// Here we have the patten BUT it may be out of order

	//If we want to find the pattern in the correct order
	// iterate the results 1 by 1 until you find the pattern again
	// we will use a hashmap of int -> int
	// when we find the pattern again we gucci

	
	newLoads := []int {}
	correctPattern := []int {}

	i := 0
	for i < len(results) {
		
		newLoads = append(newLoads, rotateManyTimes(grid, 1)[0])
		currentPattern := findPattern(newLoads)
		
		if compareHashMaps(lastPatternMap, convertArrToMap(currentPattern)) {
			correctPattern = currentPattern
			break
		}
		i++
	}
	
	uniqueItems := i - (2 * len(correctPattern)) + 1
	remaining := float64(1000000000 - uniqueItems)
	remainingElements := math.Mod(remaining, float64(len(correctPattern)))
	
	extrapolatedIdx := 0
	if remainingElements == 0 {
		extrapolatedIdx = len(correctPattern) - 1
	} else {
		extrapolatedIdx = int(remainingElements) - 1
	}
	return correctPattern[extrapolatedIdx]
}

// !MAYBE! the findPattern algorithm needs fixing !works for now!
// start from end 
// since the grid may not have reached an equilibrium position from the start
// In essence we need to chop off a starting slice and an ending slice 
// starts on the last el 
// Tries to find the same el from the previous index and backwards
// then compares from those 2 positions the previous elements
// when it cant go anymore we have found the pattern

// since we append in reverse we need to reverse the pattern we found 
// Also we need to put it the correct order


// !! 
// With this approach if we encounter the same element 2 times
// it will find that element as the pattern and will ignore the bigger picture
// In that case we must keep going untill we cant find a bigger pattern

func findPattern(results []int) []int {
	pattern := []int {}
	lastIdx := len(results) - 1

	if lastIdx < 1 {
		return pattern
	}

	searchingElement := results[lastIdx]
	prevIdx := lastIdx - 1
	
	times := 0 
	for prevIdx > -1 {
		
		prevVal := results[prevIdx]
		lastIdxCopy := lastIdx
		currentPattern := []int {}

		// We want to find the previous element same as last 
		for prevIdx > - 1 {
			
			prevVal = results[prevIdx]
			if prevVal == searchingElement {
				break
			}
			prevIdx-- 
		}
		if prevIdx == -1 {
			break
		}
		
		prevIdxCopy := prevIdx
		
		for results[lastIdxCopy] == results[prevIdxCopy]{

			currentPattern = append(currentPattern, results[lastIdxCopy])
			lastIdxCopy--
			prevIdxCopy--

			if lastIdxCopy == prevIdx {
				break
			}
			if prevIdxCopy == -1 {
				break
			}
		}

		// lastIdxCopy == prevIdx 
		// To be a pattern the 2 parts MUST be connected
		if lastIdxCopy == prevIdx && len(currentPattern) > len(pattern) {
			if compareArr(currentPattern, pattern) {
				break
			}
			pattern = currentPattern
		}

		prevIdx--
		times++
		if times == 2 {
			break
			// if we dont break here it go on finding the pattern in repeat 
			// not correct but works for now
		}
	}
	// check we may need to half it e.g. we found it repeated
	if compareArr(pattern[:len(pattern)/2], pattern[len(pattern)/2:]) {
		pattern = pattern[:len(pattern)/2]
	}
	return reverseArr(pattern)
	
}


////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////


func compareArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reverseArr(arr []int) []int {
	reversedArr := []int {}
	for i := len(arr) - 1; i > - 1; i-- {
		reversedArr = append(reversedArr, arr[i])
	}
	return reversedArr
}

func swapArrAtIdx(arr []int, idx int) []int {
	return append(arr[idx:], arr[:idx]...)
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


func convertArrToMap(arr []int) map[int]int {
	currentMap := map[int]int {}
	for _, el := range arr {
		_, exists := currentMap[el]
		if exists {
			currentMap[el]++
		} else {
			currentMap[el] = 1
		}
	}
	return currentMap
}

func compareHashMaps(map1, map2 map[int]int) bool {
	areSame := true
	for k, v := range map1 {
		v2, exists := map2[k]
		if exists && v == v2 {
			//pass
		} else {
			areSame = false
			break
		}
	}
	return areSame
}
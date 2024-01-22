package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aggelos.com/go/aoc/ex3/utils"
)


func main() {
	matrix := getMatrixFromData()

	part1Result := solveMatrix(matrix)
	part2Result := solveMatrixPart2(matrix)
	fmt.Println("The result is PART1: ", part1Result)
	fmt.Println("The result is PART2: ", part2Result)

	if 546563 == part1Result && 91031374 == part2Result {
		fmt.Println("G O O D ")
	}
}



func solveMatrix(matrix [][]string) int {

	var (
		res = 0
		rows = len(matrix)
		cols = len(matrix[0])
	)
	
	// We search to find numbers then look around to see if we will
	// include that number
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			currentChar := matrix[i][j]
			number, err := strconv.Atoi(currentChar)

			if err != nil {
				continue
			}

			point := *utils.NewPoint(i, j)

			if lookAroundForSymbol(matrix, point){
				
				before, after, jOffset := gatherNumber(matrix[i], j, cols)
				currentNumber := before + strconv.Itoa(number) + after
				finalNumber, _ := strconv.Atoi(currentNumber)

				j += jOffset
				res += finalNumber
				
			}
		}
	}

	return res
}


func solveMatrixPart2(matrix [][]string) int {

	var (
		result = 0
		rows = len(matrix)
		cols = len(matrix[0])
	)
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == "*" {
				point := *utils.NewPoint(i, j)
				result += lookAroundForNumbers(matrix, point)
			}
		}
	}
	return result
}


func lookAroundForSymbol(matrix [][]string, point utils.Point) bool{

	var (
		rows = len(matrix)
		cols = len(matrix[0])
	)

	for _, offset := range utils.Offsets {

		x := point.X + offset[0]
		y := point.Y + offset[1]
		
		currPoint := *utils.NewPoint(x, y)
		if currPoint.IsValid(rows, cols) && isSymbol(matrix[x][y]) {
			return true
			
		}
	}
	return false
}


func lookAroundForNumbers(matrix [][]string, point utils.Point) int{

	var (
		rows = len(matrix)
		cols = len(matrix[0])
		result = 1
		numbersFound = 0
		visitedSet = map[utils.Point]struct{}{}
	)
	
	for _, offset := range utils.Offsets {

		x := point.X + offset[0]
		y := point.Y + offset[1]
		
		currPoint := *utils.NewPoint(x, y)
			
		if currPoint.IsValid(rows, cols) {

			number, err := strconv.Atoi(matrix[x][y])
			if err != nil {
				continue
			}
			
			_, isVisited := visitedSet[currPoint]
			if isVisited {
				continue
			}

			// if we found Number Gather it 

			numbersFound += 1
			if numbersFound > 2{
				return 0
			}
			before, after, jOffset := gatherNumber(matrix[x], y, cols)

			// Handle jOffset so we don't double count numbers
			// This will happen only on UP and DOWN
			// e.g.
			// 11111
			// ..*..
			// 11111

			for k := 0; k < jOffset + 1; k++{
				point := *utils.NewPoint(x, k + y)
				visitedSet[point] = struct{}{}
			}
			
			currentNumber := before + strconv.Itoa(number) + after
			finalNumber, _ := strconv.Atoi(currentNumber)
			//fmt.Println(finalNumber, newY, jOffset, visitedSet)
			result *= finalNumber

		}
	}
	if numbersFound == 2 {
		return result
	}
	return 0
}



func gatherNumber(matrix []string, idx int, size int) (string, string, int) {
	
	var (
		before = ""
		after = ""
		jOffset = 0
	)
	// go back until not number
	//Dont include the current idx HENCE - 1
	for i := idx - 1; i > -1; i-- {
		_, err := strconv.Atoi(matrix[i])
		if err != nil {
			break
		}
		before += matrix[i]
	}
	// go forward untill not number 
	// add it 
	// skip to forward
	for i := idx + 1; i < size; i++ {
		_, err := strconv.Atoi(matrix[i])
		if err != nil {
			break
		}
		after += matrix[i]
		jOffset++
	}

	return reverseString(before), after, jOffset
}



func isSymbol(char string) bool{
	_, err := strconv.Atoi(char)
	return err != nil && char != "." 
}

func reverseString(t string) string {
	runes := []rune(t)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getRootPath() string{
	rootPath, _ := os.Getwd()
	rootPath += "\\data\\input.txt"
	return rootPath
}

func getMatrixFromData() [][]string{
	path := getRootPath()
	file, _ := os.Open(path)
	defer file.Close()
	return formatFileTo2DMatrix(file)
}

func formatFileTo2DMatrix(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var matrix [][]string
	
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))		
	}
	return matrix
}


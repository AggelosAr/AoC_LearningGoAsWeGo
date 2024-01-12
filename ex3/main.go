package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	matrix := getMatrixFromData(getRootPath())
	
	fmt.Println("The result is PART1: ", solveMatrix(matrix))

	fmt.Println("The result is PART2: ", solveMatrixPart2(matrix))
}


func getRootPath() string{
	rootPath, _ := os.Getwd()
	rootPath += "\\data\\input.txt"
	return rootPath
}


func getMatrixFromData(path string) [][]string{
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Errot while opening the file:", err)
	}
	defer file.Close()

	matrix := formatFileTo2DMatrix(file)
	return matrix
}


func formatFileTo2DMatrix(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var matrix [][]string
	
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, formatLine(line))		
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return matrix
}


func formatLine(line string) []string{
	return strings.Split(line, "")
}


func solveMatrix(matrix [][]string) int {
	result := 0

	rows := len(matrix)
	cols := len(matrix[0])

	// We search to find numbers then look around to see if we will
	// include that number
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			currentChar := matrix[i][j]
			number, err := strconv.Atoi(currentChar)

			if err != nil {
				continue
			}

			if lookAroundForSymbol([2]int{i, j}, rows, cols, matrix){
				
				before, after, jOffset := gatherNumber(matrix[i], j, cols)
				//fmt.Printf("LINE <%d><%d> NUMBER <%d> HAS SYMBOL OFFSET<%d> -- BEFORE<%s> AFTER<%s>\n", i,j, number, jOffset, before, after)

				j += jOffset
				currentNumber := before + strconv.Itoa(number) + after
				finalNumber, _ := strconv.Atoi(currentNumber)
				result += finalNumber
				
			}
		}
	}

	return result
}


func solveMatrixPart2(matrix [][]string) int {
	result := 0

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			currentChar := matrix[i][j]
			if isGear(currentChar){
				result += lookAroundForNumbers([2]int{i, j}, rows, cols, matrix)
			}
		}
	}
	return result
}

var offsets = [8][2]int {{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1}, // {0, 0} not needed WOW OPTIMIZATIONS K
	{1, -1}, {1, 0}, {1, 1}}


func lookAroundForSymbol(idx [2]int, rows int, cols int, matrix [][]string) bool{

	for _, offset := range offsets {

		newX := idx[0] + offset[0]
		newY := idx[1] + offset[1]
		
		if checkLimits(newX, rows) && checkLimits(newY, cols) {
			if isSymbol(matrix[newX][newY]) {
				return true
			}
		}
	}
	return false
}

type Point struct {
	X, Y int
}

func lookAroundForNumbers(idx [2]int, rows int, cols int, matrix [][]string) int{

	result := 1
	numbersFound := 0
	
	visitedSet := map[string]struct{}{}

	for _, offset := range offsets {

		newX := idx[0] + offset[0]
		newY := idx[1] + offset[1]
		
		currentIDX := Point{newX, newY}
		if checkLimits(newX, rows) && checkLimits(newY, cols) {
			number, err := strconv.Atoi(matrix[newX][newY])
			if err != nil {
				continue
			}
			
			if isInVisitedSet(visitedSet, currentIDX) {
				continue
			}

			// if we found Number Gather it 

			numbersFound += 1
			if numbersFound > 2{
				return 0
			}
			before, after, jOffset := gatherNumber(matrix[newX], newY, cols)

			// Handle jOffset so we don't double count numbers
			// This will happen only on UP and DOWN
			// e.g.
			// 11111
			// ..*..
			// 11111

			for k := 0; k < jOffset + 1; k++{
				point := Point{newX, k + newY}
				addToVisitedSet(visitedSet, point)
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

//!
func addToVisitedSet(visitedSet map[string]struct{}, point Point) {
    key := fmt.Sprintf("(%d,%d)", point.X, point.Y)
    visitedSet[key] = struct{}{}
}
//!
func isInVisitedSet(visitedSet map[string]struct{}, point Point) bool {
    key := fmt.Sprintf("(%d,%d)", point.X, point.Y)
    _, exists := visitedSet[key]
    return exists
}




func checkLimits( idx int, limit int) bool {
	return -1 < idx && idx < limit
}


func isSymbol(char string) bool{
	_, err := strconv.Atoi(char)
	return err != nil && char != "." 
}


func isGear(char string) bool{
	return char == "*" 
}



func gatherNumber(matrix []string, idx int, size int) (string, string, int) {
	// go back until not number
	before := ""
	after := ""
	jOffset := 0
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

//!
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
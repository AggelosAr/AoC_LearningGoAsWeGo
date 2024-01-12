package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"time"
	"strings"
)


func main() {
	rawData := readData()
	data := formatData(rawData)

	resultPart1 := solver(data, "Last")
	fmt.Println("The result (PART1) : ", resultPart1)

	/*
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Printf("Res <%d> -- Time took %s\n", res, elapsed)
	*/

	resultPart2 := solver(data, "First")
	fmt.Println("The result (PART2) : ", resultPart2)
	if resultPart1 == 2075724761 && resultPart2 == 1072 {
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

func formatData(text string) [][] int{

	lines := strings.Split(text, "\n")
	data := [][] int {}

	for idx := 0; idx < len(lines); idx++ {
		if lines[idx] != "" {
			data = append(data, formatLine(lines[idx]))
		}
	}
	return data
}


func formatLine(line string) [] int {

	parts := strings.Fields(line)
	numbers := [] int {}

	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing number in line : ", err)
		} else {
			numbers = append(numbers, number)
		}
	}
	return numbers
}


func solver(data [][] int, positionToADD string) int {
	result := 0
	for _, row := range data {
		result += solveCurrentRow(row, positionToADD)
	}
	return result
}


func solveCurrentRow(row []int, positionToADD string) int {
	
	position := getPosition(len(row), positionToADD)
	requiredNumbers := []int {row[position]}
	currentLine := row
	
	for {
		nextLine, canStop := getNextLine(currentLine)
		
		if canStop {
			break
		}
		
		currentLine = nextLine
		position = getPosition(len(currentLine), positionToADD)
		requiredNumbers = append(requiredNumbers, currentLine[position])
		
	}

	return getAnswer(requiredNumbers, positionToADD)
}


func getPosition(size int, positionToADD string) int {
	position := 0
	if positionToADD  == "Last" {
		position = size - 1
	}
	return position
}


func getAnswer(numbers []int, positionToADD string) int{
	if positionToADD == "First" {
		// here we cant just sum the numbers 
		// we must go backwards and calculate the next on each step
		return extrapolateBackwards(numbers)
	} 
	return sumArray(numbers)
}


func sumArray(arr []int) int {
	result := 0
	for _, num := range arr {
		result += num
	}
	return result
}

func extrapolateBackwards(nums []int ) int {
	from := 0
	to := 0

	for idx := len(nums) - 1; idx > - 1; idx-- {
		to = nums[idx]
		from = to - from
	}
	return from
} 


func getNextLine(currentLine [] int) ([]int, bool) {
	canStop := true
	nextLine := []int {}
	diff := 0

	for i := 0; i < len(currentLine) - 1; i++ {
		diff = currentLine[i + 1] - currentLine[i]

		if diff != 0 {
			canStop = false
		}

		nextLine = append(nextLine, diff)
	}

	return nextLine, canStop
}



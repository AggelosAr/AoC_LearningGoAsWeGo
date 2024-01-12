package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {

	data := readData()
	times, distances := parseData(data)

	resultPart1 := solverPart1(times, distances)
	fmt.Println("The result PART1 : ", resultPart1)

	start := time.Now()
	resultPart2 := solverPart2(times, distances)
	elapsed := time.Since(start)
	fmt.Println("The result PART2 : ", resultPart2)
	fmt.Printf("Brute Force PART2 took %s\n", elapsed)

	start = time.Now()
	resultPart2Bin := solverPart2WithBinarySearch(times, distances)
	elapsed = time.Since(start)
	fmt.Println("The result PART2 (Bin search): ", resultPart2Bin)
	fmt.Printf("Binary Search PART2 took %s\n", elapsed)

	if resultPart2Bin == resultPart2Bin {
		if resultPart2Bin == 45647654 {
			fmt.Println(" N I C E ")
		}
	} else {
		fmt.Println(" B A D ")
	}
}


func getRootPath() string{
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


func parseData(text string) ([]int, []int) {
	lines := strings.Split(text, "\n")
	return extractNumbers(lines[0]), extractNumbers(lines[1])
}


func extractNumbers(text string) []int {

	numberPart := strings.Split(text, ":")[1]
	extractedNumbers := []int {}

	for _, numberText := range strings.Fields(numberPart) {

		number, err := strconv.Atoi(numberText)

		if err != nil {
			continue
		}
		extractedNumbers = append(extractedNumbers, number)
	}
	return extractedNumbers
}


func solverPart1(times []int, distances []int) int {

	totalWays := 1

	for race := 0; race < len(times); race++ {
		totalWays *= findWays(times[race], distances[race])
	}
	return totalWays
}


func findWays(time int, distance int) int {

	ways := 0
	reachedWin := false

	for speed := 0; speed < time; speed++ {
		
		potentialDistance := getDistanceTravelled(speed, time, distance)
		if potentialDistance > distance {
			ways += 1
			reachedWin = true
		} else if reachedWin {
			break
		}
	}
	return ways
}

// can be solved with binary search to search the left and right limits
// 2 bin searches ->  1 for left min limit and 1 for right max limit
// Another approach is -> 1 bin search and then serial search to left and right
// 2 bin seaches is supposed to be better


func solverPart2(times []int, distances []int) int {

	singletTime := concatIntArrayToSingleNumber(times)
	singleDistance := concatIntArrayToSingleNumber(distances)
	
	return findWays(singletTime, singleDistance)
}


func concatIntArrayToSingleNumber( arr []int ) int {
	stringArr := [] string {}

	for _, num := range arr {
		stringArr = append(stringArr, strconv.Itoa(num))
	}

	singleStr := strings.Join(stringArr, "")

	resultNum, err := strconv.Atoi(singleStr)
	
	if err != nil {
		fmt.Println("Error parsing text to int : ", err)
		return 0
	}
	return resultNum
}


func solverPart2WithBinarySearch(times []int, distances []int) int {

	singletTime := concatIntArrayToSingleNumber(times)
	singleDistance := concatIntArrayToSingleNumber(distances)
	return findWaysWithBinSearch(singletTime, singleDistance)
}


func findWaysWithBinSearch(time int, distance int) int {
	leftLimit := binarySearchLow(time, distance)
	rightLimit := binarySearchHight(time, distance)
	return rightLimit - leftLimit + 1
}


func binarySearchLow(time int, targetDistance int) int{

	lowSpeed := 0
	hightSpeed := time

	for lowSpeed <= hightSpeed {
		//currentSpeed := (lowSpeed + hightSpeed) / 2
		currentSpeed := lowSpeed + ( hightSpeed - lowSpeed ) / 2

		currentDistance := getDistanceTravelled(currentSpeed, time, targetDistance)
		
		if currentDistance < targetDistance {
			lowSpeed = currentSpeed + 1
			
		} else {
			hightSpeed = currentSpeed - 1
		}
	}
	return lowSpeed
}


func binarySearchHight(time int, targetDistance int) int{

	lowSpeed := 0
	hightSpeed := time

	for lowSpeed <= hightSpeed {
		//currentSpeed := (lowSpeed + hightSpeed) / 2
		currentSpeed := lowSpeed + ( hightSpeed - lowSpeed ) / 2

		currentDistance := getDistanceTravelled(currentSpeed, time, targetDistance)
		
		if currentDistance < targetDistance {
			hightSpeed = currentSpeed - 1
			
		} else {
			lowSpeed = currentSpeed + 1
		}
	}
	return hightSpeed
}


func getDistanceTravelled(speed int, time int, distance int) int {

	timeRemaining := time - speed
	potentialDistance := speed * timeRemaining
	return potentialDistance
}

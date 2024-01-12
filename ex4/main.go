package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {

	
	resultPart1 := solverPart1()
	fmt.Println("The result is PART1 : ", resultPart1)

	resultPart2 := solverPart2()
	fmt.Println("The result is PART2 : ", resultPart2)
}


func getDataPath() string{
	wd, _ := os.Getwd()
	return wd + "\\data\\input.txt"
}


func solverPart1() int {

	file, err := os.Open(getDataPath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	
	for scanner.Scan() {
		
		line := scanner.Text()
		lottery := formatLine(line)
		matches := getMatches(lottery)

		result += computeScore(matches)
		
	}

	return result
}



type Lottery struct {
	winning map[int]struct{}//left
	current map[int]struct{}//right
}


func formatLine(line string) Lottery{
	numbers := strings.Split(strings.Split(line, ":")[1], "|")
	currentLottery := Lottery{winning: map[int]struct{}{}, current: map[int]struct{}{}}
	popupateLotteryField(currentLottery.winning, numbers[0])
	popupateLotteryField(currentLottery.current, numbers[1])
	return currentLottery
}

func popupateLotteryField(field  map[int]struct{}, text string) {

	numbers := strings.Split(strings.TrimSpace(text), " ")
	//fmt.Println(numbers)
	for _, numberStr := range numbers {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
			// spaces are counted as zero !!
		}
		field[number] = struct{}{}
	}

}

func getMatches(lottery Lottery) int{
	// Must check if any of the current numbers are included in the winning
	// e.g. intersection of the 2 sets
	matches := getIntersection(lottery)
	return matches
}

func computeScore(matches int) int{
	// each match is worth 1 
	// then it doubles for each match // e.g. 2*(matches-1)
	result := int(math.Pow(float64(2), float64(matches - 1)))
	return result
}


func getIntersection(lottery Lottery) int {
	matches := 0

	for k := range lottery.winning {
		_, exists := lottery.current[k]
		if exists {
			matches++
		}
	}

	return matches
}



func solverPart2() int {

	file, err := os.Open(getDataPath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	multipliers := map[int]int {}
	cardIdx := 1

	for scanner.Scan() {
		
		line := scanner.Text()
		lottery := formatLine(line)
		matches := getMatches(lottery)

		_, exists := multipliers[cardIdx]
		if !exists {
			multipliers[cardIdx] = 1
		} else {
			multipliers[cardIdx]++
		}
		
		if matches > 0 {
		// First we must see how many cards of the current card we have e.g.currentMult
		// So we can multiply the next cards by that number
			multiplier := multipliers[cardIdx]
			for i := 1; i < matches + 1; i++{
				_, exists := multipliers[cardIdx + i]
				if !exists{
					multipliers[cardIdx + i] = multiplier
				}else{
					multipliers[cardIdx + i] += multiplier 
				}
				
			}
		} 

		cardIdx++
	}

	
	// We need to add all the multiplier together
	result := 0
	for _, v := range multipliers {
		result += v
	}
	return result
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	
	"aggelos.com/go/aoc/ex4/utils"
)




func main() {
	lotteries := getData()
	
	resultPart1 := solverPart1(lotteries)
	fmt.Println("The result is PART1 : ", resultPart1)

	resultPart2 := solverPart2(lotteries)
	fmt.Println("The result is PART2 : ", resultPart2)

	if resultPart1 == 20407 && resultPart2 == 23806951 {
		fmt.Println(" G O O D ")
	}
}



func solverPart1(lotteries []utils.Lottery) int {
	r := 0
	for _, l := range lotteries {
		r += l.ComputeScore()
	}
	return r
}



func solverPart2(lotteries []utils.Lottery) int {

	multipliers := map[int]int{}
	
	for idx, l := range lotteries {
		
		_, exists := multipliers[idx]
		if !exists {
			multipliers[idx] = 1
		} else {
			multipliers[idx]++
		}
		
		if l.Matches > 0 {
		// First we must see how many cards of the current card we have e.g.currentMult
		// So we can multiply the next cards by that number
			multiplier := multipliers[idx]
			for i := 1; i < l.Matches + 1; i++{
				_, exists := multipliers[idx + i]
				if !exists{
					multipliers[idx + i] = multiplier
				}else{
					multipliers[idx + i] += multiplier 
				}
			}
		} 
	}

	// We need to add all the multiplier together
	r := 0
	for _, v := range multipliers {
		r += v
	}
	return r
}



func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []utils.Lottery {

	file, err := os.Open(getFilePath())

	if err != nil {
		fmt.Println("Error opening the file : ", err)
	}

	scanner := bufio.NewScanner(file)
	lotteries := []utils.Lottery{}

	for scanner.Scan() {
		l := *utils.NewLottery(scanner.Text())
		lotteries = append(lotteries, l)
	}
	
	return lotteries
}
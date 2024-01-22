package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"aggelos.com/go/aoc/ex2/utils"
)

func main() {

	games := getData()
	
	start := time.Now()
	part1result := solvePart1(games)
	elapsed := time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART1 - RESULT - ", part1result)
	
	start = time.Now()
	part2result := solvePart2(games)
	elapsed = time.Since(start)
	fmt.Printf("Time took %s\n", elapsed)
	fmt.Println("PART2 - RESULT - ", part2result)
	
	if part1result == 2176 && part2result == 63700 {
		fmt.Println(" G O O D ")
	} else {
		fmt.Println(" B A D ")
	}

}


func solvePart1(games []utils.Game) int {
	r := 0
	for id, game := range games {
		if checkGame(game){
			r += id + 1
		}
	}
	return r
}


func solvePart2(games []utils.Game) int {
	r := 0
	for _, game := range games {
		r += findPower(game)
	}
	return r
}


func checkGame(game utils.Game) bool {

	var (
		redMust = 12
		greenMust = 13
		blueMust = 14
	)
	
	for _, turn := range game.Turns {
		if turn.R > redMust{
			return false
		}else if turn.G > greenMust{
			return false
		}else if turn.B > blueMust{
			return false
		}
	}

	return true
}



func findPower(game utils.Game) int {
	
	colors := utils.NewColors()

	for _, turn := range game.Turns {
		colors.R = max(colors.R, turn.R)
		colors.G = max(colors.G, turn.G)
		colors.B = max(colors.B, turn.B)
		
	}
	return colors.R * colors.G * colors.B
}


func getFilePath() string {
	rootPath, _ := os.Getwd()
	return filepath.Join(rootPath, "data", "input.txt")
}


func getData() []utils.Game { 

	file, err := os.Open(getFilePath())
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	games := []utils.Game {}

	for scanner.Scan() {
		line := scanner.Text()
		
		game := utils.NewGame(line)
		games = append(games, *game)
	}

	return games
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rootPath, _ := os.Getwd()
	rootPath += "\\data\\input.txt"

	file, err := os.Open(rootPath)
	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {

		line := scanner.Text()
		
		id, game := formatLineToArray(line)
		/* PART 1
		if checkGame(game){
			result += id
			fmt.Println(id, game)
		}
		*/
		// PART 2
		id ++ // silence it 
		result += findPower(game)

	}
	fmt.Printf("The result is : %d\n", result)

}


func formatLineToArray(line string) (int, []string) {
	tempData := strings.Split(line, ":")
	tempId := strings.Split(tempData[0], " ")[1]
	id, _ := strconv.Atoi(tempId)
	gameData := tempData[1]
	game := strings.Split(gameData, ";")
	return id, game
}

func checkGame(game []string) bool {
	redMust := 12
	greenMust := 13
	blueMust := 14

	// [ "3 blue, 4 red" ,  "3 blue, 4 red" ]
	for _, turn := range game {
		// ["3 blue", "4 red"]
		
		balls := strings.Split(turn, ",")
		
		for _, ball := range balls {
			trimmedBall := strings.TrimSpace(ball)
			tempBall := strings.Split(trimmedBall, " ")
			
			number, _ := strconv.Atoi(tempBall[0])
			color := tempBall[1]

			if color == "red" && number > redMust{
				return false
			}else if color == "green" && number > greenMust{
				return false
			}else if color == "blue" && number > blueMust{
				return false
			}
		}

	}

	return true
}


type maxColors struct {
	red int
	green int
	blue int 
}
func findPower(game []string) int {
	

	currentColors := maxColors{}
	// [ "3 blue, 4 red" ,  "3 blue, 4 red" ]
	for _, turn := range game {
		// ["3 blue", "4 red"]
		
		balls := strings.Split(turn, ",")

		
		
		for _, ball := range balls {
			trimmedBall := strings.TrimSpace(ball)
			tempBall := strings.Split(trimmedBall, " ")
			
			number, _ := strconv.Atoi(tempBall[0])
			color := tempBall[1]

			if color == "red"{
				currentColors.red = max(number, currentColors.red)
			}else if color == "green" {
				currentColors.green = max(number, currentColors.green)
			}else if color == "blue"{
				currentColors.blue = max(number, currentColors.blue)
			}
		}

	}

	return currentColors.red * currentColors.green * currentColors.blue
}
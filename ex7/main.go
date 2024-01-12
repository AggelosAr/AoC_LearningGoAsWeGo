package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	handutils "aggelos.com/go/aoc/ex7/handUtils"
)


func main() {
	data := readData()
	hands := formatData(data)

	resultPart1 := solver(hands, false)
	fmt.Println("The result (PART1) : ", resultPart1)

	
	resultPart2 := solver(hands, true)
	fmt.Println("The result (PART2) : ", resultPart2)

	if resultPart1 == 250898830 && resultPart2 == 252127335 {
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


func formatData(text string) []handutils.Hand {
	rawHands := strings.Split(text, "\n")
	return formatHands(rawHands)
}


func formatHands(rawHands []string) []handutils.Hand {

	var hands []handutils.Hand 

	currentHand := handutils.Hand{}
	for _, hand := range rawHands {

		handParts := strings.Fields(hand)
		// Ignore the last new Line
		if len(handParts) == 0 {
			break
		}

		currentHand.Cards = handParts[0]
		num, _ := strconv.Atoi(handParts[1])

		currentHand.Bid = num

		hands = append(hands, currentHand)
	}
	return hands
}


func solver(hands []handutils.Hand, isJWild bool) int {

	types := map[int][]handutils.Hand {}

	for _, hand := range hands {
		currentType := hand.GetStrength(isJWild)
		types[currentType] = append(types[currentType], hand)
	}

	// Hands should be ranked from weakest to strongest
	winnings := 0
	rank := 1

	for handType := 0; handType < 7; handType++ {

		currentHands, _ := types[handType]
		currentWinnings := 0
		
		if len(currentHands) == 1 {
			//fmt.Println(currentHands[0], rank)
			currentWinnings += rank * currentHands[0].Bid
			rank++
		} else if len(currentHands) > 1 {

			sort.Sort(handutils.ByCards{Hands: currentHands, IsJWild: isJWild})
			for _, hand := range currentHands {
				//fmt.Println(hand, rank)
				currentWinnings += rank * hand.Bid
				rank++
			}
		}	
		winnings += currentWinnings
	}
	return winnings
}
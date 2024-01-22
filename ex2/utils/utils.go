package utils

import (
	"strconv"
	"strings"
)

type Game struct {
	Turns []Colors
}

func NewGame(line string) *Game {

	newGame := Game{
		Turns: []Colors{},
	}

	parts := strings.Split(line, ":")

	for _, turn := range strings.Split(parts[1], ";") {

		colors := NewColors()
		for _, ball := range strings.Split(turn, ",") {

			currBall := strings.Split(strings.TrimSpace(ball), " ")
			number, _ := strconv.Atoi(currBall[0])
			color := currBall[1]

			if color == "red" {
				colors.R = number
			} else if color == "green" {
				colors.G = number
			} else if color == "blue" {
				colors.B = number
			}
		}
		
		newGame.Turns = append(newGame.Turns, *colors)
	}

	return &newGame
}



type Colors struct {
	R, G, B int
}

func NewColors() *Colors {
	return &Colors{}
}

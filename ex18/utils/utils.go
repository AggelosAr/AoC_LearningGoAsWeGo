package utils

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}


func (p Point) IsTarget(sizeX, sizeY int) bool {
	return p.X == sizeX - 1 && p.Y == sizeY - 1
}

func (p Point) Add(pointB Point) Point {
	return Point{X: p.X + pointB.X, Y: p.Y + pointB.Y}
}


func (p *Point) AddInPlace(pointB Point) {
	p.X += pointB.X
	p.Y += pointB.Y
}


func (p Point) CrossPoints(p2 Point) int {
	return p.X * p2.Y - p.Y * p2.X
}

func (p Point) IsValid(sizeX, sizeY int) bool {
	return (-1 < p.X && p.X < sizeX) && (-1 < p.Y && p.Y < sizeY)
}


type Cell struct {
	Direction string
	Distance int
	Color string
}

func GetNewCell(text string) Cell {
	
	parts := strings.Fields(text)
	num, _ := strconv.Atoi(parts[1])
	newCell := Cell{
		Direction: parts[0],
		Distance: num,
		Color: strings.TrimRight(strings.TrimLeft(parts[2], "("), ")"),
	}
	return newCell
}

func GetNewCell2(text string) Cell {
	
	color := strings.TrimRight(strings.TrimLeft(strings.Fields(text)[2], "("), ")")
	direction, distance := decodeColor(color)

	newCell := Cell{
		Direction: direction,
		Distance: distance,
		Color: color,
	}
	
	return newCell
}

func decodeColor(color string) (string, int) {

	direction, _ := strconv.ParseInt(color[6:], 16, 64)
	distance, _ := strconv.ParseInt(color[1:6], 16, 64)

	//fmt.Printf("%v -> %v, %v\n", color, dDistance, directionMapper[int(dDirection)])
	//fmt.Println("__________________________")
	//newDistance := int(distance) / 10000
	//distance = int64(newDistance)
	
	return directionMapper[int(direction)], int(distance)
}


func (c Cell) GetPoint() Point {
	newPoint := Point{}

	return newPoint
}

var Offsets = map[string]Point{
	"U" : {X: -1, Y: 0},
	"D" : {X: 1, Y: 0},
	"L" : {X: 0, Y: -1},
	"R" : {X: 0, Y: 1},
}

var Rotations = map[string][]string{
	"U" : { "U", "L", "R"},
	"D" : { "D", "L", "R"},
	"L" : { "U", "D", "L"},
	"R" : { "U", "D", "R"},
}

var directionMapper = map[int]string {
	0: "R",
	1: "D", 
	2: "L",
	3: "U",
}

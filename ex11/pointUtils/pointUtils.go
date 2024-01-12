package pointUtils

import (
	"container/list"
	"fmt"
)

type Point struct {
	X, Y int
}

func (p Point) IsValid(sizeX int, sizeY int) bool {
	return (-1 < p.X && p.X < sizeX) && (-1 < p.Y && p.Y < sizeY)
}

func (p Point) Add(pointB Point) Point {
	return Point{X: p.X + pointB.X, Y: p.Y + pointB.Y}
}


func (p Point) Diff(pointB Point) ([2]int, [2]int) {
	rows := [2]int {min(p.X, pointB.X), max(p.X, pointB.X)}
	cols := [2]int {min(p.Y, pointB.Y), max(p.Y, pointB.Y)}
	return rows, cols
}

func (p Point) Compare(pointB Point) bool {
	return p.X == pointB.X && p.Y == pointB.Y
}

func (p Point) GetCrossPoints(sizeX int, sizeY int) []Point {
	crossPoints := []Point{}

	for _, offset := range offsets {
		currentPoint := p.Add(offset)
		if currentPoint.IsValid(sizeX, sizeY) {
			crossPoints = append(crossPoints, currentPoint)
		}
	}

	return crossPoints
}

func (p Point) Key() string {
	return fmt.Sprintf("(%d-%d)", p.X, p.Y)
}

var offsets = map[string]Point{
	"UP":    {X: -1, Y: 0},
	"DOWN":  {X: 1, Y: 0},
	"LEFT":  {X: 0, Y: -1},
	"RIGHT": {X: 0, Y: 1},
}


func GetDoublyList() PointDoublyList{
	return PointDoublyList{List: list.New()}
}

type PointDoublyList struct {
	List *list.List
}


func (l *PointDoublyList) Clear(){
	l.List.Init()
}

func (l *PointDoublyList) PopLeft() Point{
	frontElement := l.List.Front()
	l.List.Remove(frontElement) 
	return frontElement.Value.(Point)
}

func (l *PointDoublyList) PushBack(point Point) {
	l.List.PushBack(point)
}

func (l *PointDoublyList) HasMore() bool {
	return l.List.Len() > 0
}
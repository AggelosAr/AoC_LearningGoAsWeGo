package pointUtils

import (
	"errors"
	"container/list"
	"fmt"
)

type Point struct {
	X, Y int
	Symbol string
}


func (p Point) CrossPoints(p2 Point) int {
	return p.X * p2.Y - p.Y * p2.X
}

func (p Point) ToKey() string {
	return fmt.Sprintf("(%d-%d)", p.X, p.Y)
}


func (p Point) Add(offset Offset) Point{
	return Point{X: p.X + offset.X, Y: p.Y + offset.Y}
}


func (p Point) Validate(rows int, cols int) error {
	if p.X < 0 || p.X >= rows {
		return errors.New("Out of range")
	}
	if p.Y < 0 || p.Y >= cols {
		return errors.New("Out of range")
	}
	return nil
}


type BorderPoint struct {
	Point Point
	In []Point
}


type BorderPoints struct {
	Points map[string]BorderPoint
}


func (bp *BorderPoints) Add(newBorderPoints []BorderPoint) {
	for _, currentBp := range newBorderPoints {
		bp.Points[currentBp.Point.ToKey()] =  currentBp
	}
}

func (p Point) Compare(pointB Point) bool {
	return p.X == pointB.X && p.Y == pointB.Y
}


type Offset struct {
	X, Y int
}


var DirectionsMap = map[string][]string {
	"S" : {"UP", "RIGHT", "DOWN", "LEFT"},
	"|" : {"UP", "DOWN"},
	"-" : {"RIGHT", "LEFT"},
	"L" : {"UP", "RIGHT"},
	"J" : {"UP", "LEFT"},
	"7" : {"DOWN", "LEFT"},
	"F" : {"RIGHT", "DOWN"},
}

var DirectionsToOffset = map[string]Offset {
	"UP" : {X: -1, Y: 0},
	"DOWN" : {X: 1, Y: 0},
	"LEFT": {X: 0, Y: -1},
	"RIGHT": {X: 0, Y: 1},
}


type PointDoublyLinkedList struct {
	List *list.List
}


func NewPointQ() *PointDoublyLinkedList {
	return &PointDoublyLinkedList{List: list.New()}
}


func (q *PointDoublyLinkedList) Add(v Point) {
	q.List.PushBack(v)
}


func (q *PointDoublyLinkedList) PopLeft() (Point, error){

	if q.List.Len() == 0 {
		return Point{}, errors.New("Q empty")
	}

	element  := q.List.Front()
	value  := element.Value.(Point)
	q.List.Remove(element)

	return value, nil
} 


// chech chat gpt for line if this is valid GO code 
// valid in the case GOOD 
// TODO 
type Line struct {
	Points []Point
}

func GetNewLine() Line{

	return Line{Points: []Point {}}
}

func (l Line) Last() Point {
	return l.Points[len(l.Points) - 1]
}

func (l Line) IsComplete() bool { 
	_, exists := TurningPoints[l.Last().Symbol]
	return exists
}

func (l Line) Clear() {
	l.Points = []Point {}
}

func (l *Line) Add(point Point) {
	l.Points = append(l.Points, point)
}

func (l Line) GetBorderPoints() []BorderPoint {
	borderPoints := []BorderPoint {}
	currentBp := BorderPoint{}
	for _, point := range l.Points {
		currentBp.Point = point
		currentBp.In = []Point {point}
		borderPoints = append(borderPoints, currentBp)

	}
	return borderPoints
}


var TurningPoints = map[string]struct{} {
	"7" : {},
	"F" : {},
	"J" : {},
	"L" : {},
	"S" : {},
}

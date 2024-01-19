package utils

import (
	"container/list"
	"errors"
	
)

type Point struct {
	X, Y int
}




func (p Point) Add(offset Point) Point {
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




var DirectionsToOffset = map[string]Point{
	"UP":    {X: -1, Y: 0},
	"DOWN":  {X: 1, Y: 0},
	"LEFT":  {X: 0, Y: -1},
	"RIGHT": {X: 0, Y: 1},
}

var reverseDir = map[string]string{
	"UP":    "DOWN",
	"DOWN":  "UP",
	"LEFT":  "RIGHT",
	"RIGHT": "LEFT",
}

type Q struct {
	List *list.List
}

func (q *Q) HasNext() bool {
	return q.List.Len() > 0
}

func NewPointQ() *Q {
	return &Q{List: list.New()}
}

func (q *Q) Add(v Point) {
	q.List.PushBack(v)
}

func (q *Q) PopLeft() (Point, error) {

	if q.List.Len() == 0 {
		return Point{}, errors.New("Q empty")
	}

	element := q.List.Front()
	value := element.Value.(Point)
	q.List.Remove(element)

	return value, nil
}

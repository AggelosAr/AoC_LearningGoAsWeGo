package utils

import (
	"container/heap"

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

func (p Point) IsValid(sizeX, sizeY int) bool {
	return (-1 < p.X && p.X < sizeX) && (-1 < p.Y && p.Y < sizeY)
}

type State struct {
	Point    Point
	History  string
	Streak int
}



func (c State) GetNextStates(minMoves, maxMoves, rows, cols int) []State {
	
	nextStates := []State {}

	for _, direction := range Rotations[c.History] {

		nextPoint := c.Point.Add(Offsets[direction])
		if !nextPoint.IsValid(rows, cols) {
			continue
		}

		newStreak := 0
		
		if c.Streak < minMoves {
			if c.History != direction {
				continue
			} else {
				newStreak = c.Streak + 1
			}
		} else {
			if c.Streak < maxMoves {
				if c.History != direction {
					newStreak = 1
				} else {
					newStreak = c.Streak + 1
				}
			} else {
				if c.History == direction {
					continue
				} else {
					newStreak = 1
				}
			}
		}

		newState := State{
			Point:   nextPoint,
			History: direction,
			Streak:  newStreak,
		}
		nextStates = append(nextStates, newState)
	}
	return nextStates
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



type pqi[T any] struct {
	v T
	p int
}
type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }
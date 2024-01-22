package utils

import (
	"errors"
	"fmt"
)

type Point struct {
	X, Y int
}


func (p Point) Compare(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) ToKey() string {
	return fmt.Sprintf("(%d-%d)", p.X, p.Y)
}


func (p Point) Add(offset Point) Point{
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


type Graph struct {
	Data map[Point]map[Point]int
}

func (g *Graph) AddConnection(p1, p2 Point, length int) {

	_, p1MapExists := g.Data[p1]
	//_, p2MapExists := g.Data[p2]

	if !p1MapExists {
		g.Data[p1] = map[Point]int {}
	}

	//if !p2MapExists {
	//	g.Data[p2] = map[Point]int {}
	//}

	g.Data[p1][p2] = length
	//g.Data[p2][p1] = length
}





func (g Graph) Dfs(point, target Point, visitedMap map[Point]struct{}) int {

	
	if point == target {
		return 0
	}

	visitedMap[point] = struct{}{}
	distance := 0

	for nextPoint := range g.Data[point] {

		_, isVisited := visitedMap[nextPoint]
		if isVisited {
			continue
		}
		
		cDistance := g.Data[point][nextPoint]
		totalDistance := cDistance + g.Dfs(nextPoint, target, visitedMap)
		distance = max(distance, totalDistance)
	}

	delete(visitedMap, point) // backtrack
	return distance
}


func GetGraphFromGrid(grid [][]string, sPoint Point) Graph {

	newGraph := Graph{}
	newGraph.Data = map[Point]map[Point]int{}
	visitedMap := map[Point]struct{} {}

	newGraph.FillGraph(sPoint, sPoint, 0, grid, visitedMap)
	// 5 3
	fmt.Println(newGraph.Data[Point{X: 0, Y: 1}])
	fmt.Println(newGraph.Data[Point{X: 5, Y: 4}]) // 3 kati  - 4 kati
	fmt.Println(newGraph.Data[Point{X: 6, Y: 3}])
	return newGraph
}



// with shorten Paths
func (g *Graph) FillGraph(sPoint, point Point, length int, grid [][]string, visitedMap map[Point]struct{}) {
	
	rows := len(grid)
	cols := len(grid[0])
	
	_, isVisited := visitedMap[point]
	if isVisited {
		return
	}
	visitedMap[point] = struct{}{}


	nextPoints := []Point {}
	for _, offset := range DirectionsToOffset {

		nextPoint := point.Add(offset)
		
		if nextPoint.Validate(rows, cols) != nil {
			continue
		}
		if grid[nextPoint.X][nextPoint.Y] == "#" {
			continue
		}
		_, isVisited := visitedMap[nextPoint]
		if isVisited {
			continue
		} 
		nextPoints = append(nextPoints, nextPoint)
	}

	if len(nextPoints) == 1 {
		g.FillGraph(sPoint, nextPoints[0], length + 1, grid, visitedMap)
	} else {

		for _, nextPoint := range nextPoints {
			g.AddConnection(sPoint, nextPoint, length + 1)	
			g.FillGraph(nextPoint, nextPoint, 0, grid, visitedMap)
		}
		
	}

}


var AvailableDirections = map[string][]string {
	"." : {"U", "D", "L", "R"},
	">" : {"R"}, 
	"<" : {"L"},
	"^" : {"U"},
	"v" : {"D"}, 
}


var DirectionsToOffset = map[string]Point {
	"U" : {X: -1, Y: 0},
	"D" : {X: 1, Y: 0},
	"L": {X: 0, Y: -1},
	"R": {X: 0, Y: 1},
}

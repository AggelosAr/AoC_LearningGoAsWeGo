package utils



type Point struct {
	X, Y int
}


func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p Point) IsValid (rows int, cols int) bool {
	if p.X < 0 || p.X >= rows {
		return false
	}
	if p.Y < 0 || p.Y >= cols {
		return false
	}
	return true
}


var Offsets = [8][2]int {{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1}, 
	{1, -1}, {1, 0}, {1, 1}}
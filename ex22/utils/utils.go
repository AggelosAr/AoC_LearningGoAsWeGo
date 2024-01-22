package utils

import "fmt"

type Brick struct {
	Start Coord
	End   Coord
	Id int
}

func (b Brick) String() string {
    return fmt.Sprintf("BRICK <%d> Start: %s, End: %s", b.Id, b.Start, b.End)
}

func GetNewBrick(id, x1, y1, z1, x2, y2, z2 int) Brick {
	newBrick := Brick{Id: id}
	coordStart := Coord{X: x1, Y: y1, Z: z1}
	coordEnd := Coord{X: x2, Y: y2, Z: z2}

	newBrick.Start = coordStart
	newBrick.End = coordEnd

	return newBrick
}

type Coord struct {
	X, Y, Z int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.X, c.Y, c.Z)
}
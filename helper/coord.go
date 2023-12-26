package helper

import (
	"fmt"
	"math"
)

type Coord struct {
	X int
	Y int
}

const (
	NORTH = iota
	EAST  = iota
	SOUTH = iota
	WEST  = iota
)

// Return all coordinate neighbours
func (c Coord) GetNeighbours(diagonal bool) []Coord {
	ret := []Coord{
		Coord{c.X - 1, c.Y},
		Coord{c.X, c.Y - 1},
		Coord{c.X, c.Y + 1},
		Coord{c.X + 1, c.Y},
	}
	if diagonal {
		ret = append(ret, []Coord{
			Coord{c.X - 1, c.Y + 1},
			Coord{c.X - 1, c.Y - 1},
			Coord{c.X + 1, c.Y + 1},
			Coord{c.X + 1, c.Y - 1},
		}...)
	}
	return ret
}

func (c Coord) GetSafeNeighbours(diagonal bool, xlen int, ylen int) []Coord {
	n := c.GetNeighbours(diagonal)
	r := make([]Coord, 0)
	for _, neighbour := range n {
		if neighbour.X < 0 || neighbour.X >= xlen || neighbour.Y < 0 || neighbour.Y >= ylen {
			continue
		}
		r = append(r, neighbour)
	}
	return r
}

func (c Coord) MoveDirection(dir int) Coord {
	switch dir {
	case NORTH:
		return Coord{X: c.X, Y: c.Y - 1}
	case EAST:
		return Coord{X: c.X + 1, Y: c.Y}
	case WEST:
		return Coord{X: c.X - 1, Y: c.Y}
	case SOUTH:
		return Coord{X: c.X, Y: c.Y + 1}
	}
	return c
}

func (c Coord) Move(dir int) Coord {
	r := c
	switch dir {
	case NORTH:
		r.Y -= 1
	case EAST:
		r.X += 1
	case WEST:
		r.X -= 1
	case SOUTH:
		r.Y += 1
	}
	return r
}

func (c Coord) GetOrderedSquare() []Coord {
	ret := []Coord{
		Coord{c.X - 1, c.Y - 1},
		Coord{c.X, c.Y - 1},
		Coord{c.X + 1, c.Y - 1},
		Coord{c.X - 1, c.Y},
		Coord{c.X, c.Y},
		Coord{c.X + 1, c.Y},
		Coord{c.X - 1, c.Y + 1},
		Coord{c.X, c.Y + 1},
		Coord{c.X + 1, c.Y + 1},
	}
	return ret
}

// Return all coordinate neighbours, excluding negatives
func (c Coord) GetNeighboursPos(diagonal bool) []Coord {
	ret := make([]Coord, 0)
	ret = append(ret, Coord{c.X, c.Y + 1})
	ret = append(ret, Coord{c.X + 1, c.Y})
	if diagonal {
		ret = append(ret, Coord{c.X + 1, c.Y + 1})
	}

	if c.X > 0 {
		ret = append(ret, Coord{c.X - 1, c.Y})
		if diagonal {
			ret = append(ret, Coord{c.X - 1, c.Y + 1})
			if c.Y > 0 {
				ret = append(ret, Coord{c.X - 1, c.Y - 1})
			}
		}
	}
	if c.Y > 0 {
		ret = append(ret, Coord{c.X, c.Y - 1})
		if diagonal {
			ret = append(ret, Coord{c.X + 1, c.Y - 1})
		}
	}
	return ret
}

func (c Coord) ToString() string {
	return fmt.Sprintf("%v,%v", c.X, c.Y)
}

func ManhattanDistance(c1 Coord, c2 Coord) int {
	return int(math.Abs(float64(c1.X-c2.X)) + math.Abs(float64(c1.Y-c2.Y)))
}

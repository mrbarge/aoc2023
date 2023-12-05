package helper

import "math"

type Coord struct {
	X int
	Y int
}

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

func ManhattanDistance(c1 Coord, c2 Coord) int {
	return int(math.Abs(float64(c1.X - c2.X)) + math.Abs(float64(c1.Y - c2.Y)))
}
package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
)

type Beam struct {
	c   helper.Coord
	dir int
}

func simulate(b Beam, input []string) int {
	beams := []*Beam{&b}

	seenTiles := make(map[string]int)
	for len(beams) > 0 {
		// iterate backwards so we can remove beams
		for i := len(beams) - 1; i >= 0; i-- {
			beam := beams[i]

			newbeams := make([]*Beam, 0)

			// Record beam's presence on a tile
			key := fmt.Sprintf("%v-%v", beam.c.X, beam.c.Y)

			// Have we been here before in this direction?
			if v, ok := seenTiles[key]; ok {
				if v == beam.dir {
					// we have, stop tracking and move on
					beams = append(beams[:i], beams[i+1:]...)
					continue
				}
			}
			seenTiles[key] = beam.dir

			// Check where the beam is
			switch input[beam.c.Y][beam.c.X] {
			case '.':
				// Move in the same dir
				beam.c = beam.c.Move(beam.dir)
				break
			case '\\':
				switch beam.dir {
				case helper.NORTH:
					beam.c = beam.c.Move(helper.WEST)
					beam.dir = helper.WEST
				case helper.SOUTH:
					beam.c = beam.c.Move(helper.EAST)
					beam.dir = helper.EAST
				case helper.EAST:
					beam.c = beam.c.Move(helper.SOUTH)
					beam.dir = helper.SOUTH
				case helper.WEST:
					beam.c = beam.c.Move(helper.NORTH)
					beam.dir = helper.NORTH
				}
			case '/':
				// Move north or south
				switch beam.dir {
				case helper.NORTH:
					beam.c = beam.c.Move(helper.EAST)
					beam.dir = helper.EAST
				case helper.SOUTH:
					beam.c = beam.c.Move(helper.WEST)
					beam.dir = helper.WEST
				case helper.EAST:
					beam.c = beam.c.Move(helper.NORTH)
					beam.dir = helper.NORTH
				case helper.WEST:
					beam.c = beam.c.Move(helper.SOUTH)
					beam.dir = helper.SOUTH
				}
			case '-':
				// split hozirontal, if travelling north or south
				if beam.dir == helper.NORTH || beam.dir == helper.SOUTH {
					newbeams = append(newbeams, &Beam{
						c:   beam.c.Move(helper.WEST),
						dir: helper.WEST,
					}, &Beam{
						c:   beam.c.Move(helper.EAST),
						dir: helper.EAST,
					})
					// remove existing beam
					beams = append(beams[:i], beams[i+1:]...)
				} else {
					// just keep on moving
					beam.c = beam.c.Move(beam.dir)
				}
			case '|':
				// split vertical
				if beam.dir == helper.EAST || beam.dir == helper.WEST {
					newbeams = append(newbeams, &Beam{
						c:   beam.c.Move(helper.NORTH),
						dir: helper.NORTH,
					}, &Beam{
						c:   beam.c.Move(helper.SOUTH),
						dir: helper.SOUTH,
					})
					// remove existing beam
					beams = append(beams[:i], beams[i+1:]...)
				} else {
					// just keep on moving
					beam.c = beam.c.Move(beam.dir)
				}
			}
			beams = append(beams, newbeams...)
		}

		// Remove beams that have left the grid
		for i := len(beams) - 1; i >= 0; i-- {
			beam := beams[i]
			if beam.c.X < 0 || beam.c.X >= len(input[0]) || beam.c.Y < 0 || beam.c.Y >= len(input) {
				// it's no longer valid, remove it
				beams = append(beams[:i], beams[i+1:]...)
			}
		}
	}

	return len(seenTiles)
}

func problem(input []string, partTwo bool) (int, error) {
	if !partTwo {
		b := Beam{
			c:   helper.Coord{X: 0, Y: 0},
			dir: helper.EAST,
		}
		return simulate(b, input), nil
	}

	max := math.MinInt
	// top edge
	ylen := len(input) - 1
	xlen := len(input[0]) - 1
	for y := 0; y <= ylen; y++ {
		for x := 0; x <= xlen; x++ {
			ans := 0
			ans2 := 0
			// if not an edge, ignore
			if x != 0 && y != 0 && x != xlen && y != ylen {
				continue
			}
			// top edge
			if y == 0 {
				ans = simulate(Beam{
					c:   helper.Coord{X: x, Y: y},
					dir: helper.SOUTH,
				}, input)
				// special corner cases
				if x == 0 {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.EAST,
					}, input)
				} else if x == xlen {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.WEST,
					}, input)
				}
			} else if x == 0 {
				// left edge
				ans = simulate(Beam{
					c:   helper.Coord{X: x, Y: y},
					dir: helper.EAST,
				}, input)
				// special corner cases
				if y == ylen {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.NORTH,
					}, input)
				}
			} else if x == xlen {
				// right edge
				ans = simulate(Beam{
					c:   helper.Coord{X: x, Y: y},
					dir: helper.WEST,
				}, input)
				// special corner cases
				if y == 0 {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.SOUTH,
					}, input)
				} else if y == ylen {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.NORTH,
					}, input)
				}
			} else if y == ylen {
				// bottom edge
				ans = simulate(Beam{
					c:   helper.Coord{X: x, Y: y},
					dir: helper.NORTH,
				}, input)
				// special corner cases
				if x == xlen {
					ans2 = simulate(Beam{
						c:   helper.Coord{X: x, Y: y},
						dir: helper.WEST,
					}, input)
				}
			}

			if ans > max {
				max = ans
			} else if ans2 > max {
				max = ans2
			}
		}
	}

	return max, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"slices"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type Room struct {
	c     *helper.Coord
	steps int
	ends  map[int]int
	raw   string
}

func problem(lines []string, partTwo bool) (int, error) {

	grid, start := readData(lines)
	// Let's assume we start heading from east
	moving := EAST

	// override start dir for test
	//start.ends = map[int]int{
	//	SOUTH: WEST,
	//	WEST:  SOUTH,
	//}
	//moving = WEST

	startRoom := start
	var curRoom *Room
	stepCount := 0
	for startRoom != curRoom {
		// first first step
		if curRoom == nil {
			curRoom = startRoom
		}
		curRoom.steps = stepCount

		nextMove := curRoom.NextStep(moving)

		switch nextMove {
		case NORTH:
			curRoom = grid[curRoom.c.Y-1][curRoom.c.X]
			// flag that we're entering from the south
			moving = SOUTH
		case EAST:
			curRoom = grid[curRoom.c.Y][curRoom.c.X+1]
			moving = WEST
		case SOUTH:
			curRoom = grid[curRoom.c.Y+1][curRoom.c.X]
			moving = NORTH
		case WEST:
			curRoom = grid[curRoom.c.Y][curRoom.c.X-1]
			moving = EAST
		}
		stepCount++
	}

	answer := 0
	if partTwo {
		for _, row := range grid {
			within := false
			for _, room := range row {
				if room.steps > 0 || start == room {
					if slices.Contains([]string{"J", "|", "L"}, room.raw) || start == room {
						within = !within
					}
				} else {
					if within {
						answer++
					}
				}
			}
		}
	} else {
		answer = stepCount / 2
	}
	return answer, nil
}

func (r *Room) NextStep(from int) int {
	return r.ends[from]
}

func readData(lines []string) ([][]*Room, *Room) {
	ret := make([][]*Room, 0)
	var start *Room

	for y, r := range lines {
		row := make([]*Room, len(r))
		for x, v := range r {
			room := Room{
				c: &helper.Coord{
					X: x,
					Y: y,
				},
				steps: 0,
				raw:   string(v),
			}
			switch v {
			case '|':
				room.ends = map[int]int{NORTH: SOUTH, SOUTH: NORTH}
			case '-':
				room.ends = map[int]int{EAST: WEST, WEST: EAST}
			case 'L':
				room.ends = map[int]int{NORTH: EAST, EAST: NORTH}
			case 'J':
				room.ends = map[int]int{NORTH: WEST, WEST: NORTH}
			case '7':
				room.ends = map[int]int{WEST: SOUTH, SOUTH: WEST}
			case 'F':
				room.ends = map[int]int{EAST: SOUTH, SOUTH: EAST}
			case 'S':
				room.ends = map[int]int{NORTH: EAST, EAST: NORTH}
				start = &room
			}
			row[x] = &room
		}
		ret = append(ret, row)
	}

	return ret, start
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

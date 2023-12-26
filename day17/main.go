package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"slices"
	"strconv"
)

const (
	P1_MAX_MOVE = 3
	P2_MAX_MOVE = 10
	P2_MIN_MOVE = 4
)

type Walk struct {
	c         helper.Coord
	dir       int
	samesteps int
	total     int
	walked    []helper.Coord
}

func problem(input []string, partTwo bool) (int, error) {
	var ans int
	if partTwo {
		ans = walk(input, P2_MAX_MOVE, true)
	} else {
		ans = walk(input, P1_MAX_MOVE, false)
	}
	return ans, nil
}

func walk(grid []string, maxMove int, partTwo bool) int {

	finish := helper.Coord{X: len(grid[0]) - 1, Y: len(grid) - 1}
	min := math.MaxInt
	seen := make(map[string]bool)

	queue := []Walk{
		{
			c:         helper.Coord{X: 0, Y: 0},
			dir:       helper.NORTH,
			samesteps: 4, // force a dir change for part 2
			total:     0,
			walked:    make([]helper.Coord, 0),
		},
	}

	// We don't count the starting heat
	startheat, _ := strconv.Atoi(string(grid[0][0]))

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.total > min {
			// abandon ship!
			continue
		}

		key := fmt.Sprintf("%v:%v", next.c.ToString(), next.total)
		if _, ok := seen[key]; ok {
			continue
		} else {
			seen[key] = true
		}

		if next.c == finish {
			if partTwo {
				if next.samesteps != P2_MIN_MOVE {
					// need to have walked 4 steps when reaching the finish
					continue
				}
			}
			if next.total < min {
				min = next.total
				continue
			}
		}

		nextval, _ := strconv.Atoi(string(grid[next.c.Y][next.c.X]))
		for d := range []int{helper.NORTH, helper.SOUTH, helper.EAST, helper.WEST} {
			if next.dir == d && next.samesteps >= maxMove {
				// we cannot take this movement
				continue
			}
			nc := next.c.MoveDirection(d)
			if nc.X < 0 || nc.X >= len(grid[0]) || nc.Y < 0 || nc.Y >= len(grid) {
				// bad coordinate
				continue
			}
			if slices.Contains(next.walked, nc) {
				// don't walk somewhere we've walked
				continue
			}

			if partTwo {
				if next.samesteps < P2_MIN_MOVE && d != next.dir {
					// we need to keep moving in the same direction
					continue
				}
			}

			// We can only turn 90 degrees
			if (next.dir == helper.NORTH && d == helper.SOUTH) ||
				(next.dir == helper.SOUTH && d == helper.NORTH) ||
				(next.dir == helper.EAST && d == helper.WEST) ||
				(next.dir == helper.WEST && d == helper.EAST) {
				continue
			}

			nextwalked := make([]helper.Coord, 0)
			for _, v := range next.walked {
				nextwalked = append(nextwalked, v)
			}
			nextwalked = append(nextwalked, nc)

			nextsteps := 1
			if next.dir == d {
				nextsteps = next.samesteps + 1
			}

			nextstep := Walk{
				c:         nc,
				dir:       d,
				samesteps: nextsteps,
				total:     next.total + nextval,
				walked:    nextwalked,
			}
			queue = append(queue, nextstep)
		}
	}

	return min - startheat
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Part one: %v\n", ans)
	}

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

func partone(input []string, maxsteps int) (int, error) {
	grid, start := readData(input)
	seen := make(map[helper.Coord]bool)

	steps := 0
	seen[start] = true
	for steps < maxsteps {
		additions := make([]helper.Coord, 0)
		for coord, _ := range seen {
			neighbours := coord.GetNeighbours(false)
			for _, neighbour := range neighbours {
				if _, ok := seen[neighbour]; ok {
					continue
				}
				if neighbour.X < 0 || neighbour.X >= len(input[0]) || neighbour.Y < 0 || neighbour.Y >= len(input) {
					continue
				}
				if grid[neighbour.Y][neighbour.X] == true {
					// rock or start
					continue
				}
				additions = append(additions, neighbour)
			}
		}
		seen = make(map[helper.Coord]bool)
		for _, c := range additions {
			seen[c] = true
		}
		steps++
	}
	return len(seen), nil
}

func parttwo(input []string, maxsteps int) (int, error) {
	grid, start := readData(input)
	seen := make(map[helper.Coord]bool)

	ylen := len(input)
	xlen := len(input[0])

	steps := 0
	seen[start] = true

	p := make([]int, 0)

	for steps < maxsteps {
		additions := make([]helper.Coord, 0)
		for coord, _ := range seen {
			neighbours := coord.GetNeighbours(false)
			for _, neighbour := range neighbours {
				if _, ok := seen[neighbour]; ok {
					continue
				}
				if grid[((neighbour.Y%ylen)+ylen)%ylen][((neighbour.X%xlen)+xlen)%xlen] == true {
					// rock or start
					continue
				}
				additions = append(additions, neighbour)
			}
		}
		seen = make(map[helper.Coord]bool)
		for _, c := range additions {
			seen[c] = true
		}
		steps++

		if steps%ylen == maxsteps%ylen {
			p = append(p, len(seen))
			if len(p) == 3 {
				p0 := p[0]
				p1 := p[1] - p[0]
				p2 := p[2] - p[1]
				return p0 + (p1 * (maxsteps / ylen)) + ((maxsteps/ylen)*((maxsteps/ylen)-1)/2)*(p2-p1), nil
			}
		}
	}
	return 0, nil
}

func readData(input []string) ([]map[int]bool, helper.Coord) {
	r := make([]map[int]bool, 0)
	var start helper.Coord
	for y, line := range input {
		row := make(map[int]bool)
		for x, v := range line {
			row[x] = v == '#'
			if v == 'S' {
				start = helper.Coord{X: x, Y: y}
			}
		}
		r = append(r, row)
	}
	return r, start
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines, 6)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines, 26501365)
	fmt.Printf("Part two: %v\n", ans)

}

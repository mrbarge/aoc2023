package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

func problem(input []string, partTwo bool) (int, error) {
	grid, start := readData(input)
	seen := make(map[helper.Coord]bool)

	steps := 0
	seen[start] = true
	for steps < 64 {
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

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	//ans, err = problem(lines, true)
	//fmt.Printf("Part two: %v\n", ans)

}

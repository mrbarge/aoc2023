package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"slices"
	"strings"
)

func problem(lines []string, partTwo bool) (int, error) {

	galaxies := expand(lines, partTwo)
	up := findUniquePairs(galaxies)

	ans := 0
	for _, iv := range up {
		sp := helper.ManhattanDistance(iv[0], iv[1])
		ans += sp
	}

	return ans, nil
}

func findUniquePairs(array []helper.Coord) [][]helper.Coord {
	var pairs [][]helper.Coord
	existingPairs := make(map[string]bool)

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			pair := []helper.Coord{array[i], array[j]}
			pairString := fmt.Sprintf("%v %v", array[i], array[j])

			_, exists := existingPairs[pairString]
			if !exists {
				pairs = append(pairs, pair)
				existingPairs[pairString] = true
			}
		}
	}

	return pairs
}

func expand(data []string, partTwo bool) []helper.Coord {
	// find rows with no galaxies
	rows := make([]int, 0)
	cols := make([]int, 0)
	galaxyCoords := make([]helper.Coord, 0)
	for y, r := range data {
		if !strings.Contains(r, "#") {
			rows = append(rows, y)
		}
	}

	for x := 0; x < len(data[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(data); y++ {
			if data[y][x] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			cols = append(cols, x)
		}
	}

	ny := 0
	for y := 0; y < len(data); y++ {
		duplicateRow := slices.Contains(rows, y)
		nx := 0
		for x := 0; x < len(data[0]); x++ {
			duplicateCol := slices.Contains(cols, x)
			galaxy := data[y][x] == '#'
			if duplicateCol {
				if partTwo {
					nx += 999999
				} else {
					nx++
				}
			}
			if galaxy {
				galaxyCoords = append(galaxyCoords, helper.Coord{X: x + nx, Y: y + ny})
			}
		}
		if duplicateRow {
			if partTwo {
				ny += 999999
			} else {
				ny++
			}
		}
	}

	return galaxyCoords
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

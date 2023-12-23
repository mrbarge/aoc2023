package main

import (
	"fmt"
	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

func problem(input []string, partTwo bool) (int, error) {
	var g dijkstra.Graph
	if partTwo {
		g = buildGraphP2(input)
	} else {
		g = buildGraph(input)
	}

	//from := "1,0"
	//to := "21,22"
	from := "1,0"
	to := "139,140"

	path, _, _ := g.Path(from, to)
	return len(path) - 1, nil
}

func buildGraphP2(input []string) dijkstra.Graph {
	g := dijkstra.Graph{}

	for y, row := range input {
		for x, v := range row {
			if v == '#' {
				continue
			}

			key := fmt.Sprintf("%v,%v", x, y)
			g[key] = make(map[string]int)
			c := helper.Coord{X: x, Y: y}

			n := c.GetSafeNeighbours(false, len(row), len(input))
			for _, neighbour := range n {
				if input[neighbour.Y][neighbour.X] == '#' {
					continue
				} else {
					nk := fmt.Sprintf("%v,%v", neighbour.X, neighbour.Y)
					g[key][nk] = -1
				}
			}
		}
	}
	return g
}

func buildGraph(input []string) dijkstra.Graph {
	g := dijkstra.Graph{}

	for y, row := range input {
		for x, v := range row {
			if v == '#' {
				continue
			}

			key := fmt.Sprintf("%v,%v", x, y)
			g[key] = make(map[string]int)
			c := helper.Coord{X: x, Y: y}

			if v == 'v' {
				nk := fmt.Sprintf("%v,%v", c.X, c.Y+1)
				g[key][nk] = -1
			} else if v == '>' {
				nk := fmt.Sprintf("%v,%v", c.X+1, c.Y)
				g[key][nk] = -1
			} else {
				n := c.GetSafeNeighbours(false, len(row), len(input))
				for _, neighbour := range n {
					if input[neighbour.Y][neighbour.X] == '#' {
						continue
					} else if input[neighbour.Y][neighbour.X] == '.' {
						nk := fmt.Sprintf("%v,%v", neighbour.X, neighbour.Y)
						g[key][nk] = -1
					} else if input[neighbour.Y][neighbour.X] == 'v' {
						// can't enter here from below
						if c.Y > neighbour.Y {
							continue
						}
						nk := fmt.Sprintf("%v,%v", neighbour.X, neighbour.Y)
						g[key][nk] = -1
					} else if input[neighbour.Y][neighbour.X] == '>' {
						// can't enter here from the left
						if c.X > neighbour.X {
							continue
						}
						nk := fmt.Sprintf("%v,%v", neighbour.X, neighbour.Y)
						g[key][nk] = -1
					}
				}
			}
		}
	}
	return g
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, true)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Part one: %v\n", ans)

	//ans, err = problem(lines, true)
	//fmt.Printf("Part two: %v\n", ans)

}

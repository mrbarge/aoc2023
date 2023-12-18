package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func problem(input []string, partTwo bool) (int, error) {

	coords := make([]helper.Coord, 0)
	c := helper.Coord{X: 1, Y: 1}

	coords = append(coords, c)
	perimeter := 0

	for _, instruction := range input {
		dir := strings.Fields(instruction)[0]
		path, _ := strconv.Atoi(strings.Fields(instruction)[1])
		cv := strings.Fields(instruction)[2]
		colour := cv[1 : len(cv)-1]

		if partTwo {
			nd := new(big.Int)
			nd.SetString(colour[1:6], 16)
			path = int(nd.Int64())
			switch colour[6] {
			case '0':
				dir = "R"
			case '1':
				dir = "D"
			case '2':
				dir = "L"
			case '3':
				dir = "U"
			}
		}

		switch dir {
		case "R":
			c.X += path
		case "D":
			c.Y += path
		case "U":
			c.Y -= path
		case "L":
			c.X -= path
		}
		perimeter += path
		coords = append(coords, c)
	}

	ans := area(coords)
	return int(ans) + int(perimeter/2) + 1, nil
}

func area(coords []helper.Coord) float64 {
	sum := 0.0
	p0 := coords[len(coords)-1]
	for _, p1 := range coords {
		sum += float64(p0.Y*p1.X - p0.X*p1.Y)
		p0 = p1
	}
	return math.Abs(sum) / 2
}

func main() {
	fh, _ := os.Open("input.txt")
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

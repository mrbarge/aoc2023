package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

//const RANGE_LOW float64 = 7
//const RANGE_HIGH float64 = 27

const RANGE_LOW = 200000000000000
const RANGE_HIGH = 400000000000000

type Hails []Hail
type Hail struct {
	c helper.Coord3D // coordinate
	v helper.Coord3D // velocity
}

func canIntersect(a Hail, b Hail) bool {
	a1 := float64(a.v.Y) / float64(a.v.X)
	b1 := float64(a.c.Y) - a1*float64(a.c.X)
	a2 := float64(b.v.Y) / float64(b.v.X)
	b2 := float64(b.c.Y) - a2*float64(b.c.X)

	if almostEqual(a1, a2) {
		if almostEqual(b1, b2) {
			return true
		}
		return false
	}

	cx := (b2 - b1) / (a1 - a2)
	cy := cx*a1 + b1
	f := (cx > float64(a.c.X)) == (a.v.X > 0) && (cx > float64(b.c.X)) == (b.v.X > 0)

	if f && RANGE_LOW <= cx && cx <= RANGE_HIGH && RANGE_LOW <= cy && cy <= RANGE_HIGH {
		return true
	}
	return false
}

func problem(input []string, partTwo bool) (int, error) {
	hailstones := readData(input)
	ans := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if canIntersect(hailstones[i], hailstones[j]) {
				ans++
			}
		}
	}
	return ans, nil
}

// From https://stackoverflow.com/questions/47969385/go-float-comparison
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func readData(input []string) []Hail {
	r := make([]Hail, 0)
	for _, line := range input {
		sc := strings.Split(line, " @ ")[0]
		sv := strings.Split(line, " @ ")[1]
		sce := strings.Split(sc, ", ")
		sve := strings.Split(sv, ", ")
		cx, _ := strconv.Atoi(strings.TrimSpace(sce[0]))
		cy, _ := strconv.Atoi(strings.TrimSpace(sce[1]))
		cz, _ := strconv.Atoi(strings.TrimSpace(sce[2]))
		vx, _ := strconv.Atoi(strings.TrimSpace(sve[0]))
		vy, _ := strconv.Atoi(strings.TrimSpace(sve[1]))
		vz, _ := strconv.Atoi(strings.TrimSpace(sve[2]))
		c := Hail{
			c: helper.Coord3D{X: cx, Y: cy, Z: cz},
			v: helper.Coord3D{X: vx, Y: vy, Z: vz},
		}
		r = append(r, c)
	}
	return r
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

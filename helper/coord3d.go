package helper

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Coord3D struct {
	X int
	Y int
	Z int
	V int
}

func (c Coord3D) AsString() string {
	return fmt.Sprintf("%d/%d/%d", c.X, c.Y, c.Z)
}

func (c Coord3D) EmptySides(s []Coord3D) int {
	m := make(map[string]bool)
	n := c.AdjNeighbours()

	count := 0
	m[c.AsString()] = true
	for _, v := range s {
		m[v.AsString()] = true
	}

	for _, v := range n {
		if _, seen := m[v.AsString()]; !seen {
			count++
		}
	}

	return count
}

func (c Coord3D) AllEnclosed(s []Coord3D) bool {
	m := make(map[string]bool)
	n := c.AdjNeighbours()

	m[c.AsString()] = true
	for _, v := range s {
		m[v.AsString()] = true
	}

	all := true
	for _, v := range n {
		if _, seen := m[v.AsString()]; !seen {
			all = false
			break
		}
	}

	return all
}

func xRange(s []Coord3D) []int {
	r := make([]int, 0)
	for _, v := range s {
		r = append(r, v.X)
	}
	return r
}

func yRange(s []Coord3D) []int {
	r := make([]int, 0)
	for _, v := range s {
		r = append(r, v.Y)
	}
	return r
}

func zRange(s []Coord3D) []int {
	r := make([]int, 0)
	for _, v := range s {
		r = append(r, v.Z)
	}
	return r
}

func Ranges(s []Coord3D) (xmin int, xmax int, ymin int, ymax int, zmin int, zmax int) {
	xr := xRange(s)
	yr := yRange(s)
	zr := zRange(s)
	sort.Ints(xr)
	sort.Ints(yr)
	sort.Ints(zr)
	xmin = xr[0]
	xmax = xr[len(xr)-1]
	ymin = yr[0]
	ymax = yr[len(yr)-1]
	zmin = zr[0]
	zmax = zr[len(zr)-1]
	return
}

func (c Coord3D) AllNeighbours() []Coord3D {
	r := make([]Coord3D, 0)

	for i := range []int{-1, 0, 1} {
		for j := range []int{-1, 0, 1} {
			for k := range []int{-1, 0, 1} {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				r = append(r, Coord3D{X: c.X + i, Y: c.Y + j, Z: c.Z + k})
			}
		}
	}
	return r
}

func (c Coord3D) AdjNeighbours() []Coord3D {
	return []Coord3D{
		{X: c.X + 1, Y: c.Y, Z: c.Z},
		{X: c.X - 1, Y: c.Y, Z: c.Z},
		{X: c.X, Y: c.Y + 1, Z: c.Z},
		{X: c.X, Y: c.Y - 1, Z: c.Z},
		{X: c.X, Y: c.Y, Z: c.Z + 1},
		{X: c.X, Y: c.Y, Z: c.Z - 1},
	}
}

func ReadCoord3D(s string) Coord3D {
	r := Coord3D{}
	elems := strings.Split(s, ",")
	r.X, _ = strconv.Atoi(elems[0])
	r.Y, _ = strconv.Atoi(elems[1])
	r.Z, _ = strconv.Atoi(elems[2])
	return r
}

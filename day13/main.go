package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

func problem(grids [][]string, partTwo bool) (int, error) {
	ans := 0
	for _, grid := range grids {

		res, isHoriz := mirror(grid)
		if !partTwo {
			if isHoriz {
				fmt.Printf("Horizontal mirror found at row %v\n", res)
				ans += res * 100
			} else {
				fmt.Printf("Vertical mirror found at col %v\n", res)
				ans += res
			}
		} else {
			found := false
			for y, row := range grid {
				for x, _ := range row {
					newgrid := make([]string, len(grid))
					copy(newgrid, grid)
					if newgrid[y][x] == '.' {
						newgrid[y] = newgrid[y][:x] + "#" + newgrid[y][x+1:]
					} else {
						newgrid[y] = newgrid[y][:x] + "." + newgrid[y][x+1:]
					}

					res2, isHoriz2 := mirror(newgrid)
					if res2 == -1 {
						// ignore
						continue
					}
					if res == res2 && isHoriz == isHoriz2 {
						// ignore
						continue
					}

					if isHoriz2 {
						fmt.Printf("Changed %v,%v and horizontal mirror found at row %v\n", x, y, res2)
						ans += res2 * 100
						found = true
						break
					} else {
						fmt.Printf("Changed %v,%v and vertical mirror found at col %v\n", x, y, res2)
						ans += res2
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}

	}
	return ans, nil
}

func mirror(grid []string) (int, bool) {
	l := len(grid) - 1
	// horizontal
	for i := 0; i < l; i++ {
		line := grid[i]
		if line != grid[l] {
			continue
		}

		// possible mirror, validate that
		isMirror := true
		ix := i + 1
		ij := l - 1
		for ix != ij && ix < ij {
			if grid[ix] != grid[ij] {
				isMirror = false
				break
			}
			ix++
			ij--
		}

		if isMirror {
			// found horizontal mirror
			return ij + 1, true
		}
	}

	for i := l; i > 0; i-- {
		line := grid[i]
		if line != grid[0] {
			continue
		}

		// possible mirror, validate that
		isMirror := true
		ix := i - 1
		ij := 1
		for ix != ij && ix > ij {
			if grid[ix] != grid[ij] {
				isMirror = false
				break
			}
			ix--
			ij++
		}

		if isMirror {
			// found horizontal mirror
			return ij, true
		}
	}

	// vertical
	vl := len(grid[0]) - 1
	for i := 0; i < vl; i++ {
		line := getCol(grid, i)
		if line != getCol(grid, vl) {
			continue
		}

		// possible mirror, validate that
		isMirror := true
		ix := i + 1
		ij := vl - 1
		for ix != ij && ix < ij {
			ixl := getCol(grid, ix)
			ijl := getCol(grid, ij)
			if ixl != ijl {
				isMirror = false
				break
			}
			ix++
			ij--
		}

		if isMirror {
			// found vertical mirror
			return ij + 1, false
		}
	}

	// vertical
	for i := vl; i > 0; i-- {
		line := getCol(grid, i)
		if line != getCol(grid, 0) {
			continue
		}

		// possible mirror, validate that
		isMirror := true
		ix := i - 1
		ij := 1
		for ix != ij && ix > ij {
			ixl := getCol(grid, ix)
			ijl := getCol(grid, ij)
			if ixl != ijl {
				isMirror = false
				break
			}
			ix--
			ij++
		}

		if isMirror {
			// found vertical mirror
			return ij, false
		}
	}
	return -1, false
}

func getHorizontalFreq(l []string) (r map[string][]int) {
	r = make(map[string][]int)
	for row, line := range l {
		if _, ok := r[line]; !ok {
			r[line] = []int{row}
		} else {
			r[line] = append(r[line], row)
		}
	}
	return r
}

func getVerticalFreq(l []string) (r map[string][]int) {
	r = make(map[string][]int)
	for i := 0; i < len(l[0]); i++ {
		c := getCol(l, i)
		if _, ok := r[c]; !ok {
			r[c] = []int{i}
		} else {
			r[c] = append(r[c], i)
		}
	}
	return r
}

func getCol(l []string, col int) (r string) {
	for _, v := range l {
		r += string(v[col])
	}
	return r
}

func splitInput(l []string) [][]string {
	r := make([][]string, 0)

	tmp := make([]string, 0)
	for _, line := range l {
		if line == "" {
			r = append(r, tmp)
			tmp = make([]string, 0)
		} else {
			tmp = append(tmp, line)
		}
	}
	r = append(r, tmp)
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	grids := splitInput(lines)
	ans, err := problem(grids, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(grids, true)
	fmt.Printf("Part two: %v\n", ans)

}

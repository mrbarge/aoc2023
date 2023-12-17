package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

var seenIt map[string]int

func problem(grid []string, partTwo bool) (int, error) {
	seenIt = make(map[string]int)
	ans := 0
	if partTwo {
		converted := convert(grid)
		for i := 1; i <= 1000000000; i++ {
			converted = cycle(converted, 0, i)
			converted = cycle(converted, 3, i)
			converted = cycle(converted, 2, i)
			converted = cycle(converted, 1, i)
			key := makeit(converted, 1)
			if _, ok := seenIt[key]; ok {
				if (1000000000-i)%(seenIt[key]-i) == 0 {
					break
				}
			} else {
				seenIt[key] = i
			}
		}
		ans = calcLoad(converted)
	} else {
		converted := convert(grid)
		converted = cycle(converted, 0, 0)
		ans = calcLoad(converted)
	}

	return ans, nil
}

func makeit(grid [][]string, dir int) string {
	r := ""
	for _, row := range grid {
		for _, v := range row {
			r += v
		}
	}
	return r
}

func calcLoad(grid [][]string) int {
	ans := 0
	for y, row := range grid {
		for _, v := range row {
			if v == "O" {
				ans += len(grid) - y
			}
		}
	}
	return ans
}

func convert(grid []string) [][]string {
	r := make([][]string, 0)
	for _, row := range grid {
		r2 := make([]string, 0)
		for _, v := range row {
			r2 = append(r2, string(v))
		}
		r = append(r, r2)
	}
	return r
}

func printit(r [][]string) {
	for _, row := range r {
		for _, v := range row {
			fmt.Printf("%v", v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func cycle(r [][]string, dir int, cycle int) [][]string {
	if dir == 0 {
		// north
		for col := 0; col < len(r[0]); col++ {
			for row := 0; row < len(r); row++ {
				if r[row][col] == "O" {
					tmp := row
					for row2 := tmp - 1; tmp > 0; row2-- {
						if r[row2][col] == "." {
							// can move up
							r[row2][col] = "O"
							r[tmp][col] = "."
						} else if r[row2][col] == "#" || r[row2][col] == "O" {
							// can't move up
							break
						}
						tmp--
					}
				}
			}
		}
	} else if dir == 1 {
		// east
		for row := 0; row < len(r); row++ {
			for col := len(r[0]) - 1; col >= 0; col-- {
				if r[row][col] == "O" {
					tmp := col
					for col2 := tmp + 1; tmp < len(r[0])-1; col2++ {
						if r[row][col2] == "." {
							// can move right
							r[row][col2] = "O"
							r[row][tmp] = "."
						} else if r[row][col2] == "#" || r[row][col2] == "O" {
							// can't move right
							break
						}
						tmp++
					}
				}
			}
		}
	} else if dir == 2 {
		// south
		for col := 0; col < len(r[0]); col++ {
			for row := len(r) - 1; row >= 0; row-- {
				if r[row][col] == "O" {
					tmp := row
					for row2 := tmp + 1; tmp < len(r)-1; row2++ {
						if r[row2][col] == "." {
							// can move down
							r[row2][col] = "O"
							r[tmp][col] = "."
						} else if r[row2][col] == "#" || r[row2][col] == "O" {
							// can't move up
							break
						}
						tmp++
					}
				}
			}
		}
	} else if dir == 3 {
		// west
		for row := 0; row < len(r); row++ {
			for col := 0; col < len(r[0]); col++ {
				if r[row][col] == "O" {
					tmp := col
					for col2 := tmp - 1; tmp > 0; col2-- {
						if r[row][col2] == "." {
							// can move right
							r[row][col2] = "O"
							r[row][tmp] = "."
						} else if r[row][col2] == "#" || r[row][col2] == "O" {
							// can't move right
							break
						}
						tmp--
					}
				}
			}
		}
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

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

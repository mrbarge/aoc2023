package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
)

func problem(lines []string, partTwo bool) (int, error) {

	numbers := make([]int, 0)
	for y, l := range lines {
		buildnum := ""
		validnum := false
		for x, c := range l {
			coord := helper.Coord{X: x, Y: y}
			if c < '0' || c > '9' {
				if buildnum == "" {
					continue
				}
				if validnum {
					n, _ := strconv.Atoi(buildnum)
					numbers = append(numbers, n)
					validnum = false
				}
				buildnum = ""
			} else {
				// see if this number is valid
				neighbours := coord.GetNeighboursPos(true)
				for _, neighbour := range neighbours {
					if neighbour.X < len(l) && neighbour.Y < len(lines) && isSymbol(lines[neighbour.Y][neighbour.X]) {
						validnum = true
					}
				}
				buildnum += string(c)
			}
		}
		if buildnum != "" && validnum {
			n, _ := strconv.Atoi(buildnum)
			numbers = append(numbers, n)
			validnum = false
			buildnum = ""
		}
	}

	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum, nil
}

func partTwo(lines []string) (int, error) {

	numbers := make([]int, 0)
	partNum := make([][]int, 0)
	for _, l := range lines {
		partNum = append(partNum, make([]int, len(l)))
	}
	curPartNum := 1
	parts := make(map[int]int)

	for y, l := range lines {
		buildnum := ""
		validnum := false
		for x, c := range l {
			coord := helper.Coord{X: x, Y: y}
			if c < '0' || c > '9' {
				if buildnum == "" {
					partNum[y][x] = 0
					continue
				}
				if validnum {
					n, _ := strconv.Atoi(buildnum)
					numbers = append(numbers, n)
					validnum = false
					parts[curPartNum] = n
					for tx := x - len(buildnum); tx < x; tx++ {
						partNum[y][tx] = n
					}
					curPartNum++
				}
				buildnum = ""
			} else {
				// see if this number is valid
				neighbours := coord.GetNeighboursPos(true)
				for _, neighbour := range neighbours {
					if neighbour.X < len(l) && neighbour.Y < len(lines) && lines[neighbour.Y][neighbour.X] == '*' {
						validnum = true
					}
				}
				buildnum += string(c)
			}
		}
		if buildnum != "" && validnum {
			n, _ := strconv.Atoi(buildnum)
			numbers = append(numbers, n)
			parts[curPartNum] = n
			for tx := len(l) - len(buildnum); tx < len(l); tx++ {
				partNum[y][tx] = n
			}
			curPartNum++
			validnum = false
			buildnum = ""
		}
	}

	sum := 0
	for y, l := range lines {
		for x, c := range l {
			if c == '*' {
				seenParts := make(map[int]int)
				coord := helper.Coord{X: x, Y: y}
				neighbours := coord.GetNeighboursPos(true)
				for _, neighbour := range neighbours {
					pn := partNum[neighbour.Y][neighbour.X]
					if pn > 0 {
						seenParts[pn] = pn
					}
				}
				if len(seenParts) == 2 {
					ns := make([]int, 0)
					for _, v := range seenParts {
						ns = append(ns, v)
					}
					ratio := 1
					for _, v := range seenParts {
						ratio *= v
					}
					sum += ratio
				}
			}
		}
	}

	return sum, nil
}

func isSymbol(r uint8) bool {
	return r != '.' && !(r >= '0' && r <= '9')
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

	ans2, err := partTwo(lines)
	fmt.Printf("Part two: %v\n", ans2)
}

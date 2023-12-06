package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
	"strings"
)

func partOne(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		s1 := strings.Index(line, ":")
		s2 := strings.Index(line, "|")
		winners := strings.Fields(line[s1+1 : s2-1])
		entries := strings.Fields(line[s2+2:])

		winnerNums := make(map[int]bool)
		for _, v := range winners {
			vn, _ := strconv.Atoi(v)
			winnerNums[vn] = true
		}

		score := 0
		for _, v := range entries {
			vn, _ := strconv.Atoi(v)
			if _, ok := winnerNums[vn]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		sum += score
	}
	return sum, nil
}

func partTwo(lines []string) (int, error) {
	sum := 0
	cardQueue := make(map[int]int)
	cardWinners := make(map[int]int)

	for i, line := range lines {
		cardQueue[i+1] = 1

		s1 := strings.Index(line, ":")
		s2 := strings.Index(line, "|")
		winners := strings.Fields(line[s1+1 : s2-1])
		entries := strings.Fields(line[s2+2:])

		winnerNums := make(map[int]bool)
		for _, v := range winners {
			vn, _ := strconv.Atoi(v)
			winnerNums[vn] = true
		}

		totalWinners := 0
		for _, v := range entries {
			vn, _ := strconv.Atoi(v)
			if _, ok := winnerNums[vn]; ok {
				totalWinners++
			}
		}
		cardWinners[i+1] = totalWinners
	}

	cardIdx := 1
	for cardIdx <= len(lines) {
		tw := cardWinners[cardIdx]
		for ti := cardIdx + 1; ti < cardIdx+1+tw; ti++ {
			cardQueue[ti] += 1
		}
		cardQueue[cardIdx] -= 1
		if cardQueue[cardIdx] == 0 {
			cardIdx++
		}
		sum++
	}
	return sum, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans2, err := partTwo(lines)
	fmt.Printf("Part two: %v\n", ans2)

}

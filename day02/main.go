package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func problem(lines []string, partTwo bool) (int, error) {

	gamelineRE := regexp.MustCompile(`^Game (.+): (.+)$`)
	result := 0
	for _, line := range lines {
		ballmax := make(map[string]int)
		gamematch := gamelineRE.FindStringSubmatch(line)
		if len(gamematch) > 0 {
			gameNum, _ := strconv.Atoi(gamematch[1])
			sets := strings.Split(gamematch[2], "; ")
			badSets := false
			for _, set := range sets {
				fixedset := strings.Replace(set, ",", "", -1)
				balls := strings.Split(fixedset, " ")
				for i := 0; i < len(balls); i++ {
					bv, _ := strconv.Atoi(balls[i])
					i++
					bc := balls[i]
					if _, ok := ballmax[bc]; ok {
						if bv > ballmax[bc] {
							ballmax[bc] = bv
						}
					} else {
						ballmax[bc] = bv
					}
					if (bc == "red" && bv > 12) || (bc == "blue" && bv > 14) || (bc == "green" && bv > 13) {
						// It worked perfectly
						badSets = true
					}
				}
				if badSets && !partTwo {
					break
				}
			}
			if !badSets && !partTwo {
				result += gameNum
				fmt.Printf("It's a good line for %v and %v, and %v\n", gameNum, line)
				continue
			}
			if partTwo {
				result += (1 * ballmax["green"] * ballmax["red"] * ballmax["blue"])
			}
		}
	}
	return result, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, true)
	fmt.Printf("Part one: %v\n", ans)
}

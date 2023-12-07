package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
	"strings"
)

func problem(lines []string, partTwo bool) (int, error) {
	races := make(map[int]int)
	racetimes := strings.Fields(lines[0][strings.Index(lines[0], ":")+1:])
	records := strings.Fields(lines[1][strings.Index(lines[1], ":")+1:])
	for i, v := range racetimes {
		iv, _ := strconv.Atoi(v)
		rv, _ := strconv.Atoi(records[i])
		races[iv] = rv
	}

	sum := 1
	for time, distance := range races {
		w := runRace(time, distance)
		fmt.Printf("For %v/%v, winners = %v\n", time, distance, w)
		sum *= w
	}
	return sum, nil
}

func runRace(time int, record int) int {
	winners := 0

	for i := 1; i < time; i++ {
		speed := i
		remainingTime := time - i
		distance := speed * remainingTime
		if distance > record {
			winners++
		}
	}
	return winners
}

func main() {
	fh, _ := os.Open("inputp2.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, true)
	fmt.Printf("Part one: %v\n", ans)

}

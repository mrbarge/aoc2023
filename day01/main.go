package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2023/helper"
)

func problem(lines []string, partTwo bool) (int, error) {

	valmap := make(map[string]int)
	for i := 1; i < 10; i++ {
		valmap[strconv.Itoa(i)] = i
	}

	if partTwo {
		valmap["one"] = 1
		valmap["two"] = 2
		valmap["three"] = 3
		valmap["four"] = 4
		valmap["five"] = 5
		valmap["six"] = 6
		valmap["seven"] = 7
		valmap["eight"] = 8
		valmap["nine"] = 9
	}

	sum := 0
	for _, line := range lines {
		f, l := findRange(valmap, line)
		n, err := strconv.Atoi(fmt.Sprintf("%d%d", f, l))
		if err != nil {
			return 0, err
		}
		sum += n
	}

	return sum, nil
}

func findRange(valmap map[string]int, input string) (int, int) {
	var firstKey, lastKey string
	first := math.MaxInt
	last := math.MinInt
	for k, _ := range valmap {
		f := strings.Index(input, k)
		if f >= 0 && f < first {
			first = f
			firstKey = k
		}
		l := strings.LastIndex(input, k)
		if l >= 0 && l > last {
			last = l
			lastKey = k
		}
	}
	return valmap[firstKey], valmap[lastKey]
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

package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
)

func problem(lines []string, partTwo bool) (int, error) {

	instructions := lines[0]

	nodemap := make(map[string]map[string]string)
	for _, line := range lines[2:] {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]

		if _, ok := nodemap[node]; !ok {
			nodemap[node] = make(map[string]string)
		}
		nodemap[node]["L"] = left
		nodemap[node]["R"] = right
	}

	curNode := "AAA"
	steps := 0
	iidx := 0
	for true {
		if curNode == "ZZZ" {
			break
		}
		if iidx == len(instructions) {
			iidx = 0
		}
		v := string(instructions[iidx])
		curNode = nodemap[curNode][string(v)]
		steps++
		iidx++
	}
	return steps, nil
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

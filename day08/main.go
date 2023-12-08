package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strings"
)

func partOne(lines []string) (int, error) {

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
	for {
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

func partTwo(lines []string) (int, error) {
	instructions := lines[0]

	curNodes := make([]string, 0)

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

		if strings.HasSuffix(node, "A") {
			curNodes = append(curNodes, node)
		}
	}

	// calculate time to Z for each starting node
	nodeSteps := make([]int, 0)
	for _, startingNode := range curNodes {
		curNode := startingNode
		steps := 0
		iidx := 0
		for {
			if strings.HasSuffix(curNode, "Z") {
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
		nodeSteps = append(nodeSteps, steps)
	}

	gcd := LCM(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...)

	return gcd, nil
}

// from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	//ans, err := partOne(lines)
	//fmt.Printf("Part one: %v\n", ans)
	//
	ans, err := partTwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}

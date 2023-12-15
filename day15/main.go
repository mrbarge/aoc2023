package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
	"strings"
)

type ValueMap struct {
	label string
	value int
}

func partOne(input string) (int, error) {
	steps := strings.Split(input, ",")

	ans := 0
	for _, step := range steps {
		ans += hash(step)
	}

	return ans, nil
}

func printBoxes(b [][]ValueMap) {
	for i, v := range b {
		if len(v) > 0 {
			fmt.Printf("Box %v: %v\n", i, v)
		}
	}
}

func partTwo(input string) (int, error) {
	steps := strings.Split(input, ",")
	boxes := make([][]ValueMap, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = make([]ValueMap, 0)
	}

	for _, step := range steps {
		if strings.Contains(step, "=") {
			label := strings.Split(step, "=")[0]
			box := boxes[hash(label)]
			value, _ := strconv.Atoi(strings.Split(step, "=")[1])

			found := false
			for n, v := range box {
				if v.label == label {
					box[n].value = value
					boxes[hash(label)] = box
					found = true
					break
				}
			}

			if !found {
				box = append(box, ValueMap{label: label, value: value})
				boxes[hash(label)] = box
			}
		} else {
			label := strings.Split(step, "-")[0]
			box := boxes[hash(label)]
			for n, v := range box {
				if v.label == label {
					newbox := append(box[:n], box[n+1:]...)
					boxes[hash(label)] = newbox
					break
				}
			}
		}
		printBoxes(boxes)
	}

	ans := 0
	for i, box := range boxes {
		for j, value := range box {
			n := (i + 1) * (j + 1) * value.value
			ans += n
		}
	}

	return ans, nil
}

func hash(step string) int {
	var h int

	h = 0
	for _, v := range step {
		ac := int(v)
		h += ac
		h *= 17
		h %= 256
	}

	return h
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines[0])
	fmt.Printf("Part one: %v\n", ans)

	ans, err = partTwo(lines[0])
	fmt.Printf("Part two: %v\n", ans)

}

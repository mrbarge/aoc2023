package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
	"strings"
)

func problem(lines [][]int, partTwo bool) (int, error) {

	sum := 0
	for _, line := range lines {
		ns := calcSequence(line, partTwo)
		sum += ns
	}
	return sum, nil
}

func calcSequence(i []int, partTwo bool) int {
	nextSeqs := make([][]int, 0)

	// go all the way down
	nextSeqs = append(nextSeqs, i)
	for {
		allZero := true
		nextSeq := make([]int, 0)
		lastSeq := len(nextSeqs) - 1
		for ix := 1; ix < len(nextSeqs[lastSeq]); ix++ {
			diff := nextSeqs[lastSeq][ix] - nextSeqs[lastSeq][ix-1]
			if diff != 0 {
				allZero = false
			}
			nextSeq = append(nextSeq, diff)
		}
		nextSeqs = append(nextSeqs, nextSeq)
		if allZero {
			break
		}
	}

	//fmt.Printf("%v\n", nextSeqs)
	lastDiff := 0
	if !partTwo {
		for ix := len(nextSeqs) - 1; ix > 0; ix-- {
			psl := len(nextSeqs[ix-1]) - 1
			lastDiff = lastDiff + nextSeqs[ix-1][psl]
			//fmt.Printf("  lastDiff: %v\n", lastDiff)
		}
	} else {
		for ix := len(nextSeqs) - 1; ix > 0; ix-- {
			lastDiff = nextSeqs[ix-1][0] - lastDiff
			//fmt.Printf("  lastDiff: %v\n", lastDiff)
		}
	}

	return lastDiff
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ilines := make([][]int, 0)
	for _, v := range lines {
		f := strings.Fields(v)
		iline := make([]int, 0)
		for _, fv := range f {
			i, _ := strconv.Atoi(fv)
			iline = append(iline, i)
		}
		ilines = append(ilines, iline)
	}

	ans, err := problem(ilines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(ilines, true)
	fmt.Printf("Part two: %v\n", ans)

}

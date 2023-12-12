package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strconv"
	"strings"
)

var seenCalls map[string]int

func problem(lines []string, partTwo bool) (int, error) {

	sum := 0
	seenCalls = make(map[string]int)
	for _, line := range lines {
		if partTwo {
			line = unfold(line)
		}
		springs := strings.Fields(line)[0]
		failures, _ := countMaxFailures(line)
		ans := step(springs, failures, 0)
		sum += ans
	}
	return sum, nil

}

func step(s string, failures []int, sequence int) (perms int) {
	callstr := fmt.Sprintf("%v-%v-%v", s, failures, sequence)
	if _, ok := seenCalls[callstr]; ok {
		return seenCalls[callstr]
	}

	//fmt.Printf("Called for string '%v' with failures %v and sequence %v\n", s, failures, sequence)
	failuresLeft := sum(failures)
	inFailureRun := sequence > 0
	if s == "" {
		if sequence == 0 && failuresLeft == 0 {
			seenCalls[callstr] = 1
			return 1
		}
		if inFailureRun && len(failures) == 1 && failures[0] == sequence {
			seenCalls[callstr] = 1
			return 1
		}
		seenCalls[callstr] = 0
		return 0
	}

	fl := countPossibleFailures(s)
	if inFailureRun {
		if sequence+fl < failuresLeft {
			// We can't satisfy the required number of failures with the spots remaining
			seenCalls[callstr] = 0
			return 0
		}
		if failuresLeft == 0 {
			// We have un-accounted for failures but no possible failure points in the string
			seenCalls[callstr] = 0
			return 0
		}
	} else {
		if fl < failuresLeft {
			seenCalls[callstr] = 0
			return 0
		}
	}

	operational := s[0] == '.'
	failure := s[0] == '#'
	wildcard := s[0] == '?'

	if operational && inFailureRun && sequence != failures[0] {
		// We expect a run of failures[0] failures but the chain is broken
		seenCalls[callstr] = 0
		return 0
	}
	if operational && inFailureRun && sequence == failures[0] {
		// We've satisfied a failure run so we can take it off the list of expected ones
		perms += step(s[1:], failures[1:], 0)
	}
	if wildcard && inFailureRun && sequence == failures[0] {
		// We can consume a wildcard to satisfy the expected failure run
		perms += step(s[1:], failures[1:], 0)
	}
	if wildcard || failure {
		if inFailureRun {
			// We can continue to satisfy part of the current expected failure run
			perms += step(s[1:], failures, sequence+1)
		} else {
			// We can commence a failure run
			perms += step(s[1:], failures, sequence+1)
		}
	}
	if (wildcard || operational) && !inFailureRun {
		// Move along to find the next failure opportunity
		perms += step(s[1:], failures, 0)
	}
	seenCalls[callstr] = perms
	return perms
}

func unfold(s string) string {

	p1 := strings.Fields(s)[0]
	p2 := strings.Fields(s)[1]

	var r1 = p1
	var r2 = p2
	for i := 0; i < 4; i++ {
		r1 += "?" + p1
		r2 += "," + p2
	}
	return r1 + " " + r2
}

// Determine the number of possible 'failure' points remaining in the string
func countPossibleFailures(s string) (failures int) {
	for _, v := range s {
		if v == '#' || v == '?' {
			failures += 1
		}
	}
	return failures
}

func sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return sum
}

func countMaxFailures(s string) (failures []int, max int) {
	failures = make([]int, 0)
	fails := strings.Split(strings.Split(s, " ")[1], ",")
	for _, v := range fails {
		iv, _ := strconv.Atoi(v)
		failures = append(failures, iv)
		max += iv
	}
	return failures, max
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

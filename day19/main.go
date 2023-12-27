package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	id        string
	condition []Condition
	def       string
}

type Condition struct {
	subject    string
	comparator string
	value      int
	dest       string
}

func problem1(input []string) (int, error) {
	rules := readBlock(input, false)
	data := readBlock(input, true)

	r := readRules(rules)
	d := readData(data)

	ans := 0
	for _, workflow := range d {
		accept := process(workflow, r)
		if accept {
			for _, wf := range workflow {
				ans += wf
			}
		}
	}
	return ans, nil
}

func problem2(input []string) (int, error) {
	rules := readBlock(input, false)
	//data := readBlock(input, true)

	r := readRules(rules)
	//d := readData(data)

	ans := 0
	minrange, maxrange := findRanges(r)

	for xi, xv := range minrange["x"] {
		for mi, mv := range minrange["m"] {
			for ai, av := range minrange["a"] {
				for si, sv := range minrange["s"] {
					data := map[string]int{
						"x": xv, "m": mv, "a": av, "s": sv,
					}
					accept := process(data, r)
					if accept {
						ans += (maxrange["x"][xi] - xv + 1) *
							(maxrange["m"][mi] - mv + 1) *
							(maxrange["a"][ai] - av + 1) *
							(maxrange["s"][si] - sv + 1)
					}

				}
			}
		}
	}
	fmt.Printf("%v\n", minrange)
	fmt.Printf("%v\n", maxrange)
	return ans, nil
}

func findRanges(rules []Rule) (map[string][]int, map[string][]int) {
	rmin := make(map[string][]int)
	rmax := make(map[string][]int)

	for _, rule := range rules {
		for _, condition := range rule.condition {
			if condition.comparator == ">" {
				rmin[condition.subject] = append(rmin[condition.subject], condition.value+1)
				rmax[condition.subject] = append(rmax[condition.subject], condition.value)
			} else if condition.comparator == "<" {
				rmin[condition.subject] = append(rmin[condition.subject], condition.value)
				rmax[condition.subject] = append(rmax[condition.subject], condition.value-1)
			}
		}
	}
	for k := range rmin {
		rmin[k] = append(rmin[k], 1)
		sort.Ints(rmin[k])
	}
	for k := range rmin {
		rmax[k] = append(rmax[k], 4000)
		sort.Ints(rmax[k])
	}
	return rmin, rmax
}

func (c Condition) test(data map[string]int) bool {
	if _, ok := data[c.subject]; !ok {
		return false
	}
	testval := data[c.subject]
	if c.comparator == ">" {
		return testval > c.value
	} else {
		return testval < c.value
	}
}

// p2: https://syltaen.com/advent-of-code/?year=2023&day=19

func process(data map[string]int, rules []Rule) bool {
	nextRule := "in"
	for {

		if nextRule == "A" {
			return true
		} else if nextRule == "R" {
			return false
		}

		for _, rule := range rules {
			if rule.id == nextRule {
				nextRule = ""
				for _, c := range rule.condition {
					if c.test(data) {
						nextRule = c.dest
						break
					}
				}
				if nextRule == "" {
					nextRule = rule.def
				} else {
					break
				}
			}
		}
	}
}

func readRules(rules []string) []Rule {
	ret := make([]Rule, 0)
	for _, rule := range rules {
		id := strings.Split(rule, "{")[0]
		r := Rule{
			id:        id,
			condition: make([]Condition, 0),
		}

		rawcond := strings.Split(rule[strings.Index(rule, "{")+1:len(rule)-1], ",")
		r.def = rawcond[len(rawcond)-1]
		for i, rc := range rawcond {
			if i == len(rawcond)-1 {
				// ignore last
				continue
			}
			for _, comparator := range []string{">", "<"} {
				ci := strings.Index(rc, comparator)
				di := strings.Index(rc, ":")
				if ci > 0 {
					k := rc[:ci]
					c := rc[ci : ci+1]
					t, _ := strconv.Atoi(rc[ci+1 : di])
					d := rc[di+1:]
					cond := Condition{
						subject:    k,
						comparator: c,
						value:      t,
						dest:       d,
					}
					r.condition = append(r.condition, cond)
				}

			}
		}
		ret = append(ret, r)
	}
	return ret
}

func readData(data []string) []map[string]int {
	d := make([]map[string]int, 0)
	for _, line := range data {
		dr := make(map[string]int)
		vals := strings.Split(line[1:len(line)-1], ",")
		for _, val := range vals {
			key := strings.Split(val, "=")[0]
			num, _ := strconv.Atoi(strings.Split(val, "=")[1])
			dr[key] = num
		}
		d = append(d, dr)
	}
	return d
}

func readBlock(input []string, second bool) []string {
	r := make([]string, 0)
	collect := false
	if !second {
		collect = true
	}
	for _, line := range input {
		if line == "" {
			collect = !collect
			continue
		}

		if collect {
			r = append(r, line)
		}
	}
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem1(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem2(lines)
	fmt.Printf("Part two: %v\n", ans)

}

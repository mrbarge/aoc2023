package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
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

func problem(input []string, partTwo bool) (int, error) {
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

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	//ans, err = problem(lines, true)
	//fmt.Printf("Part two: %v\n", ans)

}

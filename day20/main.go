package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"strings"
)

const HIGH = 1
const LOW = 0

const FLIPFLOP = 0
const CONJUNCTION = 1
const BROADCAST = 2

type Module struct {
	name           string
	lastPulseState map[string]bool // for conjuction pulse
	moduleType     int
	state          bool // for flipflop state
	connections    []*Module
}

type Pulse struct {
	state bool // false = low, true = high
	from  string
	to    string
}

var pulseQueue = make([]Pulse, 0)
var pulseCount = map[bool]int{
	false: 0,
	true:  0,
}

func (m *Module) receivePulse(p Pulse) {
	if m.moduleType == FLIPFLOP {
		if p.state {
			// ignore high pulses
			return
		}
		for _, dest := range m.connections {
			pulseQueue = append(pulseQueue, Pulse{
				from:  m.name,
				to:    dest.name,
				state: !m.state,
			})
			pulseCount[!m.state]++
		}
		m.state = !m.state
	}
	if m.moduleType == CONJUNCTION {
		m.lastPulseState[p.from] = p.state
		allHigh := true
		for _, pstate := range m.lastPulseState {
			if !pstate {
				allHigh = false
				break
			}
		}
		for _, dest := range m.connections {
			pulseQueue = append(pulseQueue, Pulse{
				from:  m.name,
				to:    dest.name,
				state: !allHigh,
			})
			pulseCount[!allHigh]++
		}
	}
	if m.moduleType == BROADCAST {
		for _, dest := range m.connections {
			pulseQueue = append(pulseQueue, Pulse{
				from:  m.name,
				to:    dest.name,
				state: p.state,
			})
			pulseCount[p.state]++
		}
	}
}

func problem(input []string, partTwo bool) (int, error) {
	modules, _ := readData(input)
	p2ans := 0
	for i := 0; i < 1000; i++ {
		pulseQueue = append(pulseQueue, Pulse{
			state: false,
			from:  "button",
			to:    "roadcaster",
		})
		pulseCount[false]++
		for {
			pulse := pulseQueue[0]

			if partTwo {
				if pulse.to == "rx" && !pulse.state {
					return p2ans, nil
				} else {
					p2ans++
					i -= 1 //loop infinitely..
				}
			}

			//fmt.Printf("%v -%v-> %v\n", pulse.from, pulse.state, pulse.to)
			modules[pulse.to].receivePulse(pulse)
			pulseQueue = pulseQueue[1:]

			if len(pulseQueue) == 0 {
				break
			}
		}
	}

	fmt.Printf("Low: %v, High: %v\n", pulseCount[false], pulseCount[true])
	return pulseCount[false] * pulseCount[true], nil
}

func readData(input []string) (map[string]*Module, *Module) {
	var broadcast *Module
	modules := make(map[string]*Module)

	for _, line := range input {
		elems := strings.Split(line, " -> ")
		name := elems[0][1:]
		var m *Module
		if _, ok := modules[name]; !ok {
			m = &Module{
				name:           name,
				connections:    make([]*Module, 0),
				lastPulseState: make(map[string]bool),
			}
		} else {
			m = modules[name]
		}
		modules[name] = m
		switch elems[0][0] {
		case '&':
			m.moduleType = CONJUNCTION
		case '%':
			m.moduleType = FLIPFLOP
		case 'b':
			m.moduleType = BROADCAST
			broadcast = m
		}
		dests := strings.Split(elems[1], ", ")
		for _, dest := range dests {
			if _, ok := modules[dest]; !ok {
				dm := &Module{
					name:           dest,
					connections:    make([]*Module, 0),
					lastPulseState: make(map[string]bool),
				}
				modules[dest] = dm
			}
			modules[dest].lastPulseState[name] = false
			m.connections = append(m.connections, modules[dest])
		}
	}
	return modules, broadcast
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

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)

}

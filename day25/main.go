package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"strings"
)

type Connection struct {
	w1     *Wire
	w2     *Wire
	walked int
}

type Wire struct {
	name        string
	connections []*Wire
}

func (c *Connection) Matches(w1 *Wire, w2 *Wire) bool {
	return (c.w1 == w1 && c.w2 == w2) || (c.w2 == w1 && c.w1 == w2)
}

func (c *Connection) Key() string {
	if c.w1.name < c.w2.name {
		return fmt.Sprintf("%v:%v", c.w1.name, c.w2.name)
	} else {
		return fmt.Sprintf("%v:%v", c.w2.name, c.w1.name)
	}
}

func problem(input []string) (int, error) {
	data := readData(input)

	c := connections(data)
	printit(c)
	mc1 := mostWalked(c)
	removeConnection(data, mc1)

	c = connections(data)
	mc2 := mostWalked(c)
	removeConnection(data, mc2)

	c = connections(data)
	mc3 := mostWalked(c)
	removeConnection(data, mc3)

	pool1Size := calcPoolSize(mc1)
	pool2Size := len(data) - pool1Size
	ans := pool1Size * pool2Size
	return ans, nil
}

func printit(i map[string]*Connection) {
	for _, v := range i {
		fmt.Printf("{%v %v: %v} ", v.w1.name, v.w2.name, v.walked)
	}
}

func calcPoolSize(c *Connection) int {
	seen := make(map[string]bool)
	next := make([]*Wire, 0)

	next = append(next, c.w1)
	for len(next) > 0 {
		nextwire := next[0]
		next = next[1:]

		for _, connection := range nextwire.connections {
			if _, ok := seen[connection.name]; ok {
				continue
			}
			next = append(next, connection)
			seen[connection.name] = true
		}
	}
	return len(seen)
}

func removeConnection(wires map[string]*Wire, c *Connection) {
	for _, wire := range wires {
		if wire == c.w1 || wire == c.w2 {
			tc := make([]*Wire, 0)
			for _, next := range wire.connections {
				if next != c.w1 && next != c.w2 {
					tc = append(tc, next)
				}
			}
			wire.connections = tc
		}
	}
}

func connections(in map[string]*Wire) map[string]*Connection {
	walkedWires := make(map[string]*Connection)
	for _, wire := range in {
		walkWires(in, wire, walkedWires)
	}
	return walkedWires
}

func mostWalked(i map[string]*Connection) *Connection {
	max := math.MinInt
	var maxConn *Connection
	for _, v := range i {
		if v.walked > max {
			max = v.walked
			maxConn = v
		}
	}
	return maxConn
}

func walkWires(in map[string]*Wire, w *Wire, walked map[string]*Connection) {
	seen := make(map[string]bool)
	next := make([]*Wire, 0)

	next = append(next, in[w.name])
	for len(next) > 0 {
		nextwire := next[0]
		next = next[1:]

		for _, connection := range nextwire.connections {
			if _, ok := seen[connection.name]; ok {
				continue
			}
			next = append(next, connection)
			seen[connection.name] = true

			conn := Connection{w1: nextwire, w2: connection}
			if _, ok := walked[conn.Key()]; !ok {
				walked[conn.Key()] = &conn
			}
			walked[conn.Key()].walked++
		}
	}
}

func readData(input []string) map[string]*Wire {
	seen := make(map[string]*Wire)
	for _, line := range input {
		head := strings.Split(line, ": ")[0]
		elems := strings.Fields(strings.Split(line, ": ")[1])
		var w *Wire
		if _, ok := seen[head]; !ok {
			w = &Wire{
				name:        head,
				connections: make([]*Wire, 0),
			}
			seen[head] = w
		}
		w = seen[head]
		for _, wire := range elems {
			if _, ok := seen[wire]; !ok {
				w2 := &Wire{
					name:        wire,
					connections: make([]*Wire, 0),
				}
				seen[wire] = w2
			}
			w2 := seen[wire]
			w.connections = append(w.connections, w2)
			w2.connections = append(w2.connections, w)
		}
	}
	return seen
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Part one: %v\n", ans)

}

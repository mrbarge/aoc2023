package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"sort"
	"strings"
)

type Blocks []Block
type Block struct {
	coords []helper.Coord3D
}

func (b Blocks) Len() int {
	return len(b)
}
func (b Blocks) Less(i, j int) bool {
	var mini, minj = math.MaxInt, math.MaxInt
	for _, v := range b[i].coords {
		if v.Z < mini {
			mini = v.Z
		}
	}
	for _, v := range b[j].coords {
		if v.Z < minj {
			minj = v.Z
		}
	}
	return mini < minj
}

func (b Blocks) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Block) Grounded() bool {
	for _, c := range b.coords {
		if c.Z == 1 {
			return true
		}
	}
	return false
}

func (b Block) Lower() Block {
	nb := Block{
		coords: make([]helper.Coord3D, 0),
	}
	for _, c := range b.coords {
		nb.coords = append(nb.coords, helper.Coord3D{X: c.X, Y: c.Y, Z: c.Z - 1, V: c.V})
	}
	return nb
}

func (b Block) CanLower(blocks Blocks) bool {
	for _, block := range blocks {
		if block.coords[0].V == b.coords[0].V {
			// ignore selfblock
			continue
		}

		for _, coord := range block.coords {
			for _, bc := range b.coords {
				if bc.X == coord.X && bc.Y == coord.Y && bc.Z-1 == coord.Z {
					return false
				}
			}
		}
	}
	return true
}

func (b Block) Parents(blocks Blocks) []int {
	r := make([]int, 0)
	for i, block := range blocks {
		if block.coords[0].V == b.coords[0].V {
			// ignore selfblock
			continue
		}
		for _, coord := range block.coords {
			for _, bc := range b.coords {
				if bc.X == coord.X && bc.Y == coord.Y && bc.Z+1 == coord.Z {
					r = append(r, i)
				}
			}
		}
	}
	return r
}

func (b Block) Children(blocks Blocks) []int {
	r := make([]int, 0)
	for i, block := range blocks {
		if block.coords[0].V == b.coords[0].V {
			// ignore selfblock
			continue
		}
		for _, coord := range block.coords {
			for _, bc := range b.coords {
				if bc.X == coord.X && bc.Y == coord.Y && bc.Z-1 == coord.Z {
					r = append(r, i)
				}
			}
		}
	}
	return r
}

func readBlock(s string) Block {
	var b Block
	b.coords = make([]helper.Coord3D, 0)
	fromCoord := helper.ReadCoord3D(strings.Split(s, "~")[0])
	toCoord := helper.ReadCoord3D(strings.Split(s, "~")[1])

	for zi := fromCoord.Z; zi <= toCoord.Z; zi++ {
		for yi := fromCoord.Y; yi <= toCoord.Y; yi++ {
			for xi := fromCoord.X; xi <= toCoord.X; xi++ {
				ci := helper.Coord3D{X: xi, Y: yi, Z: zi, V: 0}
				b.coords = append(b.coords, ci)
			}
		}
	}
	return b
}

func (b Block) CanDisintegrate(blocks Blocks) bool {
	parents := b.Parents(blocks)
	if len(parents) == 0 {
		return true
	}
	for _, parent := range parents {
		// If parent has another child that isn't this block we're good
		pc := blocks[parent].Children(blocks)
		if len(pc) == 1 {
			return false
		}
	}
	return true
}

func problem(input []string, partTwo bool) (int, error) {
	blocks := readData(input)
	sort.Sort(blocks)
	newblocks := fallBricks(blocks)
	ans := 0
	for i := len(newblocks) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", newblocks[i])
	}
	for _, block := range newblocks {
		if block.CanDisintegrate(newblocks) {
			ans++
		}
	}
	return ans, nil
}

func readData(input []string) Blocks {
	blocks := make([]Block, 0)
	for i, line := range input {
		b := readBlock(line)
		for j, _ := range b.coords {
			b.coords[j].V = i
		}
		blocks = append(blocks, b)
	}
	return blocks
}

func fallBricks(blocks Blocks) Blocks {
	done := false
	for !done {
		moved := false
		var newblocks Blocks = make([]Block, 0)

		for _, block := range blocks {
			if block.Grounded() {
				newblocks = append(newblocks, block)
				continue
			}
			if block.CanLower(blocks) {
				newblocks = append(newblocks, block.Lower())
				moved = true
			} else {
				newblocks = append(newblocks, block)
			}
		}
		sort.Sort(newblocks)
		blocks = newblocks

		if !moved {
			done = true
		}
	}
	return blocks
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

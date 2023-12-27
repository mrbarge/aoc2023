package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
)

type Blocks []Block
type Block struct {
	id     int
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
		id:     b.id,
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
	for _, block := range blocks {
		if block.coords[0].V == b.coords[0].V {
			// ignore selfblock
			continue
		}
		for _, coord := range block.coords {
			for _, bc := range b.coords {
				if bc.X == coord.X && bc.Y == coord.Y && bc.Z+1 == coord.Z {
					if !slices.Contains(r, block.coords[0].V) {
						r = append(r, block.coords[0].V)
					}
				}
			}
		}
	}
	return r
}

func (b Block) Children(blocks Blocks) []int {
	r := make([]int, 0)
	for _, block := range blocks {
		if block.coords[0].V == b.coords[0].V {
			// ignore selfblock
			continue
		}
		for _, coord := range block.coords {
			for _, bc := range b.coords {
				if bc.X == coord.X && bc.Y == coord.Y && bc.Z-1 == coord.Z {
					if !slices.Contains(r, block.coords[0].V) {
						r = append(r, block.coords[0].V)
					}
				}
			}
		}
	}
	return r
}

func readBlock(s string, id int) Block {
	var b Block
	b.coords = make([]helper.Coord3D, 0)
	b.id = id
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

func (b Block) CountParents(blocks Blocks) int {
	queue := make([]Block, 0)

	seen := make(map[int]bool)

	sp := b.Parents(blocks)
	for _, v := range sp {
		queue = append(queue, blocks[v])
		seen[blocks[v].id] = true
	}

	np := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		pp := p.Parents(blocks)
		for _, bp := range pp {
			pblock := blocks[bp]
			if _, ok := seen[pblock.id]; !ok {
				queue = append(queue, pblock)
				seen[pblock.id] = true
			}
		}
		np++
	}

	return np
}

func partone(input []string) (int, error) {
	blocks := readData(input)
	sort.Sort(blocks)
	newblocks := fallBricks(blocks)
	ans := 0
	for _, nb := range newblocks {
		fmt.Printf("Block %v: %v\n", nb.id, nb.coords)
	}
	for _, block := range newblocks {
		if block.CanDisintegrate(newblocks) {
			fmt.Printf("Block %v can disintegrate\n", block.coords[0].V)
			ans++
		}
	}
	return ans, nil
}

func parttwo(input []string) (int, error) {
	blocks := readData(input)
	sort.Sort(blocks)
	newblocks := fallBricks(blocks)
	ans := 0
	for _, block := range newblocks {
		if !block.CanDisintegrate(newblocks) {
			ans += block.CountParents(newblocks)
		}
	}
	return ans, nil
}

func readData(input []string) Blocks {
	blocks := make([]Block, 0)
	for i, line := range input {
		b := readBlock(line, i)
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

	ans, err := partone(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}

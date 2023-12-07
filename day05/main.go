package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

type MapRange struct {
	SourceStart int
	SourceEnd   int
	DestStart   int
	DestEnd     int
}

type ResourceMap struct {
	Source string
	Dest   string
	Ranges []MapRange
}

func problem(lines []string, partTwo bool) (int, error) {
	seeds, maps := readMaps(lines, partTwo)

	order := []string{
		"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location",
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		val := seed
		oi := 0
		foundLocation := false
		for !foundLocation {
			// find the next map
			var nm ResourceMap
			for _, m := range maps {
				if m.Source == order[oi] {
					nm = m
				}
			}

			for _, r := range nm.Ranges {
				if val >= r.SourceStart && val <= r.SourceEnd {
					valdif := val - r.SourceStart
					val = r.DestStart + valdif
					break
				}
			}
			oi++
			if nm.Dest == "location" {
				foundLocation = true
			}
		}
		if val < lowest {
			lowest = val
		}
	}
	return lowest, nil
}

func readMaps(lines []string, partTwo bool) ([]int, []ResourceMap) {
	seeds := make([]int, 0)
	resourcemaps := make([]ResourceMap, 0)

	var currentMap *ResourceMap

	for _, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}

		si := strings.Index(line, "seeds: ")
		if si >= 0 {
			tmpseeds := make([]int, 0)
			sline := strings.Fields(line[7:])
			for _, v := range sline {
				si, _ := strconv.Atoi(v)
				tmpseeds = append(tmpseeds, si)
			}
			if partTwo {
				for i := 0; i < len(tmpseeds); i++ {
					sr := tmpseeds[i+1]
					for j := tmpseeds[i]; j < tmpseeds[i]+sr; j++ {
						seeds = append(seeds, j)
					}
					i++
				}
			} else {
				seeds = tmpseeds
			}
			continue
		}

		si = strings.Index(line, "map:")
		if si >= 0 {
			if currentMap == nil {
				currentMap = &ResourceMap{}
			} else {
				resourcemaps = append(resourcemaps, *currentMap)
			}
			mapname := strings.Fields(line)[0]
			currentMap.Source = mapname[0:strings.Index(mapname, "-")]
			currentMap.Dest = mapname[strings.LastIndex(mapname, "-")+1:]
			currentMap.Ranges = make([]MapRange, 0)
			continue
		}

		// must be ranges now
		rline := strings.Fields(line)
		deststart, _ := strconv.Atoi(rline[0])
		sourcestart, _ := strconv.Atoi(rline[1])
		valrange, _ := strconv.Atoi(rline[2])

		r := MapRange{
			SourceStart: sourcestart,
			SourceEnd:   sourcestart + valrange,
			DestStart:   deststart,
			DestEnd:     deststart + valrange,
		}
		currentMap.Ranges = append(currentMap.Ranges, r)
	}
	resourcemaps = append(resourcemaps, *currentMap)

	//for _, v := range resourcemaps {
	//	fmt.Printf("%v,%v\n", v.Source, v.Dest)
	//	for _, v2 := range v.Ranges {
	//		fmt.Printf("%v,%v,%v,%v\n", v2.DestStart, v2.SourceStart, v2.DestEnd, v2.SourceEnd)
	//	}
	//}
	return seeds, resourcemaps
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, true)
	fmt.Printf("Part one: %v\n", ans)

}

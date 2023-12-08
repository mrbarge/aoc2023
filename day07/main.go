package main

import (
	"fmt"
	"github.com/mrbarge/aoc2023/helper"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	bid   int
	cards []int
	raw   string
}

type HandList []Hand

var cardMap = map[uint8]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const FIVE_OF_A_KIND = 100
const FOUR_OF_A_KIND = 90
const FULL_HOUSE = 80
const THREE_OF_A_KIND = 70
const TWO_PAIR = 60
const ONE_PAIR = 50
const HIGH_CARD = 10

func (list HandList) Len() int {
	return len(list)
}

func (list HandList) Less(i, j int) bool {
	ti := list[i].TypeP2()
	tj := list[j].TypeP2()
	if ti < tj {
		return true
	} else if ti > tj {
		return false
	}

	// They're equal, so, compare further
	for x := 0; x < len(list[i].cards); x++ {
		if list[i].cards[x] < list[j].cards[x] {
			return true
		} else if list[i].cards[x] > list[j].cards[x] {
			return false
		}
	}

	return true
}

func (list HandList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func genPermutations(s string) []string {

	if !strings.ContainsRune(s, 'J') {
		return []string{s}
	}

	perms := []string{}
	for _, v := range []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"} {
		rp := strings.Replace(s, "J", v, 1)
		perms = append(perms, genPermutations(rp)...)
	}

	return perms
}

func (h Hand) TypeP2() int {

	maxType := 0
	handPerms := genPermutations(h.raw)

	for _, perm := range handPerms {

		t := getType(perm)
		if t > maxType {
			maxType = t
		}
	}

	return maxType
}

func getType(s string) int {
	ci := make([]int, 0)
	for i, _ := range s {
		ci = append(ci, cardMap[s[i]])
	}

	freq := make(map[int]int)
	for _, v := range ci {
		if _, ok := freq[v]; ok {
			freq[v]++
		} else {
			freq[v] = 1
		}
	}

	if len(freq) == 1 {
		return FIVE_OF_A_KIND
	} else if len(freq) == 2 {
		for _, v := range freq {
			if v == 4 {
				return FOUR_OF_A_KIND
			} else if v == 3 {
				return FULL_HOUSE
			}
		}
	} else if len(freq) == 3 {
		for _, v := range freq {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	} else if len(freq) == 4 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func (h Hand) Type() int {
	freq := make(map[int]int)
	for _, v := range h.cards {
		if _, ok := freq[v]; ok {
			freq[v]++
		} else {
			freq[v] = 1
		}
	}

	if len(freq) == 1 {
		return FIVE_OF_A_KIND
	} else if len(freq) == 2 {
		for _, v := range freq {
			if v == 4 {
				return FOUR_OF_A_KIND
			} else if v == 3 {
				return FULL_HOUSE
			}
		}
	} else if len(freq) == 3 {
		for _, v := range freq {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	} else if len(freq) == 4 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func problem(lines []string, partTwo bool) (int, error) {

	hands := HandList{}
	for _, line := range lines {
		sf := strings.Fields(line)
		bid, _ := strconv.Atoi(sf[1])
		c := make([]int, 0)
		for i := 0; i < len(sf[0]); i++ {
			c = append(c, cardMap[sf[0][i]])
		}
		h := Hand{
			cards: c,
			bid:   bid,
			raw:   sf[0],
		}
		hands = append(hands, h)
	}
	sort.Sort(hands)

	sum := 0
	for i, v := range hands {
		sum += (i + 1) * v.bid
	}
	return sum, nil
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

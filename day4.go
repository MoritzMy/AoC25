package main

import (
	"AoC2025/util"
	"fmt"
)

func main() {
	star1()
	star2()
}

func star1() {
	raw, err := util.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}

	paperMap := make([][]rune, len(raw))
	for i := range raw {
		paperMap[i] = []rune(raw[i])
	}

	count := 0
	count = getSinglePosition(paperMap)
	fmt.Println(count)
}

func star2() {
	raw, err := util.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}

	paperMap := make([][]rune, len(raw))
	for i := range raw {
		paperMap[i] = []rune(raw[i])
	}

	count := 0

	for {
		removed := getSinglePosition(paperMap)
		if removed == 0 {
			break
		}
		count += removed
	}
	fmt.Println(count)
}

func getSinglePosition(paperMap [][]rune) int {
	toRemove := make([][2]int, 0)
	for y := 0; y < len(paperMap); y++ {

		for x := 0; x < len(paperMap[y]); x++ {
			if paperMap[y][x] == '@' && validateSurrounding(paperMap, x, y) {
				toRemove = append(toRemove, [2]int{x, y})
			}
		}
	}
	for _, p := range toRemove {
		paperMap[p[1]][p[0]] = '.'
	}
	return len(toRemove)
}

func validateSurrounding(paperMap [][]rune, x int, y int) bool {
	count := 0
	lowX := max(0, x-1)
	highX := min(len(paperMap[0])-1, x+1)

	lowY := max(0, y-1)
	highY := min(len(paperMap)-1, y+1)

	for yy := lowY; yy <= highY; yy++ {
		for xx := lowX; xx <= highX; xx++ {
			if paperMap[yy][xx] == '@' || paperMap[yy][xx] == 'x' {
				count++
			}
		}
	}
	count--
	return count < 4
}

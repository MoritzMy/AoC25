package main

import (
	"AoC2025/util"
	"fmt"
	"math"
)

func main() {
	star1()
	star2()
}

func star1() {
	lines, err := util.ReadFile("input/day3.txt")
	if err != nil {
		panic(err)
	}
	joltage := 0
	for _, line := range lines {
		firstMax, firstIndex := getMax(line, 1)
		secondMax, _ := getMax(line[firstIndex+1:], 0)
		joltage += firstMax*10 + secondMax
	}
	fmt.Println(joltage)
}

func star2() {
	lines, err := util.ReadFile("input/day3.txt")
	if err != nil {
		panic(err)
	}
	joltage := 0
	for _, line := range lines {
		previousIndex := -1

		bankMaxPower := 0
		for i := 11; i >= 0; i-- {
			line = line[previousIndex+1:]
			maxPower, index := getMax(line, i)
			bankMaxPower += maxPower * int(math.Pow10(i))
			previousIndex = index
		}
		joltage += bankMaxPower
	}
	fmt.Println(joltage)
}

func getMax(bank string, excludeLast int) (int, int) {
	limit := len(bank) - excludeLast
	if limit < 0 {
		limit = 0
	}

	maxVal := -1
	maxIndex := -1

	for i := 0; i < limit; i++ {
		num := int(bank[i] - '0')
		if num > maxVal {
			maxVal = num
			maxIndex = i
		}
	}

	return maxVal, maxIndex
}

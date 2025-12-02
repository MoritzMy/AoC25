package main

import (
	"fmt"
	"strconv"
	"time"

	"AoC2025/util"
)

func main() {
	start := time.Now()
	star1()
	end := time.Now()
	fmt.Println("Solution 1:", end.Sub(start))
	start = time.Now()
	star2()
	end = time.Now()
	fmt.Println("Solution 2:", end.Sub(start))
}

func star1() {
	lines, err := util.ReadFile("input/day1.txt")

	if err != nil {
		panic(err)
	}

	nullCounter := 0
	dialCount := 50

	for _, line := range lines {
		direction := line[0]
		distance, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			distance *= -1
		}
		dialCount = (dialCount + distance) % 100
		if dialCount == 0 {
			nullCounter++
		}
	}

	fmt.Println(nullCounter)
}

func star2() {
	lines, err := util.ReadFile("input/day1.txt")

	if err != nil {
		panic(err)
	}

	nullCounter := 0
	dialCount := 50

	for _, line := range lines {
		direction := rune(line[0])
		distance, err := strconv.Atoi(line[1:])
		rotation := Rotation{direction: direction, distance: distance}
		nullCounter += rotation.distance / 100
		rotation.distance %= 100
		if err != nil {
			panic(err)
		}
		if rotation.direction == 'L' {
			rotation.distance *= -1
		}
		oldDialCount := dialCount
		dialCount = dialCount + rotation.distance
		nullCounter += countZeroPasses(oldDialCount, dialCount)

		dialCount = normalizeDial(dialCount)
	}

	fmt.Println(nullCounter)
}

func countZeroPasses(oldCount int, newCount int) int {
	if oldCount > 0 && newCount < 0 {
		return 1
	} else if newCount > 99 {
		return 1
	} else if newCount == 0 {
		return 1
	}
	return 0
}

func normalizeDial(dial int) int {
	return (dial%100 + 100) % 100
}

type Rotation struct {
	direction rune
	distance  int
}

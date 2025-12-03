package main

import (
	"AoC2025/util"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
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
	lines, err := util.ReadFile("input/day2.txt")
	invalidIdCount := 0
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		for _, numRange := range strings.Split(line, ",") {
			nums := strings.Split(numRange, "-")
			start := nums[0]
			end := nums[1]
			endInt, _ := strconv.Atoi(end)
			startInt, _ := strconv.Atoi(start)
			for i := startInt; i <= endInt; i++ {
				indexString := strconv.Itoa(i)
				if len(indexString)%2 == 1 {
					continue
				} else if indexString[0] == '0' {
					invalidIdCount += i
					continue
				}

				midpoint := len(indexString) / 2
				if indexString[:midpoint] == indexString[midpoint:] {
					invalidIdCount += i
				}

			}
		}
	}
	fmt.Println("STAR 1 SOLUTION: ", invalidIdCount)

}

func star2() {
	lines, err := util.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}
	invalidIdCount := 0
	for _, line := range lines {
		for _, numRange := range strings.Split(line, ",") {
			nums := strings.Split(numRange, "-")
			start := nums[0]
			end := nums[1]
			endInt, _ := strconv.Atoi(end)
			startInt, _ := strconv.Atoi(start)
			for i := startInt; i <= endInt; i++ {
				indexString := strconv.Itoa(i)
				for sectionLength := 1; sectionLength < len(indexString)/2+1; sectionLength++ {
					sections, err := sectionizeString(indexString, sectionLength)
					if err != nil {
						continue
					}
					if validateSectionEquality(sections) {
						invalidIdCount += i
						break
					}
				}
			}
		}
	}
	fmt.Println("STAR 2 SOLUTION: ", invalidIdCount)
}

func sectionizeString(str string, sectionSize int) ([]string, error) {
	if len(str)%sectionSize != 0 {
		return nil, errors.New("String can't be sectioned - String is not a multiple of section size")
	}
	iterations := len(str) / sectionSize
	sections := make([]string, 0)
	for i := 0; i < iterations; i++ {
		section := str[0:sectionSize]
		str = str[sectionSize:]
		sections = append(sections, section)

	}
	return sections, nil
}

func validateSectionEquality(sections []string) bool {
	for _, section := range sections {
		if section != sections[0] {
			return false
		}
	}
	return true
}

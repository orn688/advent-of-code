package day01

import (
	"slices"
	"strconv"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	rows, err := util.ParseGridOfIntegers(input, 2)
	if err != nil {
		return "", err
	}

	var left, right []int
	for _, row := range rows {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance int
	for i := range left {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}
		totalDistance += dist
	}

	return strconv.Itoa(totalDistance), nil
}

func Part2(input string) (string, error) {
	rows, err := util.ParseGridOfIntegers(input, 2)
	if err != nil {
		return "", err
	}

	var left, right []int
	for _, row := range rows {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	rightCounts := map[int]int{}
	for _, val := range right {
		rightCounts[val]++
	}

	var total int
	for _, val := range left {
		total += val * rightCounts[val]
	}

	return strconv.Itoa(total), nil
}

package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	rows, err := parseInput(input)
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
	rows, err := parseInput(input)
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

func parseInput(input string) ([][2]int, error) {
	var parsed [][2]int
	for _, line := range strings.Split(input, "\n") {
		aString, bString, ok := strings.Cut(line, "   ")
		a, aErr := strconv.Atoi(aString)
		b, bErr := strconv.Atoi(bString)
		if !ok || aErr != nil || bErr != nil {
			return nil, fmt.Errorf("malformed line: %q", line)
		}

		parsed = append(parsed, [2]int{a, b})
	}
	return parsed, nil
}

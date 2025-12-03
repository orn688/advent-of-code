package day02

import (
	"cmp"
	"strconv"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	grid, err := util.ParseGridOfIntegers(input, 0)
	if err != nil {
		return "", err
	}

	var safeCount int
	for _, row := range grid {
		if safeRowPart1(row) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount), nil
}

func safeRowPart1(row []int) bool {
	if len(row) <= 1 {
		return true
	}

	order := cmp.Compare(row[0], row[1])
	for i := 1; i < len(row); i++ {
		prev, curr := row[i-1], row[i]
		if order != cmp.Compare(prev, curr) {
			return false
		}
		diff := (prev - curr) * order
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func Part2(input string) (string, error) {
	grid, err := util.ParseGridOfIntegers(input, 0)
	if err != nil {
		return "", err
	}

	var safeCount int
	for _, row := range grid {
		if safeRowPart2(row) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount), nil
}

func safeRowPart2(row []int) bool {
	// Any length-2 row is safe because any length-1 row is safe, and a row is
	// considered safe if removing an element will make it safe.
	if len(row) <= 2 {
		return true
	}

	var totalCmp int
	for i := 1; i < len(row); i++ {
		prev, curr := row[i-1], row[i]
		totalCmp += cmp.Compare(prev, curr)
	}

	return true
}

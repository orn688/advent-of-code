package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func MustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("Invalid integer value %q", s)
	}
	return val
}

func ParseGridOfIntegers(input string, expectedRowLength int) ([][]int, error) {
	var grid [][]int
	for line := range strings.SplitSeq(input, "\n") {
		var row []int
		for entry := range strings.SplitSeq(line, " ") {
			// Skip blanks, which indicates there were multiple spaces between
			// integers.
			if entry == "" {
				continue
			}
			val, err := strconv.Atoi(entry)
			if err != nil {
				return nil, fmt.Errorf("malformed line: %q", line)
			}
			row = append(row, val)
		}
		if expectedRowLength > 0 && len(row) != expectedRowLength {
			return nil, fmt.Errorf("malformed line: %q", line)
		}

		grid = append(grid, row)
	}
	return grid, nil
}

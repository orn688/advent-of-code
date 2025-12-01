package day01

import (
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	var zeroCount int

	curr := 50
	for line := range strings.SplitSeq(input, "\n") {
		dir := line[0]
		count := util.MustParseInt(line[1:])
		switch dir {
		case 'L':
			curr -= count
		case 'R':
			curr += count
		}
		curr = modulo(curr, 100)
		if curr == 0 {
			zeroCount++
		}
	}

	return strconv.Itoa(zeroCount), nil
}

func Part2(input string) (string, error) {
	var zeroCount int

	curr := 50
	for line := range strings.SplitSeq(input, "\n") {
		var increment int
		switch line[0] {
		case 'L':
			increment = -1
		case 'R':
			increment = 1
		default:
			return "", fmt.Errorf("invalid line: %q", line)
		}
		count := util.MustParseInt(line[1:])
		for range count {
			curr += increment
			curr = modulo(curr, 100)
			if curr == 0 {
				zeroCount++
			}
		}
	}

	return strconv.Itoa(zeroCount), nil
}

// modulo takes the modulo of val / divisor. Unlike Go's % operator which
// returns a remainder (which may be negative), this function will never return
// a negative result.
func modulo(val, divisor int) int {
	m := val % divisor
	if m < 0 {
		m += divisor
	}
	return m
}

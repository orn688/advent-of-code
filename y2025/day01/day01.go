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
		dir := line[0]
		count := util.MustParseInt(line[1:])
		for count > 0 {
			switch dir {
			case 'L':
				curr--
			case 'R':
				curr++
			}
			curr = modulo(curr, 100)
			if curr == 0 {
				zeroCount++
			}
			count--
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

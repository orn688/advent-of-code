package day05

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "3")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "14")
}

package day04

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "13")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "43")
}

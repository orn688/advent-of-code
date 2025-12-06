package day06

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "4277556")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "3263827")
}

package day09

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "50")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "24")
}

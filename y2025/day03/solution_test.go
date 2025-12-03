package day03

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "357")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "3121910778619")
}

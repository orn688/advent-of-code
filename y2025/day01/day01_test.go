package day01

import (
	"testing"

	"github.com/orn688/advent-of-code/util"
)

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "3")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "6")
}

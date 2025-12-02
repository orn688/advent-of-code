package day02

import (
	"strings"
	"testing"

	"github.com/orn688/advent-of-code/util"
)

var testInput = strings.Join([]string{
	"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,",
	"1698522-1698528,446443-446449,38593856-38593862,565653-565659,",
	"824824821-824824827,2121212118-2121212124",
}, "")

func TestPart1(t *testing.T) {
	util.CheckAnswer(t, Part1, testInput, "1227775554")
}

func TestPart2(t *testing.T) {
	util.CheckAnswer(t, Part2, testInput, "4174379265")
}

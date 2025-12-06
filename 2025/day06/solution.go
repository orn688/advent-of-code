package day06

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	grid, err := util.ParseGridOfIntegers(strings.Join(lines[:len(lines)-1], "\n"), 0)
	if err != nil {
		return "", err
	}

	var ops []byte
	for _, char := range []byte(lines[len(lines)-1]) {
		if char != ' ' {
			ops = append(ops, byte(char))
		}
	}

	var sum int

	for x := range grid[0] {
		var nums []int
		for y := range grid {
			nums = append(nums, grid[y][x])
		}
		sum += computeOperation(nums, ops[x])
	}

	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	opLine := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	var maxLineLength int
	for _, line := range lines {
		maxLineLength = max(maxLineLength, len(line))
	}

	var opColumns []int
	var ops []byte
	opRE := regexp.MustCompile(`[+*]`)
	for _, submatch := range opRE.FindAllStringIndex(opLine, -1) {
		index := submatch[0]
		opColumns = append(opColumns, index)
		ops = append(ops, opLine[index])
	}

	var sum int

	for i, startCol := range opColumns {
		var endCol int
		if i+1 >= len(opColumns) {
			endCol = maxLineLength
		} else {
			// minus 1 to account for the column of spaces.
			endCol = opColumns[i+1] - 1
		}

		var nums []int
		for j := endCol; j >= startCol; j-- {
			var buf strings.Builder
			for _, line := range lines {
				if j >= len(line) || line[j] == ' ' {
					continue
				}
				buf.WriteByte(line[j])
			}
			if buf.Len() > 0 {
				nums = append(nums, util.MustParseInt(buf.String()))
			}
		}

		sum += computeOperation(nums, ops[i])
	}
	return strconv.Itoa(sum), nil
}

func computeOperation(nums []int, operation byte) int {
	var result int
	var f func(a, b int) int
	switch operation {
	case '+':
		f = func(a, b int) int { return a + b }
		result = 0
	case '*':
		f = func(a, b int) int { return a * b }
		result = 1
	default:
		log.Panicf("invalid operation %q", operation)
	}

	for _, num := range nums {
		result = f(result, num)
	}
	return result
}

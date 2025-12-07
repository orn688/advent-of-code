package day07

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

const (
	startChar = 'S'
	splitChar = '^'
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	startColumn := strings.IndexRune(lines[0], startChar)
	if startColumn < 0 {
		return "", errors.New("no 'S' in first line of input")
	}

	var splits int

	beamColumns := []int{startColumn}
	for _, line := range lines[1:] {
		var newBeamColumns []int
		for _, col := range beamColumns {
			if line[col] == splitChar {
				newBeamColumns = append(newBeamColumns, col-1, col+1)
				splits++
			} else {
				newBeamColumns = append(newBeamColumns, col)
			}
		}
		beamColumns = slices.Compact(newBeamColumns)
	}

	return strconv.Itoa(splits), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	startColumn := strings.IndexRune(lines[0], startChar)
	if startColumn < 0 {
		return "", errors.New("no 'S' in first line of input")
	}

	type beam struct {
		column int
		// number of possible paths the beam could have taken to get to its
		// current point.
		possiblePaths int
	}

	beams := []beam{{column: startColumn, possiblePaths: 1}}
	for _, line := range lines[1:] {
		var newBeams []beam
		for _, b := range beams {
			if line[b.column] == splitChar {
				newBeams = append(newBeams,
					beam{column: b.column - 1, possiblePaths: b.possiblePaths},
					beam{column: b.column + 1, possiblePaths: b.possiblePaths},
				)
			} else {
				newBeams = append(newBeams, b)
			}
		}

		beams = nil

		for _, b := range newBeams {
			lastIndex := len(beams) - 1
			if len(beams) == 0 || beams[lastIndex].column != b.column {
				beams = append(beams, b)
			} else {
				// Merge two beams.
				beams[lastIndex].possiblePaths += b.possiblePaths
			}
		}
	}

	var totalPossiblePaths int
	for _, b := range beams {
		totalPossiblePaths += b.possiblePaths
	}

	return strconv.Itoa(totalPossiblePaths), nil
}

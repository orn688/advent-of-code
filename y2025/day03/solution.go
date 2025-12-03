package day03

import (
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	var total int
	for bank := range strings.SplitSeq(input, "\n") {
		total += maxPossibleJoltage(bank, 2)
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for bank := range strings.SplitSeq(input, "\n") {
		total += maxPossibleJoltage(bank, 12)
	}
	return strconv.Itoa(total), nil
}

func maxPossibleJoltage(bank string, numEnabled int) int {
	allBatteries := []byte(bank)

	var toEnable []byte
	lastEnabledIndex := -1

	// Say there are 10 batteries in the bank and numEnabled is 7. The first
	// battery to enable will be the maximum (and as a tiebreaker, earliest) of
	// the first 4, since at least 6 batteries must be left after that. Then the
	// next battery to enable will be the maximum of the first 5 that comes
	// after the index of the first chosen battery. Basically we maintain a
	// sliding window of the possible choices for the next battery to enable.
	for i := range numEnabled {
		maxJoltage := byte('0')
		start := lastEnabledIndex + 1
		end := len(allBatteries) - numEnabled + i + 1
		for j := start; j < end; j++ {
			joltage := allBatteries[j]
			if joltage > maxJoltage {
				maxJoltage = joltage
				lastEnabledIndex = j
			}
		}
		toEnable = append(toEnable, maxJoltage)
	}

	return util.MustParseInt(string(toEnable))
}

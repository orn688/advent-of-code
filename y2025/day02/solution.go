package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

func Part1(input string) (string, error) {
	invalidIDSum := 0
	for rng := range strings.SplitSeq(input, ",") {
		startS, endS, ok := strings.Cut(rng, "-")
		if !ok {
			return "", fmt.Errorf("invalid range: %q", rng)
		}
		start, end := util.MustParseInt(startS), util.MustParseInt(endS)
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			// A number with an odd number of digits can never be made of a set
			// of digits repeated twice.
			if len(s)%2 != 0 {
				continue
			}
			if s[:len(s)/2] == s[len(s)/2:] {
				invalidIDSum += i
			}
		}
	}
	return strconv.Itoa(invalidIDSum), nil
}

func Part2(input string) (string, error) {
	invalidIDSum := 0
	for rng := range strings.SplitSeq(input, ",") {
		startS, endS, ok := strings.Cut(rng, "-")
		if !ok {
			return "", fmt.Errorf("invalid range: %q", rng)
		}
		start, end := util.MustParseInt(startS), util.MustParseInt(endS)
		for i := start; i <= end; i++ {
			if invalidPart2(i) {
				invalidIDSum += i
			}
		}
	}
	return strconv.Itoa(invalidIDSum), nil
}

func invalidPart2(val int) bool {
	s := strconv.Itoa(val)
	for seqLength := 1; seqLength <= len(s)/2; seqLength++ {
		if len(s)%seqLength != 0 {
			continue
		}
		uniqueSeqs := map[string]struct{}{}
		for start := 0; start < len(s); start += seqLength {
			seq := s[start : start+seqLength]
			uniqueSeqs[seq] = struct{}{}
			if len(uniqueSeqs) > 1 {
				break
			}
		}
		if len(uniqueSeqs) == 1 {
			// "Invalid" number, it's one digit sequence repeated two or more
			// times.
			return true
		}
	}
	return false
}

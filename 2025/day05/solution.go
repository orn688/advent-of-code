package day05

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

type idRange struct {
	start int
	end   int
}

func Part1(input string) (string, error) {
	goodRanges, ids, err := parseInput(input)
	if err != nil {
		return "", err
	}

	mergedRanges := mergeRanges(goodRanges)

	var goodCount int
	for _, id := range ids {
		_, good := slices.BinarySearchFunc(mergedRanges, id, func(rng idRange, otherID int) int {
			if id < rng.start {
				return 1
			} else if id > rng.end {
				return -1
			}
			return 0
		})
		if good {
			goodCount++
		}
	}

	return strconv.Itoa(goodCount), nil
}

func Part2(input string) (string, error) {
	goodRanges, _, err := parseInput(input)
	if err != nil {
		return "", err
	}

	mergedRanges := mergeRanges(goodRanges)

	count := 0
	for _, rng := range mergedRanges {
		count += rng.end - rng.start + 1
	}
	return strconv.Itoa(count), nil
}

func parseInput(input string) ([]idRange, []int, error) {
	var goodRanges []idRange
	var ids []int
	foundBlank := false
	for line := range strings.SplitSeq(input, "\n") {
		if foundBlank {
			ids = append(ids, util.MustParseInt(line))
		} else if line == "" {
			foundBlank = true
		} else {
			start, end, ok := strings.Cut(line, "-")
			if !ok {
				return nil, nil, fmt.Errorf("bad line: %q", line)
			}
			goodRanges = append(goodRanges, idRange{
				start: util.MustParseInt(start),
				end:   util.MustParseInt(end),
			})
		}
	}
	return goodRanges, ids, nil
}

func mergeRanges(ranges []idRange) []idRange {
	slices.SortFunc(ranges, func(a, b idRange) int {
		if a.start != b.start {
			return cmp.Compare(a.start, b.start)
		}
		return cmp.Compare(a.end, b.end)
	})

	mergedRanges := ranges[:1]
	for _, rng := range ranges[1:] {
		prev := &mergedRanges[len(mergedRanges)-1]
		if rng.start <= prev.end {
			if rng.end > prev.end {
				prev.end = rng.end
			}
		} else {
			mergedRanges = append(mergedRanges, rng)
		}
	}
	return mergedRanges
}

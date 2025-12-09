package day08

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

type point [3]int

func Part1(input string) (string, error) {
	return solutionImpl(input, 1000)
}

func Part2(input string) (string, error) {
	return solutionImpl(input, 0)
}

func solutionImpl(input string, pairsToJoin int) (string, error) {
	isPart2 := pairsToJoin == 0

	points := parseInput(input)

	type pointPair struct {
		i, j     int
		distance float64
	}
	var pairs []pointPair
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, pointPair{
				i:        i,
				j:        j,
				distance: distance(points[i], points[j]),
			})
		}
	}

	slices.SortFunc(pairs, func(p1, p2 pointPair) int {
		return cmp.Compare(p1.distance, p2.distance)
	})

	circuitIDs := make(map[int]int)
	circuits := make(map[int]map[int]struct{})
	for i := range points {
		// Each point is initially only in its own circuit.
		circuitIDs[i] = i
		circuits[i] = map[int]struct{}{i: {}}
	}

	if isPart2 {
		pairsToJoin = len(pairs)
	}
	for _, pair := range pairs[:pairsToJoin] {
		// For simplicity, always merge into the circuit with the higher ID (the
		// child) into the circuit with the lower ID (the parent).
		parentID := circuitIDs[pair.i]
		childID := circuitIDs[pair.j]
		if parentID == childID {
			// Already part of the same circuit, no work required.
			continue
		}
		childCircuit := circuits[childID]
		for k := range childCircuit {
			circuitIDs[k] = parentID
		}
		maps.Insert(circuits[parentID], maps.All(childCircuit))
		// The merged circuit no longer exists.
		circuits[childID] = nil
		if isPart2 && len(circuits[parentID]) == len(points) {
			product := points[pair.i][0] * points[pair.j][0]
			return strconv.Itoa(product), nil
		}
	}
	if isPart2 {
		return "", fmt.Errorf("never ended up with just one circuit in part 2")
	}

	var sizes []int
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	product := 1
	for i := range 3 {
		product *= sizes[i]
	}

	return strconv.Itoa(product), nil
}

func distance(pt1, pt2 point) float64 {
	var sum float64
	for i := range len(pt1) {
		diff := pt1[i] - pt2[i]
		sum += math.Pow(float64(diff), 2)
	}
	return math.Sqrt(sum)
}

func parseInput(input string) []point {
	var result []point
	for line := range strings.SplitSeq(input, "\n") {
		var pt point
		for i, num := range strings.Split(line, ",") {
			pt[i] = util.MustParseInt(num)
		}
		result = append(result, pt)
	}
	return result
}

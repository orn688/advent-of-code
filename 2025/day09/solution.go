package day09

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

type point struct {
	x, y int
}

func Part1(input string) (string, error) {
	var points []point
	for line := range strings.SplitSeq(input, "\n") {
		x, y, ok := strings.Cut(line, ",")
		if !ok {
			return "", fmt.Errorf("invalid line: %q", line)
		}
		points = append(points, point{
			x: util.MustParseInt(x),
			y: util.MustParseInt(y),
		})
	}

	maxArea := 0
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			width := util.Abs(p1.x-p2.x) + 1
			height := util.Abs(p1.y-p2.y) + 1
			maxArea = max(maxArea, height*width)
		}
	}

	return strconv.Itoa(maxArea), nil
}

type line struct {
	location int
	start    int
	end      int
}

func Part2(input string) (string, error) {
	var points []point
	var maxX, maxY int
	for line := range strings.SplitSeq(input, "\n") {
		x, y, ok := strings.Cut(line, ",")
		if !ok {
			return "", fmt.Errorf("invalid line: %q", line)
		}
		pt := point{
			x: util.MustParseInt(x),
			y: util.MustParseInt(y),
		}
		points = append(points, pt)
		maxX = max(maxX, pt.x)
		maxY = max(maxY, pt.y)
	}

	var verticalLines []line
	var horizontalLines []line

	for i := range points {
		pt1 := points[i]
		pt2 := points[(i+1)%len(points)]
		if pt1.x == pt2.x {
			ys := []int{pt1.y, pt2.y}
			slices.Sort(ys)
			verticalLines = append(verticalLines, line{
				location: pt1.x,
				start:    ys[0],
				end:      ys[1],
			})
		} else {
			xs := []int{pt1.x, pt2.x}
			slices.Sort(xs)
			horizontalLines = append(horizontalLines, line{
				location: pt1.y,
				start:    xs[0],
				end:      xs[1],
			})
		}
	}

	sortFunc := func(l line) []int {
		return []int{l.location, l.start, l.end}
	}
	util.SortBy(verticalLines, sortFunc)
	util.SortBy(horizontalLines, sortFunc)

	lineIndex := func(lines []line, val int) int {
		idx, _ := slices.BinarySearchFunc(lines, val, func(l line, curr int) int {
			return cmp.Compare(l.location, curr)
		})
		return idx
	}

	maxArea := 0
	for i := range len(points) - 1 {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			top := min(p1.y, p2.y)
			bottom := max(p1.y, p2.y)
			left := min(p1.x, p2.x)
			right := max(p1.x, p2.x)

			var foundIntersecting bool

			vStart := lineIndex(verticalLines, left)
			vEnd := lineIndex(verticalLines, right)
			for i := vStart; i <= min(vEnd, len(verticalLines)-1); i++ {
				line := verticalLines[i]
				if line.location <= left || line.location >= right {
					continue
				}
				// Line is fully above or below the box.
				if (line.start <= top && line.end <= top) ||
					(line.start >= bottom && line.end >= bottom) {
					continue
				}
				foundIntersecting = true
				break
			}
			if foundIntersecting {
				continue
			}

			hStart := lineIndex(horizontalLines, top)
			hEnd := lineIndex(horizontalLines, bottom)
			for i := hStart; i <= min(hEnd, len(horizontalLines)-1); i++ {
				line := horizontalLines[i]
				if line.location <= top || line.location >= bottom {
					continue
				}
				// Line is fully to the left or right of the box.
				if (line.start <= left && line.end <= left) ||
					(line.start >= right && line.end >= right) {
					continue
				}
				foundIntersecting = true
				break
			}
			if foundIntersecting {
				continue
			}

			width := right - left + 1
			height := bottom - top + 1
			area := height * width
			maxArea = max(maxArea, area)
		}
	}

	return strconv.Itoa(maxArea), nil
}

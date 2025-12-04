package day04

import (
	"strconv"
	"strings"

	"github.com/orn688/advent-of-code/util"
)

const maxNeighbors = 4

type coordinate struct {
	x int
	y int
}

func Part1(input string) (string, error) {
	rows := util.Map(strings.Split(input, "\n"), func(s string) []rune {
		return []rune(s)
	})
	count := len(getAccessible(rows))
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	rows := util.Map(strings.Split(input, "\n"), func(s string) []rune {
		return []rune(s)
	})
	var total int
	for {
		accessible := getAccessible(rows)
		if len(accessible) == 0 {
			break
		}
		total += len(accessible)
		for _, coord := range accessible {
			rows[coord.y][coord.x] = '.'
		}
	}
	return strconv.Itoa(total), nil
}

func getAccessible(rows [][]rune) []coordinate {
	var res []coordinate
	for y, row := range rows {
		for x, char := range row {
			if char != '@' {
				continue
			}

			var neighborCount int
			for dy := -1; dy <= 1; dy++ {
				if y+dy < 0 || y+dy >= len(rows) {
					continue
				}
				for dx := -1; dx <= 1; dx++ {
					if x+dx < 0 || x+dx >= len(row) {
						continue
					}
					if rows[y+dy][x+dx] == '@' {
						neighborCount++
					}
				}
			}
			// Discount the current roll itself.
			neighborCount--

			if neighborCount < maxNeighbors {
				res = append(res, coordinate{x: x, y: y})
			}
		}
	}
	return res
}

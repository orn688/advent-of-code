package util

import (
	"log"
	"strconv"
)

func MustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("Invalid integer value %q", s)
	}
	return val
}

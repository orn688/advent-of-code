package util

import (
	"errors"
	"flag"
)

type Flags struct {
	Day int
	// Part must be 1 or 2.
	Part int
}

func ParseArgs() (*Flags, error) {
	var flags Flags
	flag.IntVar(&flags.Day, "day", 0, "")
	flag.IntVar(&flags.Part, "part", 1, "")
	flag.Parse()

	if len(flag.Args()) != 0 {
		return nil, errors.New("no positional args allowed")
	}

	if flags.Day < 1 || flags.Day > 25 {
		return nil, errors.New("-day must be between 1 and 25")
	}

	if flags.Part != 1 && flags.Part != 2 {
		return nil, errors.New("-part must be 1 or 2")
	}

	return &flags, nil
}

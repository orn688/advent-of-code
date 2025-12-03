package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/orn688/advent-of-code/util"
	"github.com/orn688/advent-of-code/y2025/day01"
	"github.com/orn688/advent-of-code/y2025/day02"
	"github.com/orn688/advent-of-code/y2025/day03"
)

var days = map[int][2]func(string) (string, error){
	1: {day01.Part1, day01.Part2},
	2: {day02.Part1, day02.Part2},
	3: {day03.Part1, day03.Part2},
}

func main() {
	if err := mainImpl(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func mainImpl(ctx context.Context) error {
	flags, err := util.ParseArgs()
	if err != nil {
		return err
	}

	daySolutions, ok := days[flags.Day]
	if !ok {
		return fmt.Errorf("no solution for day %d", flags.Day)
	}

	soln := daySolutions[flags.Part-1]
	if soln == nil {
		return fmt.Errorf("no solution for day %d part %d", flags.Day, flags.Part)
	}

	input, err := util.FetchInput(ctx, 2025, flags.Day)
	if err != nil {
		return err
	}

	startTime := time.Now()
	output, err := soln(strings.TrimRight(input, "\n"))
	if err != nil {
		return err
	}

	duration := time.Since(startTime)

	_, err = fmt.Println(output)
	fmt.Printf("-- completed in %s\n", duration)
	return err
}

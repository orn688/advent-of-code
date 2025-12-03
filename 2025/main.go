package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/orn688/advent-of-code/2025/day01"
	"github.com/orn688/advent-of-code/2025/day02"
	"github.com/orn688/advent-of-code/2025/day03"
	"github.com/orn688/advent-of-code/util"
)

const year = 2025

var days = map[int][2]func(string) (string, error){
	1: {day01.Part1, day01.Part2},
	2: {day02.Part1, day02.Part2},
	3: {day03.Part1, day03.Part2},
}

func init() {
	// Go doesn't support dynamically importing the package and function for the
	// given day and part, so we need to keep track of a mapping up above.
	//
	// Here we validate that the mapping is correct, i.e. that `dayX.PartY` is
	// mapped to key X and at index Y-1. This helps catch copy-pasta when
	// registering solutions for new days.
	var errs []error
	for day, solutions := range days {
		for i, soln := range solutions {
			part := i + 1
			// e.g. "github.com/foo/bar/pkg.FunctionName"
			absFuncName := runtime.FuncForPC(reflect.ValueOf(soln).Pointer()).Name()

			parts := strings.Split(absFuncName, "/")
			if parts[len(parts)-2] != strconv.Itoa(year) {
				errs = append(errs, fmt.Errorf("wrong year for day %d part %d", day, part))
			}
			// e.g. pkg.FunctionName"
			relFuncName := parts[len(parts)-1]

			pkg, funcName, ok := strings.Cut(relFuncName, ".")
			if !ok {
				log.Panicf("failed to get package and function names")
			}

			if pkg != fmt.Sprintf("day%02d", day) {
				errs = append(errs, fmt.Errorf("wrong package %q for day %d part %d", pkg, day, part))
			}

			if funcName != fmt.Sprintf("Part%d", part) {
				errs = append(errs, fmt.Errorf("wrong function %q for day %d part %d", funcName, day, part))
			}
		}
	}
	for _, err := range errs {
		fmt.Println("ERROR: " + err.Error())
	}
	if len(errs) > 0 {
		os.Exit(1)
	}
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

	input, err := util.FetchInput(ctx, year, flags.Day)
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

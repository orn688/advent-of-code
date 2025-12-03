package day01

import "testing"

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	got, err := Part1(testInput)
	if err != nil {
		t.Fatal(err)
	}

	want := "11"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	got, err := Part2(testInput)
	if err != nil {
		t.Fatal(err)
	}

	want := "31"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

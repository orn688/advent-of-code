package day02

import "testing"

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	got, err := Part1(testInput)
	if err != nil {
		t.Fatal(err)
	}

	want := "2"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

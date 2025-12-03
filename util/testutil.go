package util

import "testing"

func CheckAnswer(t *testing.T, impl func(string) (string, error), input, want string) {
	t.Helper()

	got, err := impl(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

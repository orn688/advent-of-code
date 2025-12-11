package util

import (
	"cmp"
	"slices"
)

func Map[T1, T2 any](in []T1, f func(T1) T2) []T2 {
	var out []T2
	for _, item := range in {
		out = append(out, f(item))
	}
	return out
}

func SortBy[T1 any, T2 cmp.Ordered](in []T1, f func(T1) []T2) {
	slices.SortFunc(in, func(a, b T1) int {
		af, bf := f(a), f(b)
		if len(af) != len(bf) {
			panic("differing lengths")
		}
		for i := range af {
			if c := cmp.Compare(af[i], bf[i]); c != 0 {
				return c
			}
		}
		return 0
	})
}

package util

func Map[T1, T2 any](in []T1, f func(T1) T2) []T2 {
	var out []T2
	for _, item := range in {
		out = append(out, f(item))
	}
	return out
}

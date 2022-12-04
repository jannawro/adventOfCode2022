package utils

type HandleDiff[T comparable] func(item1 T, item2 T) bool

func HandleDiffDefault[T comparable](val1, val2 T) bool {
	return val1 == val2
}

func Diff[T comparable](items1 []T, items2 []T) []T {
	var acc []T
	for _, item1 := range items1 {
		find := false
		for _, item2 := range items2 {
			if HandleDiffDefault(item1, item2) {
				find = true
				break
			}
		}
		if !find {
			acc = append(acc, item1)
		}
	}
	return acc
}

func SliceIntersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}
	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

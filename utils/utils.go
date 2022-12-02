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

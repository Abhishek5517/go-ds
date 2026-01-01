package sort

import "sort"

func Sort[T any](arr []T, less func(a, b T) bool) {
	sort.Slice(arr, func(i, j int) bool {
		return less(arr[i], arr[j])
	})
}

package pkg

import (
	"cmp"
	"log/slog"
	"math/rand/v2"
	"slices"
)

// ShuffleSlice randomizes the order of slice `s` via rand.Shuffle.
func ShuffleSlice[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// SortedInsert inserts t into ts, where t and ts are of cmp.Ordered type
func SortedInsert[T cmp.Ordered](ts []T, t T, logger *slog.Logger) []T {
	// find slot
	i, ok := slices.BinarySearch(ts, t)
	if !ok {
		logger.Debug("value not found in slice", slog.Any("value", t))
	}
	return slices.Insert(ts, i, t)
}

func DuplicateElements[T cmp.Ordered](ts []T) []T {
	var res []T
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			if ts[i] == ts[j] {
				res = append(res, ts[i])
			}
		}
	}
	return res
}

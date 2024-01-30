package medianTwoSortedArrays

import (
	"math"
	"sort"
)

func medianTwoSortedArrays(a1, a2 []int) float64 {
	a1 = append(a1, a2...)
	sort.Slice(a1, func(i, j int) bool {
		return a1[i] < a1[j]
	})

	mid := float64(len(a1)/2)

	if (isNotInt(mid)){
		return float64(a1[int(math.Floor(float64(mid)))])
	}
	return float64(a1[int(mid-1)] + a1[int(mid)])/2
}

func isNotInt(val float64) bool {
	return val == float64(int(val))
}
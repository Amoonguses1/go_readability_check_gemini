package main

import (
	"math"
	"sort"
)

// Returns the minimum number of magic beans that you have to remove
// so that each bag such that the number of beans in each remaining non-empty bag
// (still containing at least one bean) is equal.
func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var total int
	for _, bean := range beans {
		total += bean
	}

	removedBeansNum := math.MaxInt
	for i, bean := range beans {
		// you choose the i - 1 bags of beans empty,
		// we should choose i bags with the least amoount of beans.
		// That is, you remove beans from i+1 or later bags
		// until the number of beans in the bags equals to that of i bags.
		removedBeansNum = min(removedBeansNum, total-(len(beans)-i)*bean)
	}

	return int64(removedBeansNum)
}

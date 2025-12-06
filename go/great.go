package main

// Find a 2D array answer of size n where answer[i] = [mini, maxi]:
//
// mini is the largest value in the tree that is smaller than or equal to queries[i].
// If a such value does not exist, add -1 instead.
//
// maxi is the smallest value in the tree that is greater than or equal to queries[i].
// If a such value does not exist, add -1 instead.
func closestNodes(root *TreeNode, queries []int) [][]int {
	sorted := convertToSlice(root)
	res := [][]int{}
	for _, query := range queries {
		res = append(res, findClosestValues(sorted, query))
	}
	return res
}

// Converts binary tree to sorted slice.
func convertToSlice(root *TreeNode) []int {
	arr := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		arr = append(arr, node.Val)

		current = node.Right
	}

	return arr
}

// findClosestValues returns the minimum and maximum values in sortedArr
// that are closest to the given query with binary search.
//
// The result is a two-element slice: [minClosest, maxClosest].
// If a value is not found, its corresponding element in the slice is -1.
// If an exact match is found, the function returns [query, query].
func findClosestValues(sortedArr []int, query int) []int {
	left, right := -1, len(sortedArr)
	res := []int{-1, -1}
	for right-left > 1 {
		mid := (left + right) / 2
		if sortedArr[mid] == query {
			return []int{query, query}
		} else if sortedArr[mid] > query {
			right = mid
			res[1] = sortedArr[mid]
		} else {
			left = mid
			res[0] = sortedArr[mid]
		}
	}
	return res
}

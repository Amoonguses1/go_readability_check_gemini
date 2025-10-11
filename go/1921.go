package main

import "sort"

// eliminateMaximum returns the maximum number of monsters that can be eliminated
// before any of them reach the city.
//
// Each monster starts at a distance given by dist[i] and moves toward the city at speed[i].
// A player can eliminate at most one monster per minute, starting from minute 0.
// A monster reaches the city at time ceil(dist[i] / speed[i]).
// The function returns the maximum number of monsters the player can eliminate before any monster arrives.
func eliminateMaximum(dist []int, speed []int) int {
	arrivalTimes := make([]int, len(dist))
	for i := 0; i < len(dist); i++ {
		arrivalTimes[i] = (dist[i] + speed[i] - 1) / speed[i]
	}
	sort.Ints(arrivalTimes)

	eliminated := 0
	for curMinute, arrival := range arrivalTimes {
		// If any monster arrives before or at the current minute, the game ends.
		if arrival <= curMinute {
			break
		}
		eliminated++
	}

	return eliminated
}

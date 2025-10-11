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
	for i, d := range dist {
		// Calculate arrival time: ceil(distance / speed) using integer arithmetic.
		arrivalTimes[i] = (d + speed[i] - 1) / speed[i]
	}
	sort.Ints(arrivalTimes)

	for curMinute, arrival := range arrivalTimes {
		// If a monster arrives at or before we can eliminate it, the game ends.
		if arrival <= curMinute {
			return curMinute
		}
	}
	// If the loop completes, we can eliminate all monsters.
	return len(dist)
}

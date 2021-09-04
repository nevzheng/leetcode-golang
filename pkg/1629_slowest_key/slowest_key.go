package _629_slowest_key

func AdjacentDifference(releaseTimes []int) []int {
	ad := make([]int, len(releaseTimes))
	for i, time := range releaseTimes {
		if i == 0 {
			ad[i] = time
		} else {
			ad[i] = time - releaseTimes[i-1]
		}
	}
	return ad
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	durations := AdjacentDifference(releaseTimes)
	idx := 0
	mx := keysPressed[0]
	for i, c := range keysPressed {
		if durations[i] > durations[idx] ||
			(durations[i] == durations[idx] && keysPressed[i] > keysPressed[idx]) {
			idx = i
			mx = uint8(c)
		}
	}
	return mx
}

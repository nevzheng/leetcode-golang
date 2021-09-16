package _042_trapping_rain_water

func trap(H []int) int {
	l, r := 0, len(H)-1
	water := 0
	lMax, rMax := 0, 0
	for l < r {
		if H[l] < H[r] {
			if H[l] >= lMax {
				lMax = H[l]
			} else {
				water += lMax - H[l]
			}
			l++
		} else {
			if H[r] >= rMax {
				rMax = H[r]
			} else {
				water += rMax - H[r]
			}
			r--
		}
	}
	return water
}

func trap(height []int) int {
	i, j := 0, len(height)-1
	water := 0
	lmax, rmax := height[i], height[j]

	for i < j {
		if lmax < rmax {
			i += 1
			lmax = max(lmax, height[i])
			water += lmax - height[i]
		} else {
			j -= 1
			rmax = max(rmax, height[j])
			water += rmax - height[j]
		}
	}

	return water
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
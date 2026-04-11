func longestConsecutive(nums []int) int {
	se := make(map[int]bool)

	for _, num := range nums {
		se[num] = true
	}

	ans := 0

	for _, num := range nums {
		if se[num-1] {
			continue
		}

		count := 0
		ele := num
		for {
			if se[ele] {
				count++
				ele++
			} else {
				if count > ans { ans = count }
				break
			}
		}
	}

	return ans
}

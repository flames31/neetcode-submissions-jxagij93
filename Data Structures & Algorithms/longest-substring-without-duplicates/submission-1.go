func lengthOfLongestSubstring(s string) int {
	se := make(map[byte]bool)
	ans := 0
	i, j := 0, 0

	for j < len(s) {
		if !se[s[j]] {
			se[s[j]] = true
			if (j-i+1) > ans { ans = (j-i+1) }
			j++
		} else {
			for i < j && se[s[j]] {
				se[s[i]] = false
				i++
			} 
		}
	}

	return ans
}

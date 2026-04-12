class Solution:
	def maxArea(self, heights: List[int]) -> int:
		i, j = 0, len(heights)-1
		maxWater = 0

		while i < j:
			water = (j-i) * min(heights[i], heights[j])
			maxWater = max(maxWater, water)

			if heights[i] > heights[j]:
				j -= 1
			else:
				i += 1
		
		return maxWater
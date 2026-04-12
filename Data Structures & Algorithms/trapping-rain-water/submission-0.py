class Solution:
	def trap(self, height: List[int]) -> int:
		i, j = 0, len(height)-1
		lmax, rmax = height[i], height[j]
		water = 0

		while i < j:
			if lmax < rmax:
				i += 1
				lmax = max(lmax, height[i])
				water += lmax - height[i]
			else:
				j -= 1
				rmax = max(rmax, height[j])
				water += rmax - height[j]
		
		return water
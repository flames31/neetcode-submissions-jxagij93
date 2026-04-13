class Solution:
	def maxProfit(self, prices: List[int]) -> int:
		least = float("inf")
		maxP = 0

		for p in prices:
			least = min(least, p)
			maxP = max(maxP, p-least)
		
		return maxP
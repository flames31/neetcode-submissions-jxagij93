class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        elem = {}

        for idx, num in enumerate(nums):
            if target - num in elem:
                return [elem[target-num], idx]
            
            elem[num] = idx
        
        return [0,0]
class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        elem = set()

        for num in nums:
            if num in elem:
                return True
            
            elem.add(num)

        return False
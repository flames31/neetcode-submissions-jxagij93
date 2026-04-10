class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        sortedWords = {}
        ans = []

        for s in strs:
            sw = ''.join(sorted(s))
            if sw not in sortedWords:
                sortedWords[sw] = []
            
            sortedWords[sw].append(s)
        
        for s in sortedWords:
            ans.append(sortedWords[s])
        
        return ans
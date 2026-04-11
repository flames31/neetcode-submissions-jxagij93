class Solution:

	def encode(self, strs: List[str]) -> str:
		s = ""

		for st in strs:
			s += str(len(st)) + "#" + st
		
		return s

	def decode(self, s: str) -> List[str]:
		strs = []
		i = 0

		while i < len(s):
			word = ""
			le = ""
			while s[i] != '#':
				le += s[i]
				i += 1
			
			le = int(le)

			strs.append(s[i+1:i+le+1])
			i += le + 1 

		return strs

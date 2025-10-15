def lengthOfLongestSubstring(self, s: str) -> int:
	begin = 0
	end = 0
	window_state = set()
	res = 0

	while end < len(s):
		window_size = end - begin + 1
		char = s[end]
		if char in window_state:
			while char in window_state:
				window_state.remove(s[begin])
				begin += 1
		else:
			res = max(res, window_size)
		window_state.add(char)
		end += 1

	return res
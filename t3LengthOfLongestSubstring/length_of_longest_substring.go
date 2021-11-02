package t3LengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	mp := make(map[int32]int)
	length := 0
	start := 0
	for index, char := range s {
		if idx, ok := mp[char]; ok {
			if idx > start {
				start = idx
			}
		}
		if index-start+1 > length {
			length = index - start + 1
		}
		mp[char] = index
	}
	return length
}

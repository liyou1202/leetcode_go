package arr

// LongestCommonPrefix
// https://redirect.liyou-chen.com/cr2li5
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) <= 1 {
		return strs[0]
	}

	shortestStr := findShortestStr(strs)
	if len(shortestStr) == 0 {
		return ""
	}

	for i := 0; i < len(shortestStr); i++ {
		for _, val := range strs {
			if val[i] != shortestStr[i] {
				return shortestStr[:i]
			}
		}
		if i == len(shortestStr)-1 {
			return shortestStr
		}
	}

	return ""
}

func findShortestStr(arr []string) string {
	shortestIndex := 0
	shortestLength := len(arr[0])
	for i, val := range arr {
		if len(val) <= shortestLength {
			shortestLength = len(val)
			shortestIndex = i
		}
	}
	return arr[shortestIndex]
}

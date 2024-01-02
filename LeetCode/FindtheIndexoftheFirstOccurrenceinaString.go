func strStr(haystack string, needle string) int {
	index := strings.Index(haystack, needle)

	if index != -1 {
		return index
	} else {
		return -1
	}
}
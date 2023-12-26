func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	palindrome, _ := strconv.Atoi(string(runes))
	if palindrome == x {
		fmt.Println(palindrome)
		return true
	}
	return false
}
func removeDuplicates(nums []int) int {
	encountered := make(map[int]bool)
	uniqueNums := make([]int, 0)

	for _, value := range nums {
		if !encountered[value] {
			uniqueNums = append(uniqueNums, value)
			encountered[value] = true
		}
	}
	copy(nums, uniqueNums)
	return len(uniqueNums)
}
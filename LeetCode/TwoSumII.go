func twoSum(numbers []int, target int) []int {
	var arraySum []int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				arraySum = append(arraySum, i+1)
				arraySum = append(arraySum, j+1)
			}
		}
	}
	return arraySum
}
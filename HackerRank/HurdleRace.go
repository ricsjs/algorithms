func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var maxHeight int32 = 0
	for i := 0; i < len(height); i++ {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}

	if maxHeight > k {
		return maxHeight - k
	}

	return 0

}
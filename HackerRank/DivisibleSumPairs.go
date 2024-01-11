func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var countPairs int32 = 0
	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			if (ar[i]+ar[j])%k == 0 {
				countPairs += 1
			}
		}
	}

	return countPairs
}
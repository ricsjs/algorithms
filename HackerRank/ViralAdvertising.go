func viralAdvertising(n int32) int32 {
	// Write your code here
	shared, liked, totalLikes := 5, 0, 0

	for i := 0; i < int(n); i++ {
		liked = shared / 2
		totalLikes += liked
		shared = liked * 3
	}
	return int32(totalLikes)
}
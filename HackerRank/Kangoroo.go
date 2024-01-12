func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	for i := 0; i < 10000; i++ {
		if x1 == x2 {
			return "YES"
		}

		x1 += v1
		x2 += v2
	}
	return "NO"
}
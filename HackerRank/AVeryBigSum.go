package main

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	sum := 0
	for i := 0; i < len(ar); i++ {
		sum += int(ar[i])
	}
	return int64(sum)

}

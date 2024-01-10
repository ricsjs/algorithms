package main

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sumAr int32 = 0

	for i := 0; i < len(ar); i++ {
		sumAr += ar[i]
	}

	return sumAr
}

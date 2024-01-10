package main

import (
	"fmt"
)

/*
 * Complete the 'circularArrayRotation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 *  3. INTEGER_ARRAY queries
 */

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	n := int32(len(a))
	fmt.Println(n)

	k = k % n

	rotatedArray := make([]int32, n)
	for i := int32(0); i < n; i++ {
		newPosition := (i + k) % n
		rotatedArray[newPosition] = a[i]
	}

	results := make([]int32, len(queries))
	for i, query := range queries {
		results[i] = rotatedArray[query]
	}

	return results
}

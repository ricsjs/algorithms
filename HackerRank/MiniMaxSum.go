package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var intArr []int
	for _, v := range arr {
		intArr = append(intArr, int(v))
	}

	sort.Ints(intArr)

	minSum := intArr[0] + intArr[1] + intArr[2] + intArr[3]

	maxSum := intArr[1] + intArr[2] + intArr[3] + intArr[4]

	fmt.Println(minSum, maxSum)
}

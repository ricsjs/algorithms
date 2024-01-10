package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	positiveRatios := 0.0
	negativeRatios := 0.0
	zeroRatios := 0.0

	nPositive := []int{}
	nNegative := []int{}
	nZero := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			nPositive = append(nPositive, i)
		} else if arr[i] < 0 {
			nNegative = append(nNegative, i)
		} else if arr[i] == 0 {
			nZero = append(nZero, i)
		}
	}

	positiveRatios = float64(len(nPositive)) / float64(len(arr))
	negativeRatios = float64(len(nNegative)) / float64(len(arr))
	zeroRatios = float64(len(nZero)) / float64(len(arr))

	fmt.Printf("%.6f\n%.6f\n%.6f\n", positiveRatios, negativeRatios, zeroRatios)

}

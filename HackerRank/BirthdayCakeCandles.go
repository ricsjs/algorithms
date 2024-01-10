package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var maxHeightNumber int32
	var countNumber int32

	for _, i := range candles {
		if i > maxHeightNumber {
			maxHeightNumber = i
			countNumber = 1
		} else if i == maxHeightNumber {
			countNumber += 1
		}
	}

	return countNumber
}

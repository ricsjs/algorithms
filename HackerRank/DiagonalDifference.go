package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	n := len(arr)
	var primaryDiagonalSum int32 = 0
	var secondaryDiagonalSum int32 = 0

	for i := 0; i < len(arr); i++ {
		primaryDiagonalSum += arr[i][i]
		secondaryDiagonalSum += arr[i][n-1-i]
	}

	absoluteDifference := Abs(primaryDiagonalSum - secondaryDiagonalSum)

	return absoluteDifference

}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}

	return x
}

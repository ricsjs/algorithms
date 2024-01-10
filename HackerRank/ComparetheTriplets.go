package main

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {

	points := []int32{0, 0}

	for i := 0; i < 3; i++ {
		if a[i] < b[i] {
			points[1] += 1
		} else if a[i] > b[i] {
			points[0] += 1
		}
	}

	return points
}

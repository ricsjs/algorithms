package main

import (
	"fmt"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for k := int32(0); k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

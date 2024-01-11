package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	finalGrade := []int32{}

	for i := 0; i < len(grades); i++ {
		if grades[i]%5 == 0 || grades[i] < 38 {
			finalGrade = append(finalGrade, grades[i])
		} else {
			nextMultipleOf5 := grades[i] + (5 - grades[i]%5)
			if nextMultipleOf5-grades[i] < 3 {
				finalGrade = append(finalGrade, nextMultipleOf5)
			} else {
				finalGrade = append(finalGrade, grades[i])
			}
		}
	}

	return finalGrade
}

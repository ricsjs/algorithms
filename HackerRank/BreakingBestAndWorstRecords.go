func breakingRecords(scores []int32) []int32 {
	var countMax int32 = 0
	var countMin int32 = 0
	var maxRecord int32 = scores[0]
	var minRecord int32 = scores[0]
	arrScores := []int32{}

	for i := 1; i < len(scores); i++ {
		if scores[i] > maxRecord {
			maxRecord = scores[i]
			countMax += 1
		}
		if scores[i] < minRecord {
			minRecord = scores[i]
			countMin += 1
		}
	}

	arrScores = append(arrScores, countMax)
	arrScores = append(arrScores, countMin)

	return arrScores

}
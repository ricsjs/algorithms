var romanToNumber = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sSlice := strings.Split(s, "")
	var result int

	for i := 0; i < len(sSlice); i++ {
		value := romanToNumber[sSlice[i]]

		if i+1 < len(sSlice) {
			nextValue := romanToNumber[sSlice[i+1]]
			if nextValue > value {
				result += nextValue - value
				i++
				continue
			}
		}

		result += value
	}

	return result
}
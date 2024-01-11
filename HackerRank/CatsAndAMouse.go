func catAndMouse(x int32, y int32, z int32) string {
	strToReturn := ""

	distA := Abs(z - x)
	distB := Abs(z - y)

	if distA < distB {
		strToReturn = "Cat A"
	} else if distB < distA {
		strToReturn = "Cat B"
	} else {
		strToReturn = "Mouse C"
	}

	return strToReturn
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}

	return x
}
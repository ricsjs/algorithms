func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var billArr []int32
	sumBill := 0
	for i := 0; i < len(bill); i++ {
		if int32(i) != k {
			billArr = append(billArr, bill[i])
		}
	}
	for j := 0; j < len(billArr); j++ {
		sumBill += int(billArr[j])
	}

	var billAnna int32 = int32(sumBill) / 2

	if b > int32(billAnna) {
		billAnna = b - billAnna
		fmt.Println(billAnna)
	} else {
		fmt.Println("Bon Appetit")
	}

}
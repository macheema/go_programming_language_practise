package exercise

//Exercise43 Re write reverse to use an array pointer instead of a slice
func Exercise43(arrPtr *[10]int) {
	for i, j := 0, len(arrPtr)-1; i < j; i, j = i+1, j-1 {
		arrPtr[i], arrPtr[j] = arrPtr[j], arrPtr[i]
	}
}

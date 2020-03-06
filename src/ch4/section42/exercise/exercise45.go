package exercise

//Exercise45 Write an in-place function to eliminate adjacent duplicates in a []string slice
func Exercise45(data string) string {
	dataRunes := []rune(data)
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			dataRunes = append(dataRunes[:i], dataRunes[i+1:]...)[:len(dataRunes)-1]
			data = string(dataRunes)
			i--
		}
	}
	return string(dataRunes)
}

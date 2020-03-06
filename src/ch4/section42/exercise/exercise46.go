package exercise

import "unicode"

//Exercise46 Write an in-place function that squashes each run of adjacent Unico de spaces
//(see unicode.IsSpace) in a UTF-8-enco ded []byte slice into a single ASCII space.
func Exercise46(data []byte) string {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] && unicode.IsSpace(rune(data[i])) {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}
	return string(data)
}

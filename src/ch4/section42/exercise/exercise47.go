package exercise

import "unicode/utf8"

//Exercise47 Modif y reverse to reverse the characters of a []byte slice that represents a
//UTF-8-encoded string, in place. Can you do it without allocating new memory?
func Exercise47(data []byte) {

	for i := 0; i < len(data); {
		_, size := utf8.DecodeRune(data[i:])
		reverseBytes(data[i : i+size])
		i += size
	}
	reverseBytes(data)
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// func rev(b []byte) {
// 	size := len(b)
// 	for i := 0; i < len(b)/2; i++ {
// 		b[i], b[size-1-i] = b[size-1-i], b[i]
// 	}
// }

// // Exercise47 Reverse all the runes, and then the entire slice. The runes' bytes end up in
// // the right order.
// func Exercise47(b []byte) []byte {
// 	for i := 0; i < len(b); {
// 		_, size := utf8.DecodeRune(b[i:])
// 		rev(b[i : i+size])
// 		i += size
// 	}
// 	rev(b)
// 	return b
// }

//Exercise47 ...

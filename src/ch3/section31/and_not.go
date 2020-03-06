package section31

//BitAndNot ...
// 100
// 110
// 000
func BitAndNot(number uint64, number2 uint64) uint64 {
	return number &^ number2
}

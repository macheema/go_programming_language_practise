package exercise

//Exercise44 Write a version of rotate that operates in a single pass
func Exercise44(slice *[]int, rotateBy int) {
	(*slice) = append(*slice, (*slice)[:rotateBy]...)[rotateBy:]
}

package section42

import "fmt"

//SlicesConcept1 Slices is not assignable to Array
func SlicesConcept1() {
	arr := [2]int{1, 2}
	slice := []int{1, 2, 3}
	// arr = slice[0:2] illegal operation slice cannot assign to declaed arr variable
	slice2 := slice[0:2]
	fmt.Printf("%v\t%v\t%v", arr, slice, slice2)
}

//SlicesConcept2 Nil Slices vs Non Nill Zero Valued Slices
func SlicesConcept2() {
	var slice1 []int
	var slice2 []int = nil
	slice3 := []int(nil)
	slice4 := []int{}
	slice5 := []int{}
	slice6 := []int{}
	fmt.Printf("Slice1 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice1, cap(slice1), len(slice1), &slice1, slice1, slice1 == nil)
	fmt.Printf("Slice2 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice2, cap(slice2), len(slice2), &slice2, slice2, slice2 == nil)
	fmt.Printf("Slice3 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice3, cap(slice3), len(slice3), &slice3, slice3, slice3 == nil)
	fmt.Printf("Slice4 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice4, cap(slice4), len(slice4), &slice4, slice4, slice4 == nil)
	fmt.Printf("Slice5 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice5, cap(slice5), len(slice5), &slice5, slice5, slice5 == nil)
	fmt.Printf("Slice6 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice6, cap(slice6), len(slice6), &slice6, slice6, slice6 == nil)
	slice4 = append(slice4, 1)
	fmt.Printf("Slice4 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice4, cap(slice4), len(slice4), &slice4, slice4, slice4 == nil)
	fmt.Printf("Slice5 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice5, cap(slice5), len(slice5), &slice5, slice5, slice5 == nil)
	fmt.Printf("Slice6 -> [value : %v] [cap : %v] [len : %v] [addr : %p] [ArrayAddr : %p] [Nil : %v]\n", slice6, cap(slice6), len(slice6), &slice6, slice6, slice6 == nil)
}

//AppendToSlices ...
func AppendToSlices(iterations uint) {
	slice := make([]uint, 0)[:0]
	fmt.Printf("%d\t%d\t%v\t%p\t%p\n", len(slice), cap(slice), slice, slice, &slice)
	for i := uint(0); i < iterations; i++ {
		slice = append(slice, i)
		fmt.Printf("%d\t%d\t%v\t%p\t%p\n", len(slice), cap(slice), slice, slice, &slice)
	}
}

//AddressOfNilValue ...
func AddressOfNilValue() {
	a := []int{}
	fmt.Printf("%p\t%p\n", &a, a)
	arr := []int{1, 2, 3}
	fmt.Printf("%p\t%p\n", &arr, arr)
	arr1 := arr[0:0]
	fmt.Printf("%p\t%p\n", &arr1, arr1)
}

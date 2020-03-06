package section43

import "fmt"

//MapExist ...
func MapExist() {
	flag := false
	flag = equal(map[string]int{"A": 0}, map[string]int{"B": 66})
	fmt.Println(flag)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if xv != y[k] {
			return false
		}
		// if yv, ok := y[k]; !ok || yv != xv {
		// 	return false
		// }
	}
	return true
}

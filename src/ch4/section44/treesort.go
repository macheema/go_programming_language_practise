package section44

type tree struct {
	value int
	left  *tree
	right *tree
}

//Sort ...
func Sort(values []int) []int {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	sortedArray := []int{}
	sortedArray = appendValues(sortedArray, root)
	return sortedArray
}

func appendValues(values []int, root *tree) []int {
	if root == nil {
		return values
	}
	values = appendValues(values, root.left)
	values = append(values, root.value)
	values = appendValues(values, root.right)
	return values
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}
	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

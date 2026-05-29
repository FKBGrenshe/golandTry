package binaryTree

type tree struct {
	value      int
	leftChild  *tree
	rightChild *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	values = appendValues(values[:0], root)
}

func add(curNode *tree, v int) *tree {
	if curNode == nil {
		return &tree{value: v}
	}
	if v <= curNode.value {
		curNode.leftChild = add(curNode.leftChild, v)
	} else {
		curNode.rightChild = add(curNode.rightChild, v)
	}
	return curNode
}

func appendValues(arr []int, curNode *tree) []int {
	if curNode == nil {
		return arr
	}
	arr = appendValues(arr, curNode.leftChild)
	arr = append(arr, curNode.value)
	arr = appendValues(arr, curNode.rightChild)
	return arr
}

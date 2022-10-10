package main

import "fmt"

func main() {

	//slice := []int{1, 2, 34, 7, 27, 65, 77, 22, 46, 3}
	//
	//var x int
	//
	//for i := range slice {
	//	x += slice[i]
	//}
	//average := x / len(slice)
	//
	//difs := make([]int, len(slice))
	//var minNum int
	//
	//for i := range slice {
	//	dif := slice[i] - average
	//	difs[i] = dif
	//	if dif < 0 {
	//		difs[i] = dif * (-1)
	//	}
	//	if difs[minNum] >= difs[i] {
	//		minNum = i
	//	}
	//}

	//fmt.Println(minNum)
	//fmt.Println(slice[minNum])

	testSlice := []int{3, 4, 5, 8, 19}

	rootElement := node{
		9,
		nil,
		nil,
	}

	for i := range testSlice {
		addToTree(&rootElement, testSlice[i])
	}
}

type node struct {
	val   int
	left  *node
	right *node
}

func searchElement(root *node, searchVal int) bool {

	if root == nil {
		fmt.Println("Ck")
		return false
	}

	//if searchVal != root.val {
	//	if searchVal < root.val{
	//		searchElement(root.left, searchVal)
	//	} else {
	//		searchElement(root.right, searchVal)
	//	}
	//}
	//
	//if searchVal < root.val {
	//	if searchVal != root.val {
	//		searchElement(root.left, searchVal)
	//	}
	//}
	//
	//if searchVal < root.val {
	//	searchElement(root.left, searchVal)
	//} else {
	//	searchElement(root.right, searchVal)
	//}

	if searchVal == root.val {
		return true
	}

	if searchVal < root.val {
		if root.left.val == searchVal {
			return true
		} else {
			searchElement(root.left, searchVal)
		}
	} else {
		if root.right.val == searchVal {
			return true
		} else {
			searchElement(root.right, searchVal)
		}
	}

	return false
}

func addToTree(root *node, addVal int) {
	if root == nil {
		fmt.Println("Ck")
		return
	}

	if addVal < root.val {
		if root.left == nil {
			root.left = &node{addVal, nil, nil}
		} else {
			addToTree(root.left, addVal)
		}
	} else {
		if root.right == nil {
			root.right = &node{addVal, nil, nil}
		} else {
			addToTree(root.right, addVal)
		}
	}
}

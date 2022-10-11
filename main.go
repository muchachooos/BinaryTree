package main

import "fmt"

func main() {

	//slice := []int{1, 2, 34, 7, 27, 65, 77, 22, 46, 3}

	testSlice := []int{3, 4, 5, 8, 19}

	rootElement := node{
		9,
		nil,
		nil,
	}

	for i := range testSlice {
		addElement(&rootElement, testSlice[i])
	}

	fmt.Println(searchElement(&rootElement, 8))
	fmt.Println("1-----------")
	deleteElement(&rootElement, 8)
	fmt.Println("2----------")
	fmt.Println(searchElement(&rootElement, 8))
}

type node struct {
	val   int
	left  *node
	right *node
}

func sortSlice() {
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
}

func deleteElement(root *node, deleteVal int) {
	if root == nil {
		fmt.Println("Ck")
		return
	}

	if deleteVal == root.val {
		root = nil
		return
	}

	if deleteVal < root.val {
		deleteElement(root.left, deleteVal)
		if root.left.val == deleteVal {
			root.left = nil
			return
		}
	} else {
		deleteElement(root.right, deleteVal)
		if root.right.val == deleteVal {
			root.right = nil
			return
		}
	}
}

func searchElement(root *node, searchVal int) bool {
	if root == nil {
		fmt.Println("Ck")
		return false
	}

	if searchVal == root.val {
		return true
	}

	if searchVal < root.val {
		return searchElement(root.left, searchVal)
	} else {
		return searchElement(root.right, searchVal)
	}
}

func addElement(root *node, addVal int) {
	if root == nil {
		fmt.Println("Ck")
		return
	}

	if addVal < root.val {
		if root.left == nil {
			root.left = &node{addVal, nil, nil}
		} else {
			addElement(root.left, addVal)
		}
	} else {
		if root.right == nil {
			root.right = &node{addVal, nil, nil}
		} else {
			addElement(root.right, addVal)
		}
	}
}

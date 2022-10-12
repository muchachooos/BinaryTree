package main

import "fmt"

func main() {

	slice := []int{1, 27, 34, 7, 2, 65, 77, 22, 46, 3}

	sortSlice(slice)

}

type node struct {
	val   int
	left  *node
	right *node
}

func sortSlice(slice []int) {

	var x int

	for i := range slice {
		x += slice[i]
	}
	average := x / len(slice)

	diffs := make([]int, len(slice))
	var rootValue int

	for i := range slice {
		diffsVal := slice[i] - average
		diffs[i] = diffsVal
		if diffsVal < 0 {
			diffs[i] = diffsVal * (-1)
		}
		if diffs[rootValue] >= diffs[i] {
			rootValue = i
		}
	}

	saveVal := slice[rootValue]

	fmt.Println("pos rootValue: ", rootValue)
	fmt.Println("val rootValue: ", slice[rootValue])
	fmt.Println("del rootValue: ", slice)
	
	slice = append(slice[:rootValue], slice[rootValue+1:]...)

	fmt.Println("--------")
	fmt.Println("pos: ", rootValue)
	fmt.Println("saveVal: ", saveVal)
	fmt.Println("del rootValue: ", slice)

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

package main

import "fmt"

func main() {

	slice := []int{1, 27, 34, 7, 2, 65, 77, 22, 46, 3}

	rootValue := sortSlice(slice)

	rootElement := node{
		val:   rootValue,
		left:  nil,
		right: nil,
	}

	binaryTree := newBinaryTree(&rootElement)

	for i := range slice {
		addElement(binaryTree.rootElement, slice[i])
	}

	//fmt.Println("add-------")
	//fmt.Println("search1-------")
	//fmt.Println(searchElement(&rootElement, 22))
	//fmt.Println("delete-------")
	//deleteElement(&
	//rootElement, 999)
	//fmt.Println("search2-------")
	//fmt.Println(searchElement(&rootElement, 22))
}

type binaryTree struct {
	rootElement  *node
	countOfNodes int
	maxVal       int
	minVal       int
}

// Конструктор
func newBinaryTree(root *node) binaryTree {
	return binaryTree{
		rootElement: root,
	}
}

type node struct {
	val   int
	left  *node
	right *node
}

func (binT binaryTree) add(addVal int) {
	addElement(binT.rootElement, addVal)
}

func sortSlice(FSlice []int) int {

	var x int

	for i := range FSlice {
		x += FSlice[i]
	}
	average := x / len(FSlice)

	diffs := make([]int, len(FSlice))
	var number int

	for i := range FSlice {
		diffsVal := FSlice[i] - average
		diffs[i] = diffsVal
		if diffsVal < 0 {
			diffs[i] = diffsVal * (-1)
		}
		if diffs[number] >= diffs[i] {
			number = i
		}
	}

	FSlice = append(FSlice[:number], FSlice[number+1:]...)

	return FSlice[number]
}

func addElement(root *node, addVal int) {
	if root == nil {
		fmt.Println("nilAdd")
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

func searchElement(root *node, searchVal int) bool {
	if root == nil {
		fmt.Println("nilSearch")
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

func deleteElement(root *node, deleteVal int) {
	if root == nil {
		fmt.Println("nilDel")
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

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

	el1 := node{
		9,
		nil,
		nil,
	}

	testSlice := []int{3, 4, 5, 8, 19}

	for i := range testSlice {
		addToTree(&el1, testSlice[i])
	}

	fmt.Println("9", el1.val)
	fmt.Println("3", el1.left.val)
	fmt.Println("4", el1.left.right.val)
	fmt.Println("5", el1.left.right.right.val)
	fmt.Println("8", el1.left.right.right.right.val)
	fmt.Println("19", el1.right.val)
}

type node struct {
	val   int
	left  *node
	right *node
}

func addToTree(root *node, addVal int) {
	if root == nil {
		fmt.Println("Ck!")
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

//root.left.left.val = addEl

//if addVal < root.left.val {
//root.left.left = &node{addVal, nil, nil}
//} else {
//root.left.right = &node{addVal, nil, nil}
//}
//--------------------------------------
//if addVal > root.right.val {
//root.right.right = &node{addVal, nil, nil}
//} else {
//root.right.left = &node{addVal, nil, nil}
//}

package main

import "fmt"

func main() {

	slice := []int{1, 2, 34, 7, 27, 65, 77, 22, 46, 3}

	var x int

	for i := range slice {
		x += slice[i]
	}
	average := x / len(slice)

	difs := make([]int, len(slice))
	var minNum int

	for i := range slice {
		dif := slice[i] - average
		difs[i] = dif
		if dif < 0 {
			difs[i] = dif * (-1)
		}
		if difs[minNum] >= difs[i] {
			minNum = i
		}
	}

	//fmt.Println(minNum)
	//fmt.Println(slice[minNum])

	el1 := node{
		21,
		nil,
		nil,
	}
	addToTree(&el1, 33)
}

type node struct {
	val   int
	left  *node
	right *node
}

func addToTree(root *node, addEl int) {
	if root == nil {
		fmt.Println("Ck!")
		return
	}

	if addEl < root.val {
		if root.left == nil {
			root.left = &node{addEl, nil, nil}
		} else {
			if addEl < root.left.val {
				root.left.left = &node{addEl, nil, nil}
			} else {
				root.left.right = &node{addEl, nil, nil}
			}
		}
	} else {
		if root.right == nil {
			root.right = &node{addEl, nil, nil}
		} else {
			if addEl > root.right.val {
				root.right.right = &node{addEl, nil, nil}
			} else {
				root.right.left = &node{addEl, nil, nil}
			}
		}
	}
}

//root.left.left.val = addEl

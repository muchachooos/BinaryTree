package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	slice := []int{12, 2, 34, 7, 1, 65, 77, 9, 4, 3}

	var x int
	for _, value := range slice {
		x += value
	}
	average := x / len(slice)
	//fmt.Println(average)

	slice2 := make([]int, len(slice))

	for i, value := range slice {
		ttt := value - average
		slice2[i] = ttt
		if ttt < 0 {
			slice2[i] = ttt * (-1)
		}
	}
	fmt.Println(slice2)
	//root := node{}
}

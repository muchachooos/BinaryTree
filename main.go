package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

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

	fmt.Println(minNum)
	fmt.Println(slice[minNum])

	//root := node{}
}

//---------------------------------------------------
//slice := []int{34, 22, 34, 7, 77, 65, 27, 22, 46, 3}
//
//for i, value := range slice {
//	fmt.Println(i, value)
//}
//
//println("-------------")
//
//for i := range slice {
//	fmt.Println(i, slice[i])
//}
//
//println("-------------")
//
//for i := 0; i <= len(slice); i++ {
//	fmt.Println(i, slice[i])
//}

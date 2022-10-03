package main

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
	m := x / len(slice)

	var y int

	//root := node{}
}

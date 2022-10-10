package main

import (
	"testing"
)

func Test_addToEbala(t *testing.T) {
	el1 := node{
		9,
		nil,
		nil,
	}

	testSlice := []int{3, 4, 5, 8, 19}

	for i := range testSlice {
		addToTree(&el1, testSlice[i])
	}

	if el1.val != 9 {
		t.Fail()
	}

	if el1.left.val != 3 {
		t.Fail()
	}

	if el1.left.right.val != 4 {
		t.Fail()
	}

	if el1.left.right.right.val != 5 {
		t.Fail()
	}

	if el1.left.right.right.right.val != 8 {
		t.Fail()
	}

	if el1.right.val != 19 {
		t.Fail()
	}

	res := searchElement(&el1, 4)

	if res == false {
		t.Fail()
	}
}

package main

import (
	"testing"
)

func Test_FunctionTree(t *testing.T) {

	rootElement := node{
		9,
		nil,
		nil,
	}

	testSlice := []int{3, 4, 5, 8, 19}

	//Add function test
	for i := range testSlice {
		addElement(&rootElement, testSlice[i])
	}

	if rootElement.val != 9 {
		t.Log("add_1 failed")
		t.Fail()
	}

	if rootElement.left.val != 3 {
		t.Log("add_2 failed")
		t.Fail()
	}

	if rootElement.left.right.val != 4 {
		t.Log("add_3 failed")
		t.Fail()
	}

	if rootElement.left.right.right.val != 5 {
		t.Log("add_4 failed")
		t.Fail()
	}

	if rootElement.left.right.right.right.val != 8 {
		t.Log("add_5 failed")
		t.Fail()
	}

	if rootElement.right.val != 19 {
		t.Log("add_6 failed")
		t.Fail()
	}

	// Search function test
	for j := range testSlice {
		res := searchElement(&rootElement, testSlice[j])
		if res == false {
			t.Log("Search failed")
			t.Fail()
		}
	}

	//Different value test
	res := searchElement(&rootElement, 344)
	if res == true {
		t.Log("Search failed")
		t.Fail()
	}

	//Delete function test
	deleteElement(&rootElement, 4)
	if rootElement.left.right != nil {
		t.Log("Delete failed")
		t.Fail()
	}
}

package main

import (
	"testing"
)

func Test_sortSlice(t *testing.T) {

	testSlice := []int{3, 4, 5, 8, 9, 19}

	rootValue := sortSlice(testSlice)

	rootElement := node{
		rootValue,
		nil,
		nil,
	}

	if rootElement.val != rootValue {
		t.Log("Sort failed")
		t.Fail()
	}

	for i := range testSlice {
		if i == rootValue {
			t.Log("Delete root failed")
			t.Fail()
		}
	}
}

func Test_addElement(t *testing.T) {

	testSlice := []int{3, 4, 5, 8, 19}

	rootElement := node{
		9,
		nil,
		nil,
	}

	for i := range testSlice {
		addElement(&rootElement, testSlice[i])
	}

	if rootElement.val != 9 {
		t.Log("add_1 failed")
		t.Fail()
	}

	if rootElement.left != nil && rootElement.left.val != 3 {
		t.Log("add_2 failed")
		t.Fail()
	}

	if rootElement.left.right != nil && rootElement.left.right.val != 4 {
		t.Log("add_3 failed")
		t.Fail()
	}

	if rootElement.left.right.right != nil && rootElement.left.right.right.val != 5 {
		t.Log("add_4 failed")
		t.Fail()
	}

	if rootElement.left.right.right.right != nil && rootElement.left.right.right.right.val != 8 {
		t.Log("add_5 failed")
		t.Fail()
	}

	if rootElement.right != nil && rootElement.right.val != 19 {
		t.Log("add_6 failed")
		t.Fail()
	}
}

func Test_searchElement(t *testing.T) {

	testSlice := []int{3, 4, 5, 8, 19}

	rootElement := node{
		9,
		nil,
		nil,
	}

	for i := range testSlice {
		addElement(&rootElement, testSlice[i])
	}

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
		t.Log("Different value failed")
		t.Fail()
	}
}

func Test_deleteElement(t *testing.T) {

	testSlice := []int{3, 4, 5, 8, 19}

	rootElement := node{
		9,
		nil,
		nil,
	}

	for i := range testSlice {
		addElement(&rootElement, testSlice[i])
	}

	deleteElement(&rootElement, 4)
	if rootElement.left.right != nil {
		t.Log("Delete failed")
		t.Fail()
	}
}

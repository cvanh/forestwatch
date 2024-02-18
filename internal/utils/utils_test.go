package utils

import (
	"reflect"
	"testing"
)

type ArrayDiffObject struct {
	id   int
	name string
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
// func TestArrayDiffInt(t *testing.T) {
// 	// create 2 arrays we want to compare
// 	array1 := [3]int{0, 5}
// 	array2 := [3]int{4}

// 	// this is the end result we should expect
// 	diff := [1]int{4}

// 	out := ArrayDiff(array1, array2)

// 	if diff != out {
// 		t.Fatalf("ArrayDiff() failed. expected %v but received %v", diff, out)
// 	}
// }

// func TestArrayDiffint2(t *testing.T) {
// 	// create 2 arrays we want to compare
// 	array1 := [3]int{5, 10}
// 	array2 := [3]int{7}

// 	// this is the end result we should expect
// 	diff := [1]int{7}

// 	out := ArrayDiff(array1, array2)

// 	if diff != out {
// 		t.Fatalf("ArrayDiff() failed. expected %v but received %v", diff, out)
// 	}
// }

func TestArrayDiffObject(t *testing.T) {
	// create 2 arrays we want to compare
	array1 := []ArrayDiffObject{
		{
			id:   3,
			name: "jan",
		},
		{
			id:   10,
			name: "pieter",
		},
	}
	array2 := []ArrayDiffObject{
		{
			id:   3,
			name: "jan",
		},
	}

	// this is the end result we should expect
	diff := []ArrayDiffObject{
		{
			id:   3,
			name: "jan",
		},
	}

	out := ArrayDiff(array1, array2)

	if !reflect.DeepEqual(diff, out) {
		t.Fatalf("ArrayDiff() failed. expected %v but received %v", diff, out)
	}
}

// func TestArrayDiffEmptyIntrightArray(t *testing.T) {
// 	// create 2 arrays we want to compare
// 	array1 := [1]int{7}
// 	array2 := [0]int{}

// 	// this is the end result we should expect
// 	diff := [1]int{7}

// 	out := ArrayDiff(array1, array2)

// 	if diff != out {
// 		t.Fatalf("ArrayDiff() failed. expected %v but received %v", diff, out)
// 	}
// }

// func TestArrayDiffEmptyIntLeftArray(t *testing.T) {
// 	// create 2 arrays we want to compare
// 	array1 := [0]int{}
// 	array2 := [1]int{7}

// 	// this is the end result we should expect
// 	diff := [1]int{7}

// 	out := ArrayDiff(array1, array2)

// 	if diff != out {
// 		t.Fatalf("ArrayDiff() failed. expected %v but received %v", diff, out)
// 	}
// }

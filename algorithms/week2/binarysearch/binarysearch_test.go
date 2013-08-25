package binarysearch

import "testing"

type testcase struct {
	A []int
	V int
	E int
}

func Test_BinarySearch(t *testing.T) {
	tests := []testcase{
		// 0: empty array
		{A: []int{}, V:1, E:-1},
		// 1: one element, not in it
		{A: []int{0}, V:1, E:-1},
		// 2: one element, in it
		{A: []int{1}, V:1, E:0},
		// 3: two elements, not present
		{A: []int{0,0}, V:1, E:-1},
		// 4: two elements, first one
		{A: []int{1,2}, V:1, E:0},
		// 5: two elements, second one
		{A: []int{0,1}, V:1, E:1},
		// 6: three elements, first one
		{A: []int{0,1,2}, V:0, E:0},
		// 7: three elements, second one
		{A: []int{0,1,2}, V:1, E:1},
		// 8: three elements, third one
		{A: []int{0,1,2}, V:2, E:2},
		// 9: three elements, none
		{A: []int{0,1,2}, V:3, E:-1},

		// 10: four elements, none
		{A: []int{0,1,2,5}, V:3, E:-1},
		// 11: four elements, first one
		{A: []int{0,1,2,5}, V:0, E:0},
		// 12: four elements, second one
		{A: []int{0,1,2,5}, V:1, E:1},
		// 13: four elements, third one
		{A: []int{0,1,2,5}, V:2, E:2},
		// 14: four elements, fourth one
		{A: []int{0,1,2,5}, V:5, E:3},

	}

	for i, tc := range(tests) {
		if search(tc.A, tc.V) != tc.E {
			t.Error("test case", i)
		}
	}
}

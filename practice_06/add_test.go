// **************************************************
//
//		足し算 関数 TEST 用
//
//	         2024.08.25
//
// **************************************************
package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{10, 5, 15},
		{-1, -1, -2},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestErrorAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 3},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) != %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

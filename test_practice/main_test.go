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

// TestMul 関数のテスト
func TestMul(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},    // 2 * 3 = 6
		{-2, 3, -6},  // -2 * 3 = -6
		{0, 5, 0},    // 0 * 5 = 0
		{7, 0, 0},    // 7 * 0 = 0
		{-4, -5, 20}, // -4 * -5 = 20
		{1, 1, 1},    // 1 * 1 = 1
	}

	for _, test := range tests {
		result := Mul(test.a, test.b)
		if result != test.expected {
			t.Errorf("Mul(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

// **************************************************
//
//  掛け算 関数 TEST 用
//
//          2024.08.26
//
// **************************************************

package main

import "testing"

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

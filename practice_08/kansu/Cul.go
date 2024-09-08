// **************************************************
//
//  関数 足し算
//
//          2024.09.08
//
// **************************************************

package kansu

import "fmt"

// Add 関数
func Add(a int, b int) int {
	return a + b
}

// Sub 関数
func Sub(a int, b int) int {
	return a - b
}

// Mul 関数
func Mul(a int, b int) int {
	return a * b
}

// Div 関数（ゼロ除算にエラー処理を追加）
func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("ゼロで除算はできません")
	}
	return a / b, nil
}

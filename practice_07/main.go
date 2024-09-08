// **************************************************
//
//  関数 TEST 用
//
//          2024.09.03
//
// **************************************************

package main

import (
	"errors"
	"fmt"
)

// 四則演算の関数定義
func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0で除算することはできません")
	}
	return a / b, nil
}

// 関数のテスト
//
//(1)

func main() {
	// 四則演算関数のテスト

	fmt.Println(Add(1, 1))

	fmt.Println(Sub(1, 1))

	fmt.Println(Mul(1, 1))

	result, err := Div(1, 1)
	if err != nil {
		fmt.Println("エラーコード0:", err)
	} else {
		fmt.Println(result)
	}

	result, err = Div(1, 0)
	if err != nil {
		fmt.Println("エラーコード1:", err)
	} else {
		fmt.Println(result)
	}

	test_calculator()

}

// 関数のテスト
//
//(２)

func test_calculator() {

	var a, c int
	var b string

	a = 1
	b = "x"
	c = 0

	switch b {
	case "+":
		fmt.Println(Add(a, c))
	case "-":
		fmt.Println(Sub(a, c))
	case "*":
		fmt.Println(Mul(a, c))
	case "/":
		result, err := Div(a, c)
		if err != nil {
			fmt.Println("エラーコード2:", err)
		} else {
			fmt.Println(result)
		}
	default:
		fmt.Println("演算子が無効です")
	}
}

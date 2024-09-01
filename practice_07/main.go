// **************************************************
//
//	 四則演算 関数
//
//          2024.09.01
//
// **************************************************

package main

import (
	"errors"
	"fmt"
)

// 足し算関数

func Add(a, b int) int {

	return a + b

}

// 引き算関数

func Sub(d, e int) int {

	return d - e

}

// 掛け算関数

func Mul(f, g int) int {

	return f * g

}

// 割り算関数

func Div(h, i int) (int, error) {
	if i == 0 {
		return 0, errors.New("0で割ってはいけません")
	}
	return h / i, nil
}

func main() {
	fmt.Println(Add(1, 1)) // Output: 15
	fmt.Println(Sub(1, 1)) // Output: 5
	fmt.Println(Mul(1, 1)) // Output: 50

	// Handling the result of Div with error checking
	result, err := Div(1, 0)
	if err != nil {
		fmt.Println("エラーコード:", err)
	} else {
		fmt.Println(result)
	}
}

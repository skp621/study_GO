// **************************************************
//
//  四則演算
//
//          2024.09.08
//
// **************************************************

package main

import (
	"fmt"
	"study_GO/practice_08/kansu"
)

func main() {

	var a, b int
	var c string

	fmt.Println("左辺を入力してください")
	fmt.Scanln(&a)

	fmt.Println("演算子を入力してください")
	fmt.Scanln(&c)

	fmt.Println("右辺を入力してください")
	fmt.Scanln(&b)

	// 演算子に応じて計算を実行
	switch c {
	case "+":
		result := kansu.Add(a, b)
		fmt.Println("答え:", result)
	case "-":
		result := kansu.Sub(a, b)
		fmt.Println("答え:", result)
	case "*":
		result := kansu.Mul(a, b)
		fmt.Println("答え:", result)
	case "/":
		// 割り算でエラー処理
		result, err := kansu.Div(a, b)
		if err != nil {
			fmt.Println("エラー:", err)
		} else {
			fmt.Println("答え:", result)
		}
	default:
		fmt.Println("無効な演算子が入力されました。+、-、*、/ のいずれかを入力してください。")
	}
}

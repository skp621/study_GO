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

	// 最初の数値を入力
	fmt.Scanln(&a)

	// 演算子を入力
	fmt.Scanln(&c)

	// 次の数値を入力
	fmt.Scanln(&b)

	// 演算子に応じて計算を実行
	switch c {
	case "+":
		result := kansu.Add(a, b)
		fmt.Println("結果:", result)
	case "-":
		result := kansu.Sub(a, b)
		fmt.Println("結果:", result)
	case "*":
		result := kansu.Mul(a, b)
		fmt.Println("結果:", result)
	case "/":
		// 割り算でエラー処理
		result, err := kansu.Div(a, b)
		if err != nil {
			fmt.Println("エラー:", err)
		} else {
			fmt.Println("結果:", result)
		}
	default:
		fmt.Println("無効な演算子が入力されました。+、-、*、/ のいずれかを入力してください。")
	}
}

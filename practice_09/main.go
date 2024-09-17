// **************************************************
//
//  世界のナベアツ
//
//          2024.09.15
//
// **************************************************

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var a, b int
	var c string

	fmt.Println("開始数字を入力してください")
	fmt.Scanln(&a)

	fmt.Println("終了数字を入力してください")
	fmt.Scanln(&b)

	//入力値の大小により分岐
	var max int
	var min int

	if a > b {

		max = a
		min = b

	} else {

		max = b
		min = a
	}

	for i := min; i <= max; i++ {

		// todo: ３の倍数の判定を追加

		c = strconv.Itoa(i)

		if strings.Contains(c, "3") {
			fmt.Println("さん！")
		} else {
			fmt.Println(i)
		}
	}
}

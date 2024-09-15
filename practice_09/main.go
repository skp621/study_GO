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

	if a > b {

		for i := a; i >= b; i-- {

			c = strconv.Itoa(i)

			if strings.Contains(c, "3") {
				fmt.Println("さん！")
			} else {
				fmt.Println(i)
			}
		}

	} else {

		for i := a; i <= b; i++ {

			c = strconv.Itoa(i)

			if strings.Contains(c, "3") {
				fmt.Println("さん！")
			} else {
				fmt.Println(i)
			}
		}
	}
}

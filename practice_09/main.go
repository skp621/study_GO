// **************************************************
//
//  世界のナベアツ
//
//          2024.09.20
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

	fmt.Println("3を含むほうが、3の倍数よりも、優先されます！")

	for i := min; i <= max; i++ {

		c = strconv.Itoa(i)

		if strings.Contains(c, "3") {

			fmt.Println("さん！(含む)")

		} else {

			if i%3 == 0 {

				fmt.Println("さん！(倍数)")

			} else {

				fmt.Println(i)

			}
		}
	}
}

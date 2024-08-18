//**************************************************
//
//  FizzBuzz 関数化
//										2024.08.18
//**************************************************

package main

import "fmt"

// 関数
func fizzbuzz(j int) string {

	var k string

	if j%3 == 0 && j%5 == 0 {

		k = "FizzBuzz"
		return k

	} else if j%3 == 0 {

		k = "Fizz"
		return k

	} else if j%5 == 0 {

		k = "Buzz"
		return k

	} else {

		k = fmt.Sprintf("%d", j)
		return k

	}

}

func main() {

	for i := 1; i < 21; i++ {

		result := fizzbuzz(i)

		fmt.Println(result)

	}

}

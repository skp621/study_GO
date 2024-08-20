//**************************************************
//
//  Reverse FizzBuzz 関数
//										2024.08.20
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

	for i := 15; i >= 1; i-- {

		result := fizzbuzz(i)

		fmt.Println(result)

	}

}

//**************************************************
//
//  FizzBuzz 関数
//										2024.08.19
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

	fmt.Println("<< 1から20まで昇順でFizzBuzz判定 >>")

	for i := 1; i < 21; i++ {

		result := fizzbuzz(i)

		fmt.Println(result)

	}

	fmt.Println("<< 15から1まで降順でFizzBuzz判定 >>")

	for i := 15; i >= 1; i-- {

		result := fizzbuzz(i)

		fmt.Println(result)

	}

}

//**************************************************
//
//  FizzBuzz 関数化
//										2024.08.18
//**************************************************

package main

import "fmt"

// 関数
func fizzbuzz(j k) k {

	if j%3 == 0 && j%5 == 0 {

		k = "FizzBuzz"

	} else if j%3 == 0 {

		k = "Fizz"

	} else if j%5 == 0 {

		k = "Buzz"

	} else {

		k = j

	}

}

func main() {

	for i := 1; i < 21; i++ {

		if i%3 == 0 && i%5 == 0 {

			fmt.Println("FizzBuzz")

		} else if i%3 == 0 {

			fmt.Println("Fizz")

		} else if i%5 == 0 {

			fmt.Println("Buzz")

		} else {

			fmt.Println(i)

		}

	}

}

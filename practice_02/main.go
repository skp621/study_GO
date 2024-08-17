//**************************************************
//
//  FizzBuzz
//										2024.08.17
//**************************************************

package main

import "fmt"

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

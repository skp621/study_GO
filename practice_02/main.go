//**************************************************
//
//  FizzBuzz
//										2024.08.17
//**************************************************

package main

import "fmt"

func main() {

	for i := 1; i < 21; i++ { // perform varying i from 1 by 1 until i > 20

		if i%3 == 0 && i%5 == 0 { // and ではない &&

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

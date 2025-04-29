package main

import "fmt"

func fizzbuzz() {
	for i := range 101 {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("FizzBuzz")
	fizzbuzz()
}

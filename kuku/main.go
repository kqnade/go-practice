package main

import "fmt"

func kuku() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	kuku()
}
